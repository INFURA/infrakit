package controllers_test

import (
	"context"

	infrakitv1alpha1 "github.com/INFURA/infrakit/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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
		}
		Expect(k8sClient.Create(ctx, ptcs)).Should(Succeed())
	})
})
