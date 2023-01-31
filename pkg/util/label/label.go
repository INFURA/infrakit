package label

type (
	Labels      = map[string]string
	Annotations = map[string]string
)

// Union returns a union of all label sets passed in.
func Union(labelSets ...Labels) Labels {
	union := make(Labels)

	for _, labelSet := range labelSets {
		for key := range labelSet {
			union[key] = labelSet[key]
		}
	}

	return union
}
