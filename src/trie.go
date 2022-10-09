package src

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

type Trie struct {
}

type Node struct {
	value MyData
}

func (n Node) Hash() string {
	hash := sha3.NewLegacyKeccak256()
	var buf []byte
	hash.Write([]byte(n.value))
	buf = hash.Sum(nil)
	return hex.EncodeToString(buf)
}

func NewNode() Node {
	return Node{}
}

func NewNodeFromHex(hex string) Node {
	deserialize, err := Deserialize(hex)
	if err != nil {
		return Node{}
	}
	return Node{value: *deserialize}
}

type Path struct {
}

func (t Trie) Update(node *Node, path string, value string) error {
	curNode := node
	if curNode == nil {
		*curNode = NewNode()
	}

	if path == "" {

	}

	return nil
}
