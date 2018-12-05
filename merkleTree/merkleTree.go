package merkleTree

import (
	"bytes"
	"encoding/binary"
	"sort"

	"github.com/ethereum/go-ethereum/crypto"
)

const plasmaLength = 16777215

type ProofStep struct {
	Node  *MerkleNode
	Right *MerkleNode
}

type MerkleTree struct {
	Levels   []MerkleNode
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left 		 *MerkleNode
	Right 		 *MerkleNode
	Parent 		 *MerkleNode
	Segment 	 *Segment
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
		Segment: &Segment{
			segmentLength,
			hashFunc(append(UintToBytesArray(segmentLength), segment.Data...)),
		},
	}
}

func NewMerkleNode(left, right *MerkleNode, hashFunc func(data ...[]byte) []byte) *MerkleNode {
	var node MerkleNode

	nodeSegment := getNodeHashAndLength(
		left.Segment.SegmentLength,
		right.Segment.SegmentLength,
		left.Segment.Hash,
		right.Segment.Hash,
		hashFunc,
	)
	node.Segment = nodeSegment

	return &node
}

func NewMerkleTree(segment []InputSegment, hashFunc func(data ...[]byte) []byte) *MerkleTree {
	var nodes []MerkleNode
	var buckets []MerkleNode

	for i := 0; i < len(segment); i+=2 {
		var node1 *MerkleNode
		node1 = MakeLeaf(&segment[i], hashFunc)
		if i != len(segment) - 1 {
			node2 := MakeLeaf(&segment[i+1], hashFunc)
			node2.Left = node1
			node1.Right = node2
			nodes = append(nodes, *node1, *node2)
		} else {
			nodes = append(nodes, *node1)
		}
	}

	buckets = nodes[:]

	for len(buckets) != 1 {
		var level []MerkleNode

		for len(buckets) > 0 {
			if len(buckets) >= 2 {
				node1 := buckets[0]
				node2 := buckets[1]
				buckets = buckets[2:]
				parent := NewMerkleNode(&node1, &node2, hashFunc)
				node1.Parent = parent
				node2.Parent = parent
				node1.Right = &node2
				node2.Left = &node1
				level = append(level, *parent)
			} else {
				level = append(level, buckets[len(buckets)-1])
				buckets = []MerkleNode{}
			}
		}
		buckets = level
	}

	tree := MerkleTree{nodes, &buckets[0]}

	return &tree
}

func PrepareSegments(list []InputSegment) []InputSegment {

	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Start < list[j].Start
	})

	var listWithEmptyStruct []InputSegment

	for i := 0; i <= len(list); i++ {
		// Check for start
		if i == 0 && list[i].Start != 0 {
			startSturct := InputSegment{Start: 0, End: list[i].Start, Data: []byte{}}
			listWithEmptyStruct = append(listWithEmptyStruct, startSturct)
		}

		// Check for end
		if i == len(list)-1 {
			if list[i].End != plasmaLength {
				endS := InputSegment{Start: list[i].End, End: plasmaLength, Data: []byte{}}
				listWithEmptyStruct = append(listWithEmptyStruct, list[i])
				listWithEmptyStruct = append(listWithEmptyStruct, endS)
			} else {
				listWithEmptyStruct = append(listWithEmptyStruct, list[i])
			}
			return listWithEmptyStruct
		}

		el := list[i]
		nextEl := list[i+1]

		if nextEl.Start-el.End > 1 {
			empty := InputSegment{Start: el.End, End: nextEl.Start, Data: []byte{}}
			listWithEmptyStruct = append(listWithEmptyStruct, el)
			listWithEmptyStruct = append(listWithEmptyStruct, empty)
		} else {
			listWithEmptyStruct = append(listWithEmptyStruct, el)
		}

	}
	return listWithEmptyStruct
}

func (tree *MerkleTree) GetProof(index int) []ProofStep {
	var proof []ProofStep
	curr := tree.Levels[index]

	for curr.Parent != nil {
		var node *MerkleNode
		if curr.Right != nil {
			node = curr.Right
		} else {
			node = curr.Left
		}
		proof = append(proof, ProofStep{
			Node: node,
			Right: curr.Right,
		})
		curr = *curr.Parent
	}

	return proof
}

func Verify(proof []ProofStep, rootHash *Segment, leaf *InputSegment, hashFunc func(data ...[]byte) []byte) bool {
	curr := MakeLeaf(leaf, hashFunc).Segment

	for _, step := range proof {
		if step.Right != nil {
			left := append(UintToBytesArray(curr.SegmentLength), curr.Hash...)
			right := append(UintToBytesArray(step.Node.Segment.SegmentLength), step.Node.Segment.Hash...)
			curr = &Segment{curr.SegmentLength + step.Node.Segment.SegmentLength, hashFunc(append(left, right...))}
		} else {
			left := append(UintToBytesArray(step.Node.Segment.SegmentLength),step.Node.Segment.Hash...)
			right := append(UintToBytesArray(curr.SegmentLength), curr.Hash...)
			curr = &Segment{curr.SegmentLength + step.Node.Segment.SegmentLength, hashFunc(append(left, right...))}
		}
	}

	if curr.SegmentLength == rootHash.SegmentLength && bytes.Equal(curr.Hash, rootHash.Hash) {
		return true
	} else {
		return false
	}
}

func UintToBytesArray(value uint32) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint32(b, value)
	return b
}
