package controllers

import (
	"context"
	"errors"
	"time"

	demov1alpha1 "github.com/basit9958/configmap-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DemoClusterConfigmap Controller", func() {
	const (
		ccmName = "test-ccm"
		ccmNs   = "default"

		timeout  = time.Second * 10
		interval = time.Millisecond * 250
	)
	labelSelectors := metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app.kubernetes.io/context": "test",
		},
	}

	Context("When updating DemoClusterConfigmap status", func() {
		It("Should add the created ConfigMaps to DemoClusterConfigmap.Status.Namespaces when it's created.", func() {
			By("Creating a new DemoClusterConfigmap")
			ctx := context.Background()
			ccm := demov1alpha1.DemoClusterConfigmap{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "github.com/v1alpha1",
					Kind:       "DemoClusterConfigmap",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      ccmName,
					Namespace: ccmNs,
				},
				Spec: demov1alpha1.DemoClusterConfigmapSpec{
					Data: map[string]string{
						"this is": "the way",
					},
					GenerateTo: demov1alpha1.ClusterConfigMapGenerateTo{
						NamespaceSelectors: labelSelectors,
					},
				},
			}
			Expect(k8sClient.Create(ctx, &ccm)).Should(Succeed())
			Eventually(func() bool {
				if err := k8sClient.Get(ctx, types.NamespacedName{Namespace: ccmNs, Name: ccmName}, &ccm); err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())

			By("Checking whether a ConfigMap with the same data got created in the matching namespace.")
			Eventually(func() ([]string, error) {
				var cmList corev1.ConfigMapList
				if err := k8sClient.List(ctx, &cmList); err != nil {
					return nil, err
				}
				for _, cm := range cmList.Items {
					GinkgoWriter.Write([]byte(cm.Name))
				}
				return nil, errors.New("")
			})
		})
	})
})
