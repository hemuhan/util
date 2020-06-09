package array

import (
	"reflect"
)

/**
查找数组是否包含元素
example
a := []int{1,2,3,4}
n := Index(a,4)
// n is 3

n2 := Index(a,10)
// n2 is -1
**/

func Index(arr interface{}, item interface{}) int {
	ref := reflect.ValueOf(arr)
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice, reflect.Array:
		for n := 0; n < reflect.ValueOf(arr).Len(); n++ {
			if ref.Index(n).Interface() == item {
				return n
			}
		}
	}
	return -1
}
