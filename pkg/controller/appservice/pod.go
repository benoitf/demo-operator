package appservice

import (
	appv1alpha1 "github.com/eivantsov/demo-operator/pkg/apis/app/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// newPodForCR returns an nginx pod with the same name/namespace as the cr
func newPodForCR(cr *appv1alpha1.AppService, labels map[string]string) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx",
					Image: "nginx:1.16-alpine",
				},
			},
		},
	}
}
