package leet_05

import (
	"reflect"
	"testing"
)

func TestUniqueSortList(t *testing.T) {
	tests := []struct {
		name string
		ss   []int
		want []int
	}{
		{
			"",
			[]int{1, 3, 2, 5, 6, 3, 2, 7},
			[]int{1, 2, 3, 5, 6, 7},
		},
		{
			"",
			[]int{1, 1, 3, 2, 5, 5, 6, 3, 2, 7},
			[]int{1, 2, 3, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueSortList(tt.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkUniqueSortList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniqueSortList([]int{1, 1, 3, 2, 5, 5, 6, 3, 2, 7})
	}
}
