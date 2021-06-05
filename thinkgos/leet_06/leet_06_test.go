package leet_06

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		ss   []int
		want int
	}{
		{
			"",
			[]int{4, 1, 2, 1, 2},
			4,
		},
		{
			"",
			[]int{6, 1, 2, 1, 2, 5, 7, 5, 3, 7, 6},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.ss); got != tt.want {
				t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
