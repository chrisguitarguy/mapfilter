package mapfilter_test

import (
	"slices"
	"testing"

	"github.com/chrisguitarguy/mapfilter"
)

func TestFilter_ReturnsSequenceWithOnlyIncludedValues(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3, 4, 5})

	result := mapfilter.Filter(seq, func(i int) bool {
		return i%2 == 0
	})

	realized := slices.Collect(result)

	if !slices.Equal([]int{2, 4}, realized) {
		t.Errorf("expected [2, 4], got %#v", realized)
	}
}

func TestFilter2_ReturnsSequenceWithOnlyIncludedValues(t *testing.T) {
	seq := slices.All([]int{1, 2, 3, 4, 5})

	result := mapfilter.Filter2(seq, func(k, v int) bool {
		return v%2 == 0
	})

	realized := slices.Collect(mapfilter.ToSeq(result))

	if !slices.Equal([]int{2, 4}, realized) {
		t.Errorf("expected [2, 4], got %#v", realized)
	}
}
