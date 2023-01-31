package podtemplateclientset

import (
	"fmt"

	"github.com/INFURA/infrakit/pkg/util/label"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	KeyInfuraClientSet = "infura.io/ptcs"
)

// SubresourceLabels returns the default labels that should be applied
// to all dependent subresources that this PodTemplateClientSet owns.
func SubresourceLabels(clientSet metav1.Object) label.Labels {
	return label.Labels{
		"app":              clientSet.GetName(),
		KeyInfuraClientSet: fmt.Sprintf("%s.%s", clientSet.GetNamespace(), clientSet.GetName()),
	}
}

// SubresourceLabelsSelector returns the default labels that should be
// applied to all dependent subresources that this PodTemplateClientSet owns,
// as a label selector.
func ClientSetSubresourceLabelsSelector(clientSet metav1.Object) (labels.Selector, error) { //nolint: ireturn
	labels := SubresourceLabels(clientSet)
	labelSelector := &metav1.LabelSelector{}

	for key, value := range labels {
		metav1.AddLabelToSelector(labelSelector, key, value)
	}

	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return selector, nil
}
