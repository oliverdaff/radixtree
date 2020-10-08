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

func (tn *radixTreeNode) put(key string, value interface{}) (isNewKey bool) {
	isNewKey = false
	if len(key) == 0 {
		isNewKey = tn.value == nil
		tn.value = value
		return
	}
	next := key[0]
	if link, ok := tn.linksByFirstChar[next]; ok {
		commonPrefix := longestCommonPrefix(key, link)
		if link == commonPrefix {
			isNewKey = tn.links[link].put(key[len(commonPrefix):], value)
		} else {
			isNewKey = true
			bridgeNode := tn.links[link].createBridge(link[len(commonPrefix):])
			delete(tn.links, link)
			tn.links[commonPrefix] = bridgeNode
			tn.linksByFirstChar[next] = commonPrefix
			bridgeNode.put(key[len(commonPrefix):], value)
		}
		if isNewKey {
			tn.size++
		}
	} else {
		tn.size++
		tn.linksByFirstChar[next] = key[1:]
		tn.links[key[1:]] = newRadixTreeNode("", value)
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
