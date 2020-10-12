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
