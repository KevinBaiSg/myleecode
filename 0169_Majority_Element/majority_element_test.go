package majorityelement

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	nums := []int{1, 2, 5, 3, 2, 2, 8, 2, 2, 5, 2, 3, 2}
	num := majorityElement(nums)
	fmt.Printf("num = %d\n", num)
}
