package unique

// Slice returns a new slice with all duplicate values removed.
func Slice[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	unique := make([]T, 0, len(slice))
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}
