package radixtree

import (
	"fmt"
	"testing"
)

func TestNewRadixTreeNode(t *testing.T) {
	var tests = []struct {
		key   string
		value interface{}
	}{
		{"abc", 1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s:%s", tt.key, tt.value)
		t.Run(testname, func(t *testing.T) {
			node := newRadixTreeNode(tt.key, tt.value)
			if node.size != 1 {
				t.Errorf("Expected size of 1 got %d", node.size)
			}
			if node.linksByFirstChar[tt.key[0]] != tt.key {
				t.Errorf("Expected links by first char %s to be %s", string(tt.key[0]), tt.key)
			}
		})

	}
}

func TestRadixTreeNodePut(t *testing.T) {
	var tests = []struct {
		items map[string]interface{}
	}{
		{map[string]interface{}{"abc": 1}},
		{map[string]interface{}{
			"abc": 1,
			"ab":  2,
		}},
		{map[string]interface{}{
			"abc":   1,
			"abcde": 2,
			"ab":    3,
		}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.items)
		t.Run(testname, func(t *testing.T) {
			node := newRadixTreeNode("", nil)
			for k, v := range tt.items {
				node.put(k, v)
			}
			for k, v := range tt.items {
				actual := node.get(k)
				if actual != v {
					t.Errorf("Expected value %s for key %s, got %s",
						v, k, actual)
				}
			}
		})

	}
}

func TestRadixTreeNodeGetNodeForPrefix(t *testing.T) {
	var tests = []struct {
		items        map[string]interface{}
		searchKey    string
		nodeExpected bool
	}{
		{map[string]interface{}{"abc": 1}, "ab", true},
		{map[string]interface{}{
			"abc": 1,
			"ab":  2,
		}, "cde", false},
		{map[string]interface{}{
			"abc":   1,
			"abcde": 2,
			"xyz":   2,
		}, "abcd", true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.items)
		t.Run(testname, func(t *testing.T) {
			node := newRadixTreeNode("", nil)
			for k, v := range tt.items {
				node.put(k, v)
			}
			result, _ := node.getNodeForPrefix(tt.searchKey, make([]string, 0))
			if (result != nil) != tt.nodeExpected {
				t.Errorf("Expected %t got %v", tt.nodeExpected, result)
			}
		})

	}
}

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

func TestContains(t *testing.T) {
	var tests = []struct {
		keyValues        map[string]interface{}
		searchKey        string
		expectedContains bool
	}{
		{map[string]interface{}{}, "abc", false},
		{map[string]interface{}{
			"abc": 1,
		}, "abc", true},
		{map[string]interface{}{
			"abcd": 1,
		}, "abc", false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.keyValues)
		t.Run(testname, func(t *testing.T) {
			node := newRadixTreeNode("", nil)
			for key, value := range tt.keyValues {
				node.put(key, value)
			}
			if node.contains(tt.searchKey) != tt.expectedContains {
				t.Errorf("Expected contains to return %t for key %s",
					tt.expectedContains, tt.searchKey)
			}
		})

	}
}
