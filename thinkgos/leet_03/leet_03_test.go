package leet_03

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

var arr = []Dept{
	{1, 0, "超然科技"},
	{2, 0, "低速科技"},
	{3, 1, "科研中心"},
	{4, 1, "运营中心"},
	{5, 2, "吃喝院"},
	{6, 2, "研究院"},
	{7, 3, "aa"},
	{8, 3, "bb"},
	{9, 4, "cc"},
	{10, 5, "dd"},
	{11, 6, "ee"},
}

func TestTree(t *testing.T) {
	v, err := json.MarshalIndent(toTree(arr, 0), " ", "  ")
	require.NoError(t, err)
	t.Logf(string(v))
}

func BenchmarkTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toTree(arr, 0)
	}
}

func TestTree2(t *testing.T) {
	v, err := json.MarshalIndent(toTree2(arr, 0), " ", "  ")
	require.NoError(t, err)
	t.Logf(string(v))
}

func BenchmarkTree2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toTree2(arr, 0)
	}
}
