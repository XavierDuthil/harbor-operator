package portal

import (
	"context"

	goharborv1 "github.com/goharbor/harbor-operator/apis/goharbor.io/v1alpha3"
	harbormetav1 "github.com/goharbor/harbor-operator/apis/meta/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (r *Reconciler) GetService(ctx context.Context, portal *goharborv1.Portal) (*corev1.Service, error) {
	name := r.NormalizeName(ctx, portal.GetName())
	namespace := portal.GetNamespace()

	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{
				Name:       harbormetav1.PortalHTTPPortName,
				Port:       harbormetav1.HTTPPort,
				TargetPort: intstr.FromString(harbormetav1.PortalHTTPPortName),
				Protocol:   corev1.ProtocolTCP,
			}, {
				Name:       harbormetav1.PortalHTTPSPortName,
				Port:       harbormetav1.HTTPSPort,
				TargetPort: intstr.FromString(harbormetav1.PortalHTTPSPortName),
				Protocol:   corev1.ProtocolTCP,
			}},
			Selector: map[string]string{
				r.Label("name"):      name,
				r.Label("namespace"): namespace,
			},
		},
	}, nil
}
