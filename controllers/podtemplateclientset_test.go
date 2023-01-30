package controllers_test

import (
	"context"

	infrakitv1alpha1 "github.com/INFURA/infrakit/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("PodTemplateClientSet controller", func() {
	It("Works", func() {
		ctx := context.Background()
		ptcs := &infrakitv1alpha1.PodTemplateClientSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-ptcs",
				Namespace: "default",
			},
			Spec: infrakitv1alpha1.PodTemplateClientSetSpec{
				Template: corev1.PodTemplateSpec{
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{},
					},
				},
			},
		}
		Expect(k8sClient.Create(ctx, ptcs)).Should(Succeed())
	})
})
