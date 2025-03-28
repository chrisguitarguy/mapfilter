package mapfilter

import (
	"iter"
)

func Map[V, U any](seq iter.Seq[V], f func(V) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range seq {
			if !yield(f(v)) {
				break
			}
		}
	}
}

func Map2[K, V, O, U any](seq iter.Seq2[K, V], f func(K, V) (O, U)) iter.Seq2[O, U] {
	return func(yield func(O, U) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				break
			}
		}
	}
}
