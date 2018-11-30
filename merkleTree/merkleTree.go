package merkleTree

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
)

type MerkleProof struct {
	Tree     [][]MerkleNode
	RootHash []byte
	Segment  Segment
}

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
	SegmentLength uint32
	Data          []byte
}

var zeroHash = crypto.Keccak256([]byte{})
var zeroSegment = uint32(0)

func NewMerkleNode(left, right *MerkleNode, hashFunc func(data ...[]byte) []byte) *MerkleNode  {
	var node MerkleNode

	if right == nil {
		concatLeftNodeData := append(UintToBytesArray(left.Segment.SegmentLength), left.Segment.Data...)
		concatRightNodeData := append(UintToBytesArray(zeroSegment), zeroHash...)
		prevHashes := append(concatLeftNodeData, concatRightNodeData...)
		node.Segment.Data = hashFunc(prevHashes)
		node.Segment.SegmentLength = left.Segment.SegmentLength
	} else {
		concatLeftNodeData := append(UintToBytesArray(left.Segment.SegmentLength), left.Segment.Data...)
		concatRightNodeData := append(UintToBytesArray(right.Segment.SegmentLength), right.Segment.Data...)
		prevHashes := append(concatLeftNodeData, concatRightNodeData...)
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
			hashFunc(append(UintToBytesArray(segment.SegmentLength), segment.Data...)),
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

func (tree *MerkleTree) GetProof(segment Segment) (*MerkleProof, error){
	segmentHash := crypto.Keccak256(append(UintToBytesArray(segment.SegmentLength), segment.Data...))
	exist := false
	leafs := tree.Levels[0]
	for _, l := range leafs {
		if bytes.Equal(l.Segment.Data, segmentHash) && l.Segment.SegmentLength == segment.SegmentLength {
			exist = true
			break
		}
	}

	if exist == true {
		return &MerkleProof{
			 tree.Levels,
			 tree.RootNode.Segment.Data,
			 segment,
		}, nil
	} else {
		return nil, errors.New("Segment does not belong to the Merkle Tree")
	}
}

func UintToBytesArray(value uint32) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint32(b, value)
	return b
}