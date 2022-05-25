package maps

// FilterEntries
// https://deno.land/std@0.140.0/collections#filterentries
func FilterEntries[K comparable, V any](
	m map[K]V,
	pred func(key K, value V) bool,
) map[K]V {
	filtered := map[K]V{}
	for k, v := range m {
		if pred(k, v) {
			filtered[k] = v
		}
	}
	return filtered
}

// FilterKeys
// https://deno.land/std@0.140.0/collections#filterkeys
// TODO: test
func FilterKeys[K comparable, V any](
	m map[K]V,
	pred func(key K) bool,
) map[K]V {
	filtered := map[K]V{}
	for k, v := range m {
		if pred(k) {
			filtered[k] = v
		}
	}
	return filtered
}

// FilterValues
// https://deno.land/std@0.140.0/collections#filtervalues
// TODO: test
func FilterValues[K comparable, V any](
	m map[K]V,
	pred func(value V) bool,
) map[K]V {
	filtered := map[K]V{}
	for k, v := range m {
		if pred(v) {
			filtered[k] = v
		}
	}
	return filtered
}
