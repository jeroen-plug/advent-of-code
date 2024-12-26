package day23

type Set map[string]struct{}

func (a Set) Union(b Set) Set {
	result := make(Set)
	for k := range a {
		result[k] = struct{}{}
	}
	for k := range b {
		result[k] = struct{}{}
	}
	return result
}

func (a Set) Intersection(b Set) Set {
	result := make(Set)
	for k := range a {
		if _, exists := b[k]; exists {
			result[k] = struct{}{}
		}
	}
	return result
}

func ToSet(slice []string) Set {
	result := make(Set)
	for _, k := range slice {
		result[k] = struct{}{}
	}
	return result
}
