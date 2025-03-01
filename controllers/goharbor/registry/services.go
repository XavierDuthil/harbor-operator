package registry

import (
	"context"

	goharborv1 "github.com/goharbor/harbor-operator/apis/goharbor.io/v1alpha3"
	harbormetav1 "github.com/goharbor/harbor-operator/apis/meta/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (r *Reconciler) GetService(ctx context.Context, registry *goharborv1.Registry) (*corev1.Service, error) {
	name := r.NormalizeName(ctx, registry.GetName())
	namespace := registry.GetNamespace()

	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{
				Name:       harbormetav1.RegistryAPIPortName,
				Port:       registry.Spec.HTTP.TLS.GetInternalPort(),
				TargetPort: intstr.FromString(harbormetav1.RegistryAPIPortName),
				Protocol:   corev1.ProtocolTCP,
			}, {
				Name:       harbormetav1.RegistryMetricsPortName,
				Port:       registry.Spec.HTTP.TLS.GetInternalPort() + 1,
				TargetPort: intstr.FromString(harbormetav1.RegistryMetricsPortName),
				Protocol:   corev1.ProtocolTCP,
			}},
			Selector: map[string]string{
				r.Label("name"):      name,
				r.Label("namespace"): namespace,
			},
		},
	}, nil
}
