package radixtree

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
	subKey := key[1:]
	linksByFirstChar[key[0]] = subKey
	links[subKey] = newRadixTreeNode("", value)
	return &radixTreeNode{1, nil, links, linksByFirstChar}
}
