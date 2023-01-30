package resource_test

import (
	"strings"

	"github.com/INFURA/infrakit/pkg/resource"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8sResource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

func assertComparisonResult(
	statefulSetA, statefulSetB appsv1.StatefulSet,
	check func(resource.CompareStatefulsetResult) bool,
) {
	comparer := resource.NewStatefulSetComparer(statefulSetA)
	result, err := comparer.Compare(statefulSetB)
	Expect(err).ToNot(HaveOccurred())
	Expect(check(*result)).To(BeTrue())

	_, _ = GinkgoWriter.Write([]byte("Verbose diff:" + strings.Join(result.Reasons(), "; ")))
}

func matchButNotReplace(result resource.CompareStatefulsetResult) bool {
	return result.Match() && !result.Replace()
}

func mismatchButNotReplace(result resource.CompareStatefulsetResult) bool {
	return !result.Match() && !result.Replace()
}

func mismatchAndReplace(result resource.CompareStatefulsetResult) bool {
	return !result.Match() && result.Replace()
}

func volumeClaimTemplatesMustParse(str string) []corev1.PersistentVolumeClaim {
	return []corev1.PersistentVolumeClaim{
		{
			Spec: corev1.PersistentVolumeClaimSpec{
				Resources: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: k8sResource.MustParse(str),
					},
				},
			},
		},
	}
}

var _ = Describe("SemanticEquality", func() {
	It("Compares Ports", func() {
		portsA := []corev1.ContainerPort{
			{
				Name:          "api",
				ContainerPort: 8545,
			},
		}

		portsB := []corev1.ContainerPort{
			{
				Name:          "api",
				ContainerPort: 8545,
				Protocol:      corev1.ProtocolTCP,
			},
		}

		Expect(resource.Semantic.DeepEqual(portsA, portsB)).To(BeTrue())
	})
})

var _ = Describe("CompareStatefulSet", func() {
	DescribeTable("When statefulsets are matching",
		assertComparisonResult,
		Entry("returns matching when nothing differs",
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(5),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{Name: "foobar"},
							},
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{},
				},
			},
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(5),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{Name: "foobar"},
							},
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{},
				},
			},
			matchButNotReplace,
		),
		Entry("returns matching when only resource format differs",
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(5),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{Name: "foobar"},
							},
						},
					},
					VolumeClaimTemplates: volumeClaimTemplatesMustParse("5120Gi"), // equivalent to 5 Ti
				},
			},
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(5),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{Name: "foobar"},
							},
						},
					},
					VolumeClaimTemplates: volumeClaimTemplatesMustParse("5Ti"),
				},
			},
			matchButNotReplace,
		),
	)

	DescribeTable("When statefulsets are not matching but needn't be replaced",
		assertComparisonResult,
		Entry("returns not matching when replicas differ",
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(5),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{},
							Containers:     []corev1.Container{{Name: "foobar"}},
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{},
				},
			},
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(20),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{},
							Containers:     []corev1.Container{{Name: "foobar"}},
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{},
				},
			},
			mismatchButNotReplace,
		),
		Entry("returns not matching when containers differ",
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(20),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{},
							Containers:     []corev1.Container{{Name: "foobar"}},
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{},
				},
			},
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(20),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{{Name: "initcontainer"}},
							Containers:     []corev1.Container{{Name: "foobar"}},
							RestartPolicy:  corev1.RestartPolicyAlways,
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{},
				},
			},
			mismatchButNotReplace,
		),
		Entry("returns not matching when other template elements differ",
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(33),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers:         []corev1.Container{{Name: "foobar"}},
							ServiceAccountName: "joe",
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{},
				},
			},
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(33),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers:         []corev1.Container{{Name: "foobar"}},
							ServiceAccountName: "rogan",
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{},
				},
			},
			mismatchButNotReplace,
		),
	)

	DescribeTable("When statefulset needs replacing",
		assertComparisonResult,
		Entry("returns replace when volumeClaimTemplates count differs",
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(5),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{},
							Containers:     []corev1.Container{{Name: "foobar"}},
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{},
				},
			},
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(5),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{},
							Containers:     []corev1.Container{{Name: "foobar"}},
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
						{
							ObjectMeta: metav1.ObjectMeta{Name: "foobar"},
						},
					},
				},
			},
			mismatchAndReplace,
		),
		Entry("returns replace when volumeClaimTemplates name differs ",
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(5),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{},
							Containers:     []corev1.Container{{Name: "foobar"}},
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
						{
							ObjectMeta: metav1.ObjectMeta{Name: "linus"},
						},
					},
				},
			},
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(5),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{},
							Containers:     []corev1.Container{{Name: "foobar"}},
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
						{
							ObjectMeta: metav1.ObjectMeta{Name: "torvalds"},
						},
					},
				},
			},
			mismatchAndReplace,
		),
		Entry("returns replace when volumeClaimTemplates spec differs (most interesting case)",
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(5),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{{Name: "foobar"}},
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
						{
							Spec: corev1.PersistentVolumeClaimSpec{
								DataSource: &corev1.TypedLocalObjectReference{
									Name: "this-snapshot",
								},
							},
						},
					},
				},
			},
			appsv1.StatefulSet{
				Spec: appsv1.StatefulSetSpec{
					Replicas: pointer.Int32Ptr(5),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{{Name: "foobar"}},
						},
					},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
						{
							Spec: corev1.PersistentVolumeClaimSpec{
								DataSource: &corev1.TypedLocalObjectReference{
									Name: "no-longer-this-snapshot",
								},
							},
						},
					},
				},
			},
			mismatchAndReplace,
		),
	)
})
