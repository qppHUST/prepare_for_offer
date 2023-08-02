package sort

import (
	"fmt"
	"testing"
)

func TestThreeWayFastDischarge(t *testing.T) {
	nums := []int{0, 2, 2, 2, 0, 2, 1, 1}
	ThreeWayFastDischarge(nums)
	fmt.Printf("%v", nums)
}
