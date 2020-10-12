package radixtree

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		first, second, prefix string
	}{
		{"abc", "ade", "a"},
		{"xyz", "abc", ""},
		{"", "abc", ""},
		{"abc", "", ""},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s:%s", tt.first, tt.second)
		t.Run(testname, func(t *testing.T) {
			prefix := longestCommonPrefix(tt.first, tt.second)
			if prefix != tt.prefix {
				t.Errorf("Expected prefix %s got %s", tt.prefix, prefix)
			}
		})
	}
}
