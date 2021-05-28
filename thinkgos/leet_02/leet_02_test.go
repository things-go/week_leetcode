package leet_02

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDifference(t *testing.T) {
	var s1 = []int{1, 2, 3, 4, 5}
	var s2 = []int{2, 3, 7, 8}

	var wantAdd = []int{7, 8}
	var wantDel = []int{1, 4, 5}

	ss1 := New(s1...)
	ss2 := New(s2...)

	gotDel := ss1.Difference(ss2).List()
	gotAdd := ss2.Difference(ss1).List()

	require.Equal(t, gotAdd, wantAdd)
	require.Equal(t, gotDel, wantDel)
}
