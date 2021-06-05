package leet_04

import (
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    string
		wantErr bool
	}{
		{
			"",
			"2006年6月3日",
			"2006-06-03",
			false,
		},
		{
			"",
			"2006年10月13日",
			"2006-10-13",
			false,
		},
		{
			"",
			"2006年12月3日",
			"2006-12-03",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Format(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Format("2006年12月3日")
	}
}

func BenchmarkFormat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Format2("2006年12月3日")
	}
}
