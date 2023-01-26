package controllers_test

import (
	"context"

	infrakitv1alpha1 "github.com/INFURA/infrakit/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Proxy controller", func() {
	It("Works", func() {
		ctx := context.Background()
		proxy := &infrakitv1alpha1.Proxy{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-proxy",
				Namespace: "default",
			},
		}
		Expect(k8sClient.Create(ctx, proxy)).Should(Succeed())
	})
})
