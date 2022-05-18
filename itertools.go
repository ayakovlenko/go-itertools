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

func findIndex[A any](aa []A, pred func(a A) bool) (i int) {
	for ; i < len(aa); i++ {
		if !pred(aa[i]) {
			break
		}
	}
	return i
}

// DropWhile
// https://docs.python.org/3/library/itertools.html#itertools.dropwhile
func DropWhile[A any](aa []A, pred func(a A) bool) []A {
	i := findIndex(aa, pred)
	return aa[i:]
}

// TakeWhile
// https://docs.python.org/3/library/itertools.html#itertools.takewhile
func TakeWhile[A any](aa []A, pred func(a A) bool) []A {
	i := findIndex(aa, pred)
	return aa[:i]
}

// Pairwise
// https://docs.python.org/3/library/itertools.html#itertools.pairwise
func Pairwise[A any](aa []A) [][2]A {
	pairs := [][2]A{}

	if len(aa) < 2 {
		return pairs
	}

	for i := 0; i < len(aa)-1; i++ {
		pairs = append(pairs, [2]A{aa[i], aa[i+1]})
	}

	return pairs
}

// GroupBy
// https://deno.land/std@0.139.0/collections#groupby
func GroupBy[A any, B comparable](aa []A, grouper func(a A) B) map[B][]A {
	grouped := map[B][]A{}
	for _, a := range aa {
		k := grouper(a)
		if group, ok := grouped[k]; ok {
			grouped[k] = append(group, a)
		} else {
			grouped[k] = []A{a}
		}
	}
	return grouped
}
