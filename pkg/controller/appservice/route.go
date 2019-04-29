package appservice

import (
	appv1alpha1 "github.com/eivantsov/demo-operator/pkg/apis/app/v1alpha1"
	routev1 "github.com/openshift/api/route/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func newRoute(cr *appv1alpha1.AppService, name string, serviceName string, port int32, labels map[string]string) *routev1.Route {
	targetPort := intstr.IntOrString{
		Type:   intstr.Int,
		IntVal: int32(port),
	}
	return &routev1.Route{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Route",
			APIVersion: routev1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: routev1.RouteSpec{
			To: routev1.RouteTargetReference{
				Kind: "Service",
				Name: serviceName,
			},
			Port: &routev1.RoutePort{
				targetPort,
			},
		},
	}
}
