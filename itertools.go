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

// DropWhile
// https://docs.python.org/3/library/itertools.html#itertools.dropwhile
func DropWhile[A any](aa []A, pred func(a A) bool) []A {
	i := 0
	for ; i < len(aa); i++ {
		if !pred(aa[i]) {
			break
		}
	}

	return aa[i:]
}

// TakeWhile
// https://docs.python.org/3/library/itertools.html#itertools.takewhile
func TakeWhile[A any](aa []A, pred func(a A) bool) []A {
	i := 0
	for ; i < len(aa); i++ {
		if !pred(aa[i]) {
			break
		}
	}

	return aa[:i]
}
