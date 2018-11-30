package merkleTree

import (
	"github.com/ethereum/go-ethereum/crypto"
)

type MerkleTree struct {
	Levels    [][]MerkleNode
	RootNode  *MerkleNode
}

type MerkleNode struct {
	Left     *MerkleNode
	Right    *MerkleNode
	Segment  Segment
}

type Segment struct {
	SegmentLength int
	Data          []byte
}

var zeroHash = crypto.Keccak256([]byte{})

func NewMerkleNode(left, right *MerkleNode, hashFunc func(data ...[]byte) []byte) *MerkleNode  {
	var node MerkleNode

	if right == nil {
		prevHashes := append(left.Segment.Data, zeroHash...)
		node.Segment.Data = hashFunc(prevHashes)
		node.Segment.SegmentLength = left.Segment.SegmentLength
	} else {
		prevHashes := append(left.Segment.Data, right.Segment.Data...)
		node.Segment.Data = hashFunc(prevHashes)
		node.Segment.SegmentLength = left.Segment.SegmentLength + right.Segment.SegmentLength
	}

	node.Left = left
	node.Right = right

	return &node
}

func LeafToNode(segment Segment, hashFunc func(data ...[]byte) []byte) *MerkleNode {
	node := MerkleNode{
		Segment: Segment{
			segment.SegmentLength,
			hashFunc(segment.Data),
		},
	}
	return &node
}

func NewMerkleTree(segment []Segment, hashFunc func(data ...[]byte) []byte) *MerkleTree {
	var nodes  []MerkleNode
	var levels [][]MerkleNode

	if len(segment)%2 != 0 {
		segment = append(segment, Segment{0, []byte{}})
	}

	for _, s := range segment {
		node := LeafToNode(s, hashFunc)
		nodes = append(nodes, *node)
	}

	countOfDataNodes := len(nodes)
	counterOfLevels := 0
	for countOfDataNodes > 1 {
		if countOfDataNodes%2 == 0 {
			countOfDataNodes =  countOfDataNodes / 2
			counterOfLevels++
		} else {
			countOfDataNodes = (countOfDataNodes + 1) / 2
			counterOfLevels++
		}
	}

	levels = [][]MerkleNode{nodes}

	for i := 0; i < counterOfLevels; i++ {
		var level []MerkleNode

		lastNodeIndex := len(nodes) - 1
		for j := 0; j <= lastNodeIndex; j+=2 {
			if j == lastNodeIndex && lastNodeIndex%2 == 0 {
				node := NewMerkleNode(&nodes[j], nil, crypto.Keccak256)
				level = append(level, *node)
			} else {
				node := NewMerkleNode(&nodes[j], &nodes[j+1], crypto.Keccak256)
				level = append(level, *node)
			}
		}

		nodes = level
		levels = append(levels, level)
	}

	tree := MerkleTree{levels, &nodes[0]}

	return &tree
}