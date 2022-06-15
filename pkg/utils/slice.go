package utils

func SliceContains[T comparable](sl []T, x T) bool {
	for _, s := range sl {
		if s == x {
			return true
		}
	}
	return false
}
