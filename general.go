package gomaps

// checks if a map is empty
func IsEmpty[M ~map[K]V, K comparable, V any](m *M) bool {
	return len(*m) == 0
}

// check if two maps are equal
func IsEqual[M ~map[K]V, K comparable, V comparable](m1, m2 *M) bool {
	if len(*m1) != len(*m2) {
		return false
	}
	for k, v1 := range *m1 {
		if v2, ok := (*m2)[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// check if two maps are equal given a custom comparator function
func IsEqualFunc[M ~map[K]V, K comparable, V any](m1, m2 *M, f func(V, V) bool) bool {
	if len(*m1) != len(*m2) {
		return false
	}
	for k, v1 := range *m1 {
		if v2, ok := (*m2)[k]; !ok || !f(v1, v2) {
			return false
		}
	}
	return true
}

// clear a map
func Clear[M ~map[K]V, K comparable, V any](m *M) {
	for k := range *m {
		delete(*m, k)
	}
}

// delete key, value pair from map given function
func DeleteFunc[M ~map[K]V, K comparable, V any](m *M, f func(K) bool) {
	for k := range *m {
		if f(k) {
			delete(*m, k)
		}
	}
}

// get the keys as a slice
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// get the values as a slice
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
