//nolint:wsl,lll,funlen,godot
package resource

import (
	"fmt"
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

// This is an adapted port from postgres-operator.
// https://github.com/zalando/postgres-operator/blob/e03e9f919a0f5b8effc6ab3a6ca91f051566c196/pkg/cluster/cluster.go#L366
//
// The statefulset comparison process is visibly complicated since we have to
// compare sub-properties at one time, not the whole thing
//
// As per the API validation code for statefulsets [here](https://github.com/kubernetes/kubernetes/blob/d159ae35450ca4b2adcf35c63af3e71bc2088f09/pkg/apis/apps/validation/validation.go#L157):
// "updates to statefulset spec for fields other than 'replicas', 'template', and 'updateStrategy' are forbidden",
// therefore we'll aim to return "replace" only if we spot mismatches on fields other than the above.

type CompareStatefulsetResult struct {
	match   bool
	replace bool
	reasons []string
}

func (c *CompareStatefulsetResult) Match() bool       { return c.match }
func (c *CompareStatefulsetResult) Replace() bool     { return c.replace }
func (c *CompareStatefulsetResult) Reasons() []string { return c.reasons }

type CompareStatefulSetError struct{ msg string }

func (e *CompareStatefulSetError) Error() string {
	return e.msg
}

// StatefulSetComparer enacapsulates logic of comparing two statefulsets.
type StatefulSetComparer struct {
	Statefulset appsv1.StatefulSet
}

func NewStatefulSetComparer(statefulset appsv1.StatefulSet) *StatefulSetComparer {
	return &StatefulSetComparer{
		Statefulset: statefulset,
	}
}

// Compare runs the comparison by looking at corresponding sub-attributes of the statefulsets.
// Note that it is still not thorough, comparing the significant differences.
// Though not highly probable, more comparisons can be added as needed here.
func (c *StatefulSetComparer) Compare(statefulSet appsv1.StatefulSet) (*CompareStatefulsetResult, error) { //nolint: gocognit
	var (
		reasons      = make([]string, 0)
		match        = true
		needsReplace = false
	)

	if len(c.Statefulset.Spec.Template.Spec.Containers) == 0 {
		return nil, &CompareStatefulSetError{msg: "statefulset has no container"}
	}

	// CHAPTER 1: Compare updatable attributes belonging to Template, Replicas
	// /////////////////////////////

	if *c.Statefulset.Spec.Replicas != *statefulSet.Spec.Replicas {
		match = false
		reasons = append(reasons, "new statefulset's number of replicas doesn't match the current one")
	}
	// XXX: not looking at annotations since not currently involved in any functionality.
	// The "common.infura.io/hash" is tricking us.
	// if !reflect.DeepEqual(c.Statefulset.Annotations, statefulSet.Annotations) {
	// 	match = false
	// 	reasons = append(reasons, "new statefulset's annotations doesn't match the current one")
	// }

	// Compare init containers
	containerReasons := c.compareContainers(
		"initContainers",
		c.Statefulset.Spec.Template.Spec.InitContainers,
		statefulSet.Spec.Template.Spec.InitContainers,
	)
	if len(containerReasons) > 0 {
		match = false
		reasons = append(reasons, containerReasons...)
	}

	// Compare containers
	containerReasons = c.compareContainers(
		"containers",
		c.Statefulset.Spec.Template.Spec.Containers,
		statefulSet.Spec.Template.Spec.Containers,
	)
	if len(containerReasons) > 0 {
		match = false
		reasons = append(reasons, containerReasons...)
	}

	// Compare deeper updatable attributes
	if c.Statefulset.Spec.Template.Spec.ServiceAccountName != statefulSet.Spec.Template.Spec.ServiceAccountName {
		match = false
		reasons = append(reasons, "new statefulset's serviceAccountName service account name doesn't match the current one")
	}

	if c.Statefulset.Spec.Template.Spec.TerminationGracePeriodSeconds != nil &&
		statefulSet.Spec.Template.Spec.TerminationGracePeriodSeconds != nil &&
		*c.Statefulset.Spec.Template.Spec.TerminationGracePeriodSeconds != *statefulSet.Spec.Template.Spec.TerminationGracePeriodSeconds {
		match = false
		reasons = append(reasons, "new statefulset's terminationGracePeriodSeconds doesn't match the current one")
	}

	if !reflect.DeepEqual(c.Statefulset.Spec.Template.Spec.Affinity, statefulSet.Spec.Template.Spec.Affinity) {
		match = false
		reasons = append(reasons, "new statefulset's pod affinity doesn't match the current one")
	}

	if !reflect.DeepEqual(c.Statefulset.Spec.Template.Labels, statefulSet.Spec.Template.Labels) {
		match = false
		reasons = append(reasons, "new statefulset's metadata labels doesn't match the current one")
	}

	if !reflect.DeepEqual(c.Statefulset.Spec.Template.Annotations, statefulSet.Spec.Template.Annotations) {
		match = false
		reasons = append(reasons, "new statefulset's pod template metadata annotations doesn't match the current one")
	}

	if !reflect.DeepEqual(c.Statefulset.Spec.Template.Spec.SecurityContext, statefulSet.Spec.Template.Spec.SecurityContext) {
		match = false
		reasons = append(reasons, "new statefulset's pod template security context in spec doesn't match the current one")
	}

	// CHAPTER 2: Compare attributes that are not updateable which require replacing the statefulset
	// /////////////////////////////

	// Compare PodManagementPolicy
	if c.Statefulset.Spec.PodManagementPolicy != statefulSet.Spec.PodManagementPolicy {
		needsReplace = true
		reasons = append(reasons, "new statefulset's PodManagementPolicy doesn't match the current one")
	}

	// Compare selectors
	if (c.Statefulset.Spec.Selector != nil) && (statefulSet.Spec.Selector != nil) {
		if !reflect.DeepEqual(c.Statefulset.Spec.Selector.MatchLabels, statefulSet.Spec.Selector.MatchLabels) {
			// forbid introducing new labels in the selector on the new statefulset, as it would cripple replacements
			// due to the fact that the new statefulset won't be able to pick up old pods with non-matching labels.
			if !mapContains(c.Statefulset.Spec.Selector.MatchLabels, statefulSet.Spec.Selector.MatchLabels) {
				return nil, &CompareStatefulSetError{msg: "new statefulset introduces extra labels in the label selector, cannot continue"}
			}
			needsReplace = true
			reasons = append(reasons, "new statefulset's selector doesn't match the current one")
		}
	}

	// Compare VolumeClaimTemplates
	if len(c.Statefulset.Spec.VolumeClaimTemplates) != len(statefulSet.Spec.VolumeClaimTemplates) {
		needsReplace = true
		reasons = append(reasons, "new statefulset's volumeClaimTemplates contains different number of volumes to the old one")
	}
	for i := 0; i < len(c.Statefulset.Spec.VolumeClaimTemplates); i++ {
		name := c.Statefulset.Spec.VolumeClaimTemplates[i].Name
		// Some generated fields like creationTimestamp make it not possible to use DeepCompare on ObjectMeta
		if name != statefulSet.Spec.VolumeClaimTemplates[i].Name {
			needsReplace = true
			reasons = append(reasons, fmt.Sprintf("new statefulset's name for volume %d doesn't match the current one", i))

			continue
		}
		if !reflect.DeepEqual(c.Statefulset.Spec.VolumeClaimTemplates[i].Annotations, statefulSet.Spec.VolumeClaimTemplates[i].Annotations) {
			needsReplace = true
			reasons = append(reasons, fmt.Sprintf("new statefulset's annotations for volume %q doesn't match the current one", name))
		}

		// Semantic correctly deals with things like resource.Quantity values
		if !Semantic.DeepEqual(c.Statefulset.Spec.VolumeClaimTemplates[i].Spec, statefulSet.Spec.VolumeClaimTemplates[i].Spec) {
			name := c.Statefulset.Spec.VolumeClaimTemplates[i].Name
			needsReplace = true
			reasons = append(reasons, fmt.Sprintf("new statefulset's volumeClaimTemplates specification for volume %q doesn't match the current one", name))
		}
	}

	if needsReplace {
		match = false
	}

	return &CompareStatefulsetResult{match: match, reasons: reasons, replace: needsReplace}, nil
}

type containerCondition func(a, b v1.Container) bool

type containerCheck struct {
	condition containerCondition
	reason    string
}

func newCheck(msg string, cond containerCondition) containerCheck {
	return containerCheck{reason: msg, condition: cond}
}

// compareContainers: compare two list of Containers
// and return:
// * whether or not a rolling update is needed
// * a list of reasons in a human readable format
func (c *StatefulSetComparer) compareContainers(description string, setA, setB []v1.Container) []string {
	reasons := make([]string, 0)

	if len(setA) != len(setB) {
		reasons = append(reasons, "new statefulset container count does not match the current ones")

		return reasons
	}

	checks := []containerCheck{
		newCheck("%s: %s (index %d) name doesn't match the current one",
			func(a, b v1.Container) bool { return a.Name != b.Name }),
		newCheck("%s: %s (index %d) ports don't match the current one",
			func(a, b v1.Container) bool { return !Semantic.DeepEqual(a.Ports, b.Ports) }),
		newCheck("%s: %s (index %d) resources don't match the current ones",
			func(a, b v1.Container) bool { return !compareResources(&a.Resources, &b.Resources) }),
		newCheck("%s: %s (index %d) environment doesn't match the current one",
			func(a, b v1.Container) bool {
				// Empty slices can return different with DeepEqual depending on how constructed
				if len(a.Env) == 0 && len(b.Env) == 0 {
					return false
				}

				return !reflect.DeepEqual(a.Env, b.Env)
			}),
		newCheck("%s: %s (index %d) environment sources don't match the current one",
			func(a, b v1.Container) bool { return !reflect.DeepEqual(a.EnvFrom, b.EnvFrom) }),
	}

	for index, containerA := range setA {
		containerB := setB[index]

		for _, check := range checks {
			if check.condition(containerA, containerB) {
				reasons = append(reasons, fmt.Sprintf(check.reason, description, containerA.Name, index))
			}
		}
	}

	return reasons
}

func compareResources(a *v1.ResourceRequirements, b *v1.ResourceRequirements) bool {
	equal := true
	if a != nil {
		equal = compareResourcesAssumeFirstNotNil(a, b)
	}
	if equal && (b != nil) {
		equal = compareResourcesAssumeFirstNotNil(b, a)
	}

	return equal
}

func compareResourcesAssumeFirstNotNil(a *v1.ResourceRequirements, b *v1.ResourceRequirements) bool {
	if b == nil || (len(b.Requests) == 0) {
		return len(a.Requests) == 0
	}
	for k, v := range a.Requests {
		if (&v).Cmp(b.Requests[k]) != 0 { //nolint: gosec, scopelint
			return false
		}
	}
	for k, v := range a.Limits {
		if (&v).Cmp(b.Limits[k]) != 0 { //nolint: gosec, scopelint
			return false
		}
	}

	return true
}

// mapContains returns true if and only if haystack contains all the keys from the needle with matching corresponding values
func mapContains(haystack, needle map[string]string) bool {
	if len(haystack) < len(needle) {
		return false
	}

	for k, v := range needle {
		v2, ok := haystack[k]
		if !ok || v2 != v {
			return false
		}
	}

	return true
}
