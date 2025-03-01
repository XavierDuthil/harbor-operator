package storage

import (
	"context"

	"github.com/go-logr/logr"
	goharborv1 "github.com/goharbor/harbor-operator/apis/goharbor.io/v1alpha3"
	minio "github.com/goharbor/harbor-operator/pkg/cluster/controllers/storage/minio/api/v1"
	"github.com/goharbor/harbor-operator/pkg/cluster/k8s"
	"github.com/goharbor/harbor-operator/pkg/cluster/lcm"
	"github.com/ovh/configstore"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	Storage = "storage"

	DefaultCredsSecret = "creds"
	DefaultPrefix      = "minio-"
	DefaultZone        = "zone-harbor"
	DefaultRegion      = "us-east-1"
	DefaultBucket      = "harbor"
	DefaultServicePort = 9000
)

type MinIOController struct {
	KubeClient  client.Client
	Ctx         context.Context
	Log         logr.Logger
	Scheme      *runtime.Scheme
	Recorder    record.EventRecorder
	MinioClient Minio
	ConfigStore *configstore.Store
}

var HarborClusterMinIOGVK = schema.GroupVersionKind{
	Group:   minio.SchemeGroupVersion.Group,
	Version: minio.SchemeGroupVersion.Version,
	Kind:    minio.MinIOCRDResourceKind,
}

func NewMinIOController(options ...k8s.Option) lcm.Controller {
	o := &k8s.CtrlOptions{}

	for _, option := range options {
		option(o)
	}

	return &MinIOController{
		KubeClient:  o.Client,
		Log:         o.Log,
		Scheme:      o.Scheme,
		ConfigStore: o.ConfigStore,
	}
}

// Reconciler implements the reconcile logic of minIO service.
func (m *MinIOController) Apply(ctx context.Context, harborcluster *goharborv1.HarborCluster, _ ...lcm.Option) (*lcm.CRStatus, error) {
	// Apply minIO tenant
	if crs, err := m.applyTenant(ctx, harborcluster); err != nil {
		return crs, err
	}

	// Check readiness
	mt, ready, err := m.checkMinIOReady(ctx, harborcluster)
	if err != nil {
		return minioNotReadyStatus(GetMinIOError, err.Error()), err
	}

	if !ready {
		m.Log.Info("MinIO is not ready yet")

		return minioUnknownStatus(), nil
	}

	// Apply minIO ingress if necessary
	if crs, err := m.applyIngress(ctx, harborcluster); err != nil {
		return crs, err
	}

	// Init minio
	// TODO: init bucket in the minCR pods as creation by client may meet network connection issue.
	if err := m.minioInit(ctx, harborcluster); err != nil {
		return minioNotReadyStatus(CreateDefaultBucketError, err.Error()), err
	}

	crs, err := m.ProvisionMinIOProperties(ctx, harborcluster, mt)
	if err != nil {
		return crs, err
	}

	m.Log.Info("MinIO is ready")

	return crs, nil
}

func (m *MinIOController) Delete(ctx context.Context, harborcluster *goharborv1.HarborCluster) (*lcm.CRStatus, error) {
	minioCR, err := m.generateMinIOCR(ctx, harborcluster)
	if err != nil {
		return minioNotReadyStatus(GenerateMinIOCrError, err.Error()), err
	}

	if err := m.KubeClient.Delete(ctx, minioCR); err != nil {
		return minioUnknownStatus(), err
	}

	return nil, nil
}

func (m *MinIOController) Upgrade(_ context.Context, _ *goharborv1.HarborCluster) (*lcm.CRStatus, error) {
	panic("implement me")
}

func (m *MinIOController) minioInit(ctx context.Context, harborcluster *goharborv1.HarborCluster) error {
	accessKey, secretKey, err := m.getCredsFromSecret(ctx, harborcluster)
	if err != nil {
		return err
	}

	endpoint := m.getServiceName(harborcluster) + "." + harborcluster.Namespace + ":9000"

	edp := &MinioEndpoint{
		Endpoint:        endpoint,
		AccessKeyID:     string(accessKey),
		SecretAccessKey: string(secretKey),
		Location:        DefaultRegion,
	}

	m.MinioClient, err = NewMinioClient(edp)
	if err != nil {
		return err
	}

	exists, err := m.MinioClient.IsBucketExists(ctx, DefaultBucket)
	if err != nil || exists {
		return err
	}

	return m.MinioClient.CreateBucket(ctx, DefaultBucket)
}

func (m *MinIOController) checkMinIOReady(ctx context.Context, harborcluster *goharborv1.HarborCluster) (*minio.Tenant, bool, error) {
	minioCR := &minio.Tenant{}
	if err := m.KubeClient.Get(ctx, m.getMinIONamespacedName(harborcluster), minioCR); err != nil {
		if errors.IsNotFound(err) {
			return nil, false, nil
		}

		return nil, false, err
	}

	// For different version of minIO have different Status.
	// Ref https://github.com/minio/operator/commit/d387108ea494cf5cec57628c40d40604ac8d57ec#diff-48972613166d50a2acb9d562e33c5247
	if minioCR.Status.CurrentState == minio.StatusReady || minioCR.Status.CurrentState == minio.StatusInitialized {
		return minioCR, true, nil
	}

	// Not ready
	return minioCR, false, nil
}

func (m *MinIOController) getMinIONamespacedName(harborcluster *goharborv1.HarborCluster) types.NamespacedName {
	return types.NamespacedName{
		Namespace: harborcluster.Namespace,
		Name:      m.getServiceName(harborcluster),
	}
}

func (m *MinIOController) getMinIOSecretNamespacedName(harborcluster *goharborv1.HarborCluster) types.NamespacedName {
	secretName := harborcluster.Spec.InClusterStorage.MinIOSpec.SecretRef
	if secretName == "" {
		secretName = DefaultPrefix + harborcluster.Name + "-" + DefaultCredsSecret
	}

	return types.NamespacedName{
		Namespace: harborcluster.Namespace,
		Name:      secretName,
	}
}

func (m *MinIOController) getServiceName(harborcluster *goharborv1.HarborCluster) string {
	return DefaultPrefix + harborcluster.Name
}

func minioNotReadyStatus(reason, message string) *lcm.CRStatus {
	now := metav1.Now()

	return &lcm.CRStatus{
		Condition: goharborv1.HarborClusterCondition{
			Type:               goharborv1.StorageReady,
			Status:             corev1.ConditionFalse,
			LastTransitionTime: &now,
			Reason:             reason,
			Message:            message,
		},
		Properties: nil,
	}
}

func minioUnknownStatus() *lcm.CRStatus {
	now := metav1.Now()

	return &lcm.CRStatus{
		Condition: goharborv1.HarborClusterCondition{
			Type:               goharborv1.StorageReady,
			Status:             corev1.ConditionUnknown,
			LastTransitionTime: &now,
			Reason:             "",
			Message:            "",
		},
		Properties: nil,
	}
}

func minioReadyStatus(properties *lcm.Properties) *lcm.CRStatus {
	now := metav1.Now()

	return &lcm.CRStatus{
		Condition: goharborv1.HarborClusterCondition{
			Type:               goharborv1.StorageReady,
			Status:             corev1.ConditionTrue,
			LastTransitionTime: &now,
			Reason:             "",
			Message:            "",
		},
		Properties: *properties,
	}
}
