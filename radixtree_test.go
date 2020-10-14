package radixtree

import (
	"fmt"
	"testing"
)

func TestRadixTreePut(t *testing.T) {
	var tests = []struct {
		key           string
		value         interface{}
		errorExpected bool
	}{
		{"abc", 1, false},
		{"", 1, true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s:%s", tt.key, tt.value)
		t.Run(testname, func(t *testing.T) {
			trie := NewRadixTree()
			result := trie.Put(tt.key, tt.value)
			errorResult := result != nil
			if errorResult != tt.errorExpected {
				t.Errorf("Error expected %t but got %t",
					tt.errorExpected, errorResult)
			}
		})

	}
}

func TestRadixTreeGet(t *testing.T) {
	var tests = []struct {
		key           string
		value         interface{}
		queryKey      string
		queryValue    interface{}
		errorExpected bool
	}{
		{"abc", 1, "abc", 1, false},
		{"abc", 1, "", nil, true},
		{"abc", 1, "wrong", nil, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s:%s", tt.key, tt.value)
		t.Run(testname, func(t *testing.T) {
			trie := NewRadixTree()
			trie.Put(tt.key, tt.value)
			val, result := trie.Get(tt.queryKey)
			errorResult := result != nil
			if errorResult != tt.errorExpected {
				t.Errorf("Error expected %t but got %t",
					tt.errorExpected, errorResult)
			}
			if val != tt.queryValue {
				t.Errorf("Expected %v got %v",
					tt.queryValue, val)
			}
		})

	}
}

func TestRadixTreeDelete(t *testing.T) {
	var tests = []struct {
		key           string
		value         interface{}
		deleteKey     string
		foundExpected bool
		errorExpected bool
	}{
		{"abc", 1, "abc", true, false},
		{"abc", 1, "", false, true},
		{"abc", 1, "wrong", false, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s:%s", tt.key, tt.value)
		t.Run(testname, func(t *testing.T) {
			trie := NewRadixTree()
			trie.Put(tt.key, tt.value)
			deleted, error := trie.Delete(tt.deleteKey)
			errorResult := error != nil
			if errorResult != tt.errorExpected {
				t.Errorf("Error expected %t but got %t",
					tt.errorExpected, errorResult)
			}
			if deleted != tt.foundExpected {
				t.Errorf("Expected %v got %v",
					tt.foundExpected, deleted)
			}
		})

	}
}

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

func TestRadixTreeContains(t *testing.T) {
	var tests = []struct {
		key              string
		value            interface{}
		queryKey         string
		expectedContains bool
		errorExpected    bool
	}{
		{"abc", 1, "abc", true, false},
		{"abc", 1, "", false, true},
		{"abc", 1, "wrong", false, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s:%s", tt.key, tt.value)
		t.Run(testname, func(t *testing.T) {
			trie := NewRadixTree()
			trie.Put(tt.key, tt.value)
			contains, err := trie.Contains(tt.queryKey)
			errorResult := err != nil
			if errorResult != tt.errorExpected {
				t.Errorf("Error expected %t but got %t",
					tt.errorExpected, errorResult)
			}
			if contains != tt.expectedContains {
				t.Errorf("Expected %v got %v",
					tt.expectedContains, contains)
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
		{map[string]interface{}{
			"abcd":         1,
			"abcde":        1,
			"www.test.com": 1,
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

func TestDelete(t *testing.T) {
	var tests = []struct {
		keyValues      map[string]interface{}
		deleteKey      string
		expectedDelete bool
		expectedSize   int
	}{
		//{map[string]interface{}{}, "abc", false, 0},
		{map[string]interface{}{
			"abc": 1,
		}, "abc", true, 0},
		{map[string]interface{}{
			"abcd": 1,
		}, "abc", false, 1},
		{map[string]interface{}{
			"abcd":         1,
			"abcde":        1,
			"www.test.com": 1,
		}, "abc", false, 3},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.keyValues)
		t.Run(testname, func(t *testing.T) {
			node := newRadixTreeNode("", nil)
			for key, value := range tt.keyValues {
				node.put(key, value)
			}
			deleted, _ := node.delete(tt.deleteKey)
			if deleted != tt.expectedDelete {
				t.Errorf("Expected delete to return %t for key %s",
					tt.expectedDelete, tt.deleteKey)
			}
			if node.size != tt.expectedSize {
				t.Errorf("Got %d expected size to be %d for key %s", node.size,
					tt.expectedSize, tt.deleteKey)
			}
		})

	}
}

func TestLongestPrefix(t *testing.T) {
	var tests = []struct {
		keyValues         map[string]interface{}
		searchKey, prefix string
	}{
		//{map[string]interface{}{}, "abc", false, 0},
		{map[string]interface{}{
			"abc": 1,
		}, "abc", "abc"},
		{map[string]interface{}{
			"abcd": 1,
		}, "abcdef", "abcd"},
		{map[string]interface{}{
			"abcd":         1,
			"abcde":        1,
			"www.test.com": 1,
		}, "www.test.com/index.html", "www.test.com"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.keyValues)
		t.Run(testname, func(t *testing.T) {
			node := newRadixTreeNode("", nil)
			for key, value := range tt.keyValues {
				node.put(key, value)
			}
			prefix := node.longestPrefixOf(tt.searchKey, 0)
			if *prefix != tt.prefix {
				t.Errorf("Expected prefix to return %s for key %s",
					tt.prefix, *prefix)
			}
		})
	}
}

func TestItems(t *testing.T) {
	var tests = []struct {
		keyValues map[string]interface{}
	}{
		//{map[string]interface{}{}, "abc", false, 0},
		{map[string]interface{}{
			"abc": 1,
		}},
		{map[string]interface{}{
			"abcd": 1,
		}},
		{map[string]interface{}{
			"abcd":         1,
			"abcde":        1,
			"www.test.com": 1,
		}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.keyValues)
		t.Run(testname, func(t *testing.T) {
			node := newRadixTreeNode("", nil)
			for key, value := range tt.keyValues {
				node.put(key, value)
			}

			keyValues := make([]KeyValue, 0)
			for keyValue := range node.items(make([]string, 0)) {
				keyValues = append(keyValues, keyValue)
			}

			if len(keyValues) != len(tt.keyValues) {
				t.Errorf("Expected %d key values got %d", len(tt.keyValues), len(keyValues))
			}
		})
	}
}
