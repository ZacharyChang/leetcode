package leetcode

import (
	"fmt"
	"testing"
)

func TestWordFilter(t *testing.T) {
	s0 := Constructor([]string{"apple"})
	s1 := Constructor([]string{"pop"})
	s2 := Constructor([]string{"pro", "professional", "program", "progress", "programmer", "prepare", "player"})
	tests := []struct {
		w      WordFilter
		prefix string
		suffix string
		want   int
	}{
		{s0, "a", "e", 0},
		{s0, "b", "", -1},
		{s1, "", "", 0},
		{s1, "", "p", 0},
		{s1, "", "op", 0},
		{s1, "", "pop", 0},
		{s1, "p", "", 0},
		{s1, "pop", "pop", 0},
		{s1, "pg", "pgp", -1},
		{s2, "pro", "", 4},
		{s2, "prof", "", 1},
		{s2, "prof", "am", -1},
		{s2, "p", "o", 0},
		{s2, "", "r", 6},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if got := tt.w.F(tt.prefix, tt.suffix); got != tt.want {
				t.Errorf("WordFilter.F() = %v, want %v", got, tt.want)
			}
		})
	}
}
