package itertools

func Filter[A any](aa []A, pred func(a A) bool) []A {
	filtered := []A{}
	for _, a := range aa {
		if pred(a) {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

func Map[A any, B any](aa []A, mapper func(a A) B) []B {
	mapped := []B{}
	for _, a := range aa {
		mapped = append(mapped, mapper(a))
	}
	return mapped
}
