package array

import "testing"
import "github.com/hemuhan/utils/array"

func TestIndex(t *testing.T) {
	arr := [4]int{1, 2, 3, 4}
	n := array.Index(arr, 3)
	if n != 2 {
		t.Errorf("n is not 2, n value is :%d", n)
	}

	sarr := [3]string{"a", "b", "c"}
	sn := array.Index(sarr, "b")
	if sn != 1 {
		t.Errorf("sn is not 1, sn value is :%d", n)
	}

}
