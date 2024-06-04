package leet_03

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

var arr = []*Dept{
	{1, 0, "超然科技", nil},
	{2, 0, "低速科技", nil},
	{3, 1, "科研中心", nil},
	{4, 1, "运营中心", nil},
	{5, 2, "吃喝院", nil},
	{6, 2, "研究院", nil},
	{7, 3, "aa", nil},
	{8, 3, "bb", nil},
	{9, 4, "cc", nil},
	{10, 5, "dd", nil},
	{11, 6, "ee", nil},
}

func TestTree(t *testing.T) {
	v, err := json.MarshalIndent(IntoTree(arr, 0), " ", "  ")
	require.NoError(t, err)
	t.Logf(string(v))

	
}

func BenchmarkTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntoTree(arr, 0)
	}
}
