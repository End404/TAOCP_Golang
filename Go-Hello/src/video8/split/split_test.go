package split

import (
	"reflect"
	"testing"
)

// 单元测试和基准测试

func TestSplit(t *testing.T) {
	/*
		//测试1
		got := Split("我爱你", "爱")
		want := []string{"爱", "你"}
		//got := Split("a:b:c", ":")
		//want := []string{"a", "b", "c"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want:%v got:%v", want, got)
		}
	*/

	// 测试2 基准测试
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple":     test{input: "我爱你", sep: "爱", want: []string{"爱", "你"}},
		"multi sep":  test{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"multi sep2": test{input: "abcd", sep: "bc", want: []string{"a", "d"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %s falide, want:%v got:%v", name, tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "河")
	}
}
