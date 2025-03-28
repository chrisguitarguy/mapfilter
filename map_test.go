package mapfilter_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/chrisguitarguy/mapfilter"
)

func TestMap_ReturnsNewSeqWithValues(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3, 4, 5})

	result := mapfilter.Map(seq, func(i int) string {
		return fmt.Sprintf("%d!", i)
	})

	realized := slices.Collect(result)

	if !slices.Equal([]string{"1!", "2!", "3!", "4!", "5!"}, realized) {
		t.Errorf("expected [1!, 2!, 3!, 4!, 5!], got %#v", realized)
	}
}

func TestMap_EarlyReturnSmokeTest(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3, 4, 5})

	result := mapfilter.Map(seq, func(i int) string {
		return fmt.Sprintf("%d!", i)
	})

	for range result {
		break
	}
}

func TestMap2_ReturnsNewSeqWithValues(t *testing.T) {
	seq := slices.All([]int{1, 2, 3, 4, 5})

	result := mapfilter.Map2(seq, func(k, v int) (int, string) {
		return k, fmt.Sprintf("%d: %d!", k, v)
	})

	realized := slices.Collect(mapfilter.ToSeq(result))

	if !slices.Equal([]string{"0: 1!", "1: 2!", "2: 3!", "3: 4!", "4: 5!"}, realized) {
		t.Errorf("expected [0: 1!, 1: 2!, 2: 3!, 3: 4!, 4: 5!], got %#v", realized)
	}
}

func TestMap2_EarlyReturnSmokeTest(t *testing.T) {
	seq := slices.All([]int{1, 2, 3, 4, 5})

	result := mapfilter.Map2(seq, func(k, v int) (int, string) {
		return k, fmt.Sprintf("%d: %d!", k, v)
	})

	for range result {
		break
	}
}
