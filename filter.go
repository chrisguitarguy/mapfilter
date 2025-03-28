package mapfilter

import (
	"iter"
)

func Filter[V any](seq iter.Seq[V], check func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !check(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}

func Filter2[K, V any](seq iter.Seq2[K, V], check func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !check(k, v) {
				continue
			}

			if !yield(k, v) {
				return
			}
		}
	}
}
