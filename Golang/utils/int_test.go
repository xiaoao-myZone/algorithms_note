package utils

import (
	"reflect"
	"testing"
)

type arr_to_arr struct {
	input  []int
	output []int
}

func Test_IntReverse(t *testing.T) {
	test_case := []arr_to_arr{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
	}
	for _, one := range test_case {
		if res := IntReverse(one.input, true); reflect.DeepEqual(res, one.output) {
			t.Log("通过")
		} else {
			t.Log("不通过")
		}
	}

	for _, one := range test_case {
		if IntReverse(one.input, false); reflect.DeepEqual(one.input, one.output) {
			t.Log("通过")
		} else {
			t.Log("不通过")
		}
	}
}
