package controllers_test

import (
	"context"
	"time"

	infrakitv1alpha1 "github.com/INFURA/infrakit/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("PodTemplateClientSet controller", func() {
	ptcsName := "test-ptcs"
	ptcsNamespace := "default"

	It("Works", func() {
		ctx := context.Background()
		ptcs := &infrakitv1alpha1.PodTemplateClientSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      ptcsName,
				Namespace: ptcsNamespace,
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

		// Check to see that the ptcs is created
		ptcsLookupKey := types.NamespacedName{Name: ptcsName, Namespace: ptcsNamespace}
		createdPtcs := &infrakitv1alpha1.PodTemplateClientSet{}

		Eventually(func() bool {
			err := k8sClient.Get(ctx, ptcsLookupKey, createdPtcs)
			return err == nil
		}, time.Second*10, time.Millisecond*250).Should(BeTrue())

		// Check to see that the statefulset is created
		stsLookupKey := types.NamespacedName{Name: ptcsName, Namespace: ptcsNamespace}
		sts := &appsv1.StatefulSet{}

		Eventually(func() bool {
			err := k8sClient.Get(ctx, stsLookupKey, sts)
			return err == nil
		}, time.Second*10, time.Millisecond*250).Should(BeTrue())

		Expect(sts.Spec.Selector).ShouldNot(BeNil())
		Expect(sts.Spec.Selector.MatchLabels).ShouldNot(BeNil())
	})
})
