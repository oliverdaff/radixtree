// Package radixtree a data structure that represents a space-optimized trie
// (prefix tree) in which each node that is the only child is merged with its parent.
package radixtree

import (
	"errors"
	"sort"
	"strings"
)

// RadixTree is a external API for the trie.
type RadixTree struct {
	root *radixTreeNode
}

// NewRadixTree creates a new RadixTree
func NewRadixTree() *RadixTree {
	return &RadixTree{newRadixTreeNode("", nil)}
}

// Put stores a key value in the trie.
// Returns a error if the key has zero length.
func (rt *RadixTree) Put(key string, value interface{}) error {
	if len(key) == 0 {
		return errors.New("Zero length key")
	}
	rt.root.put(key, value)
	return nil
}

// Get returns the value associated with the key
// else returns nil.
// Get returns an error if the key passed has zero length.
func (rt *RadixTree) Get(key string) (interface{}, error) {
	if len(key) == 0 {
		return nil, errors.New("Zero length key")
	}
	return rt.root.get(key), nil
}

// Delete removes the key and value from the tree
// returns true if the key was found else returns false.
// A error is returned if the key has zero length.
func (rt *RadixTree) Delete(key string) (bool, error) {
	if len(key) == 0 {
		return false, errors.New("Zero length key")
	}
	deleted, _ := rt.root.delete(key)
	return deleted, nil
}

// Contains returns true if the trie contains the key
// else returns false.
// A error is returned if the key has zero length.
func (rt *RadixTree) Contains(key string) (bool, error) {
	if len(key) == 0 {
		return false, errors.New("Zero length key")
	}
	return rt.root.contains(key), nil
}

// IsEmpty returns true if the trie is empty.
func (rt *RadixTree) IsEmpty() bool {
	return rt.root.size == 0
}

// LongestPrefixOf returns the longest prefix of the key
// found in the trie.
// A empty string is returned no common prefix is found.
// A error is returned if the passed key has zero length.
func (rt *RadixTree) LongestPrefixOf(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("Zero length key")
	}
	if prefix := rt.root.longestPrefixOf(key, 0); prefix != nil {
		return *prefix, nil
	}
	return "", nil
}

// KeysWithPrefix searches the tree for all the
// keys for which s is a valid prefix.
// Returns a channel that returns all the keys
func (rt *RadixTree) KeysWithPrefix(s string) <-chan string {

	prefixNode, path := rt.root.getNodeForPrefix(s, make([]string, 0))
	if prefixNode != nil {
		return prefixNode.keys(path)
	}
	ch := make(chan string)
	close(ch)
	return ch
}

// Keys returns a channel that receives all keys
// in the trie.
func (rt *RadixTree) Keys() <-chan string {
	return rt.root.keys(make([]string, 0))
}

// Items returns a channel that receives all
// key value pairs in the trie.
func (rt *RadixTree) Items() <-chan KeyValue {
	return rt.root.items(make([]string, 0))
}

type radixTreeNode struct {
	size             int
	value            interface{}
	links            map[string]*radixTreeNode
	linksByFirstChar map[byte]string
}

func newRadixTreeNode(key string, value interface{}) *radixTreeNode {
	links := make(map[string]*radixTreeNode)
	linksByFirstChar := make(map[byte]string)
	if len(key) == 0 {
		return &radixTreeNode{0, value, links, linksByFirstChar}
	}
	linksByFirstChar[key[0]] = key
	links[key] = newRadixTreeNode("", value)
	return &radixTreeNode{1, nil, links, linksByFirstChar}
}

func (tn *radixTreeNode) put(key string, value interface{}) (isNewKey bool) {
	isNewKey = false
	if len(key) == 0 { // store the value in this node
		isNewKey = tn.value == nil
		tn.value = value
		return
	}
	next := key[0]
	if link, ok := tn.linksByFirstChar[next]; ok { // if first char is in linksByFirstChar
		commonPrefix := longestCommonPrefix(key, link)
		if link == commonPrefix { // is current link the common prefix
			isNewKey = tn.links[link].put(key[len(commonPrefix):], value)
		} else {
			isNewKey = true
			bridgeNode := tn.links[link].createBridge(link[len(commonPrefix):]) //create a bridge node after common
			delete(tn.links, link)                                              //remove old link
			tn.links[commonPrefix] = bridgeNode                                 //Set bridge to be common prefix
			tn.linksByFirstChar[next] = commonPrefix                            //Set linksByFirstChar as common prefix
			bridgeNode.put(key[len(commonPrefix):], value)                      // put different component after bridge node
		}
		if isNewKey {
			tn.size++
		}
	} else {
		tn.size++ // Save the link with node
		tn.linksByFirstChar[next] = key
		tn.links[key] = newRadixTreeNode("", value)
		isNewKey = true
	}

	return
}

func (tn *radixTreeNode) createBridge(subKey string) *radixTreeNode {
	bridgeNode := newRadixTreeNode("", nil)
	bridgeNode.initBridgeLinks(subKey, tn)
	bridgeNode.size = 1 + tn.size
	return bridgeNode
}

func (tn *radixTreeNode) initBridgeLinks(key string, node *radixTreeNode) {
	tn.links = make(map[string]*radixTreeNode)
	tn.linksByFirstChar = make(map[byte]string)
	tn.linksByFirstChar[key[0]] = key
	tn.links[key] = node
}

func (tn *radixTreeNode) get(key string) interface{} {
	if node := tn.getNode(key); node != nil {
		return node.value
	}
	return nil
}

func (tn *radixTreeNode) getNode(key string) *radixTreeNode {
	if len(key) == 0 {
		return tn
	}
	if link, ok := tn.linksByFirstChar[key[0]]; ok {
		commonPrefix := longestCommonPrefix(key, link)
		if commonPrefix == link {
			return tn.links[link].getNode(key[len(link):])
		}
	}
	return nil
}

func (tn *radixTreeNode) getNodeForPrefix(s string, path []string) (*radixTreeNode, []string) {
	if len(s) == 0 {
		return tn, path
	}
	next := s[0]
	if link, ok := tn.linksByFirstChar[next]; ok {
		commonPrefix := longestCommonPrefix(s, link)
		if commonPrefix == link {
			path = append(path, link)
			return tn.links[link].getNodeForPrefix(s[len(link):], path)
		} else if len(commonPrefix) == len(s) {
			path := append(path, link)
			return tn.links[link], path
		}

	}
	return nil, path
}

func (tn *radixTreeNode) contains(key string) bool {
	return tn.getNode(key) != nil
}

func (tn *radixTreeNode) delete(key string) (bool, bool) {
	deleted, empty := false, false
	if len(key) == 0 {
		deleted = tn.value != nil
		if deleted {
			tn.value = nil
			if tn.size == 0 {
				empty = true
			}
		}
	} else {
		next := key[0]
		if link, ok := tn.linksByFirstChar[next]; ok {
			commonPrefix := longestCommonPrefix(key, link)
			if link == commonPrefix {
				deleted, empty = tn.links[link].delete(key[len(commonPrefix):])
				if deleted {
					tn.size--
				}
				if empty {
					delete(tn.links, link)
					delete(tn.linksByFirstChar, next)
					empty = tn.size == 0 && tn.value == nil
				}
			}
		}
	}
	return deleted, empty
}

func (tn *radixTreeNode) longestPrefixOf(s string, index int) *string {
	var result *string
	if index == len(s) {
		if tn.value != nil {
			result = &s
		}
	} else {
		next := s[index]
		if link, ok := tn.linksByFirstChar[next]; ok {
			commonPrefix := longestCommonPrefix(s[index:], link)
			if commonPrefix == link {
				result = tn.links[link].longestPrefixOf(s, index+len(link))
			}
		}
		if result == nil && tn.value != nil {
			subS := s[:index]
			result = &subS
		}
	}
	return result
}

// KeyValue is a key value pair
// from the RadixTree.
type KeyValue struct {
	Key   string
	Value interface{}
}

func (tn *radixTreeNode) items(path []string) <-chan KeyValue {
	ch := make(chan KeyValue, 1)
	go func() {
		if tn.value != nil {
			ch <- KeyValue{strings.Join(path, ""), tn.value}
		}
		keys := make([]string, 0)
		for key := range tn.links {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			path = append(path, key)
			for item := range tn.links[key].items(path) {
				ch <- item
			}
			path = path[:len(path)-1]
		}
		close(ch)
	}()
	return ch
}

func (tn *radixTreeNode) keys(path []string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		for keyValue := range tn.items(path) {
			ch <- keyValue.Key
		}
		close(ch)
	}()
	return ch
}

func longestCommonPrefix(key string, link string) string {
	i := 0
	n := min(len(key), len(link))
	for i < n && key[i] == link[i] {
		i++
	}
	return key[:i]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
