package merkleTree

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/crypto"
)

//type MerkleProof struct {
//	Tree       [][]MerkleNode
//	RootHash   []byte
//	RootLength uint32
//	Segment    Segment
//}

type MerkleTree struct {
	Levels   [][]MerkleNode
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left    *MerkleNode
	Right   *MerkleNode
	Segment *Segment
}

type InputSegment struct {
	Start uint32
	End   uint32
	Data  []byte
}

type Segment struct {
	SegmentLength uint32
	Hash          []byte
}

var zeroHash = crypto.Keccak256([]byte{})
var zeroSegment = uint32(0)

func getNodeHashAndLength(leftLength uint32, rightLength uint32, leftHash []byte, rightHash []byte, hashFunc func(data ...[]byte) []byte) *Segment {
	segmentLength := leftLength + rightLength
	leftData := append(UintToBytesArray(leftLength), leftHash...)
	rightData := append(UintToBytesArray(rightLength), rightHash...)
	return &Segment{
		segmentLength,
		hashFunc(append(leftData, rightData...)),
	}
}

func MakeLeaf(segment *InputSegment, hashFunc func(data ...[]byte) []byte) *MerkleNode {
	segmentLength := segment.End - segment.Start
	return &MerkleNode{
		nil,
		nil,
		&Segment{
			segmentLength,
			hashFunc(append(UintToBytesArray(segmentLength), segment.Data...)),
		},
	}
}

func NewMerkleNode(left, right *MerkleNode, hashFunc func(data ...[]byte) []byte) *MerkleNode  {
	var node MerkleNode

	nodeSegment := getNodeHashAndLength(
		left.Segment.SegmentLength,
		right.Segment.SegmentLength,
		left.Segment.Hash,
		right.Segment.Hash,
		hashFunc,
	)
	node.Segment = nodeSegment

	node.Left = left
	node.Right = right

	return &node
}

func NewMerkleTree(segment []InputSegment, hashFunc func(data ...[]byte) []byte) *MerkleTree {
	var nodes []MerkleNode
	var levels [][]MerkleNode
	var notBalancedNodes []MerkleNode

	for _, s := range segment {
		node := MakeLeaf(&s, hashFunc)
		nodes = append(nodes, *node)
	}

	if len(nodes)%2 != 0 {
		notBalancedNodes = append(notBalancedNodes, nodes[len(nodes)-1])
		nodes = nodes[:len(nodes)-1]
	}

	levels = [][]MerkleNode{nodes}

	for len(nodes) > 1 {
		var level []MerkleNode

		lastNodeIndex := len(nodes) - 1
		for j := 0; j <= lastNodeIndex; j += 2 {
			if j == lastNodeIndex && j%2 == 0 {
				notBalancedNodes = append([]MerkleNode{nodes[j]}, notBalancedNodes...)
			} else {
				node := NewMerkleNode(&nodes[j], &nodes[j+1], hashFunc)
				level = append(level, *node)
			}
		}

		nodes = level
		levels = append(levels, level)

		if len(nodes)%2 != 0 && len(notBalancedNodes) != 0 {
			nodes = append(nodes, notBalancedNodes[0])
			notBalancedNodes = notBalancedNodes[1:]
		}
	}

	tree := MerkleTree{levels, &nodes[0]}

	return &tree
}

//func (tree *MerkleTree) GetProof(segment Segment) (*MerkleProof, error){
//	segmentHash := crypto.Keccak256(append(UintToBytesArray(segment.SegmentLength), segment.Data...))
//	exist := false
//	leafs := tree.Levels[0]
//	for _, l := range leafs {
//		if bytes.Equal(l.Segment.Data, segmentHash) && l.Segment.SegmentLength == segment.SegmentLength {
//			exist = true
//			break
//		}
//	}
//
//	if exist == true {
//		return &MerkleProof{
//			 tree.Levels,
//			 tree.RootNode.Segment.Data,
//			 tree.RootNode.Segment.SegmentLength,
//			 segment,
//		}, nil
//	} else {
//		return nil, errors.New("Segment does not belong to the Merkle Tree")
//	}
//}
//
//func Verify(proof *MerkleProof, rootHash []byte) bool {
//	tree := proof.Tree
//	merkleRoot := proof.RootHash
//	leafs := tree[0]
//	var nodes []Segment
//	for _, l := range leafs {
//		nodes = append(nodes, Segment{l.Segment.SegmentLength, l.Segment.Data})
//	}
//
//	for len(nodes) > 1 {
//		var level []Segment
//
//		if len(nodes)%2 != 0 {
//			nodes = append(nodes, Segment{zeroSegment, zeroHash})
//		}
//
//		for i := 0; i < len(nodes); i+=2 {
//			dataLeft := nodes[i].Data
//			dataRight := nodes[i+1].Data
//			segmentLengthLeft := nodes[i].SegmentLength
//			segmentLengthRight := nodes[i+1].SegmentLength
//			currentSegmentLength := segmentLengthLeft + segmentLengthRight
//
//			leftSegment := append(UintToBytesArray(segmentLengthLeft), dataLeft...)
//			rightSegment := append(UintToBytesArray(segmentLengthRight), dataRight...)
//
//			node := crypto.Keccak256(append(leftSegment, rightSegment...))
//			level = append(level, Segment{currentSegmentLength, node})
//		}
//
//		nodes = level
//	}
//
//	if bytes.Equal(nodes[0].Data, merkleRoot) && bytes.Equal(rootHash, merkleRoot) && nodes[0].SegmentLength == proof.RootLength {
//		return true
//	} else {
//		return false
//	}
//}

func UintToBytesArray(value uint32) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint32(b, value)
	return b
}