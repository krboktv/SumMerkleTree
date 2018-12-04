package main

import (
	"bytes"
	"fmt"
	"reflect"

	"../merkleTree"
	"github.com/ethereum/go-ethereum/crypto"
)

var zeroHash = crypto.Keccak256([]byte{})

func main() {
	fmt.Print(test_leaf_to_node1())
	fmt.Print(test_leaf_to_node2())
	fmt.Print(test_New_Merkle_Node_with_left_and_right_nodes())
	fmt.Print(test_3_leafs())
	fmt.Print(test_4_leafs())
	fmt.Print(test_5_leafs())
	fmt.Print(test_6_leafs())
	fmt.Print(test_7_leafs())
	fmt.Print(test_15_leafs())
	fmt.Print(TestSortSegments())
	fmt.Print(test_tree_with_sort())
	//fmt.Print(test_get_proof())
	//fmt.Print(test_verify_proof())
}

func test_leaf_to_node1() string {
	d1 := []byte("Лист")
	segmentStart := uint32(5)
	segmentEnd := uint32(10)
	segmentLength1 := uint32(segmentEnd - segmentStart)

	dataHash := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...))

	segment := merkleTree.InputSegment{segmentStart, segmentEnd, d1}
	hash := merkleTree.MakeLeaf(&segment, crypto.Keccak256)
	if bytes.Equal(dataHash, hash.Segment.Hash) && hash.Segment.SegmentLength == segmentLength1 {
		return "test_leaf_to_node1: true\n"
	} else {
		return "test_leaf_to_node1: false\n"
	}
}

func test_leaf_to_node2() string {
	d1 := []byte("Лист1")
	d2 := []byte("Лист2")
	d3 := []byte("Лист3")
	d4 := []byte("Лист4")
	d5 := []byte("Лист5")
	d6 := []byte{}

	segmentStart1 := uint32(0)
	segmentEnd1 := uint32(2)
	segmentLength1 := uint32(segmentEnd1 - segmentStart1)
	segmentStart2 := uint32(2)
	segmentEnd2 := uint32(6)
	segmentLength2 := uint32(segmentEnd2 - segmentStart2)
	segmentStart3 := uint32(6)
	segmentEnd3 := uint32(10)
	segmentLength3 := uint32(segmentEnd3 - segmentStart3)
	segmentStart4 := uint32(10)
	segmentEnd4 := uint32(13)
	segmentLength4 := uint32(segmentEnd4 - segmentStart4)
	segmentStart5 := uint32(13)
	segmentEnd5 := uint32(15)
	segmentLength5 := uint32(segmentEnd5 - segmentStart5)
	segmentStart6 := uint32(15)
	segmentEnd6 := uint32(20)
	segmentLength6 := uint32(segmentEnd6 - segmentStart6)

	nodes := [][]byte{
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...)),
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength2), d2...)),
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength3), d3...)),
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength4), d4...)),
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength5), d5...)),
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength6), d6...)),
	}

	segments := []merkleTree.InputSegment{
		{segmentStart1, segmentEnd1, d1},
		{segmentStart2, segmentEnd2, d2},
		{segmentStart3, segmentEnd3, d3},
		{segmentStart4, segmentEnd4, d4},
		{segmentStart5, segmentEnd5, d5},
	}

	if len(segments)%2 != 0 {
		segments = append(segments, merkleTree.InputSegment{segmentStart6, segmentEnd6, []byte{}})
	}

	var check = true
	for i, s := range segments {
		hash := merkleTree.MakeLeaf(&s, crypto.Keccak256)
		if !bytes.Equal(hash.Segment.Hash, nodes[i]) {
			check = false
			break
		}
	}

	if check == true {
		return "test_leaf_to_node2: true\n"
	} else {
		return "test_leaf_to_node2: false\n"
	}
}

func test_New_Merkle_Node_with_left_and_right_nodes() string {
	d1 := []byte("Лист1")
	d2 := []byte("Лист2")

	someSegmentLenght1 := uint32(1)
	someSegmentLenght2 := uint32(1)

	node1 := crypto.Keccak256(append(merkleTree.UintToBytesArray(someSegmentLenght1), d1...))
	node2 := crypto.Keccak256(append(merkleTree.UintToBytesArray(someSegmentLenght2), d2...))

	node12 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(someSegmentLenght1), node1...), append(merkleTree.UintToBytesArray(someSegmentLenght2), node2...)...))

	leftNode := merkleTree.MerkleNode{nil, nil, &merkleTree.Segment{someSegmentLenght1, node1}}
	rightNode := merkleTree.MerkleNode{nil, nil, &merkleTree.Segment{someSegmentLenght2, node2}}

	newNode := merkleTree.MerkleNode{&leftNode, &rightNode, &merkleTree.Segment{2, node12}}

	node := merkleTree.NewMerkleNode(&leftNode, &rightNode, crypto.Keccak256)
	if reflect.DeepEqual(node, &newNode) {
		return "test_New_Merkle_Node_with_left_and_right_nodes: true\n"
	} else {
		return "test_New_Merkle_Node_with_left_and_right_nodes: false\n"
	}
}

func test_3_leafs() string {
	d1 := []byte("Есть")
	d2 := []byte("3")
	d3 := []byte("Листа")

	segmentStart1 := uint32(0)
	segmentEnd1 := uint32(2)
	segmentLength1 := uint32(segmentEnd1 - segmentStart1)
	segmentStart2 := uint32(2)
	segmentEnd2 := uint32(6)
	segmentLength2 := uint32(segmentEnd2 - segmentStart2)
	segmentStart3 := uint32(6)
	segmentEnd3 := uint32(10)
	segmentLength3 := uint32(segmentEnd3 - segmentStart3)

	segmentLength12 := segmentLength1 + segmentLength2

	rootSegmentLength := segmentLength12 + segmentLength3

	node1 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...))
	node2 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength2), d2...))
	node3 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength3), d3...))

	node12 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1), node1...), append(merkleTree.UintToBytesArray(segmentLength2), node2...)...))
	rootHash := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength12), node12...), append(merkleTree.UintToBytesArray(segmentLength3), node3...)...))

	segments := []merkleTree.InputSegment{
		{segmentStart1, segmentEnd1, d1},
		{segmentStart2, segmentEnd2, d2},
		{segmentStart3, segmentEnd3, d3},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if bytes.Equal(rootHash, tree.RootNode.Segment.Hash) && tree.RootNode.Segment.SegmentLength == rootSegmentLength {
		return "test_3_leafs: true\n"
	} else {
		return "test_3_leafs: false\n"
	}
}

func test_4_leafs() string {
	d1 := []byte("Привет")
	d2 := []byte("Это")
	d3 := []byte("Pervij")
	d4 := []byte("Тест")

	segmentStart1 := uint32(0)
	segmentEnd1 := uint32(2)
	segmentLength1 := uint32(segmentEnd1 - segmentStart1)
	segmentStart2 := uint32(2)
	segmentEnd2 := uint32(6)
	segmentLength2 := uint32(segmentEnd2 - segmentStart2)
	segmentStart3 := uint32(6)
	segmentEnd3 := uint32(10)
	segmentLength3 := uint32(segmentEnd3 - segmentStart3)
	segmentStart4 := uint32(10)
	segmentEnd4 := uint32(13)
	segmentLength4 := uint32(segmentEnd4 - segmentStart4)

	segmentLength12 := segmentLength1 + segmentLength2
	segmentLength34 := segmentLength3 + segmentLength4

	rootSegmentLength := segmentLength12 + segmentLength34

	node1 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...))
	node2 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength2), d2...))
	node3 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength3), d3...))
	node4 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength4), d4...))

	node12 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1), node1...), append(merkleTree.UintToBytesArray(segmentLength2), node2...)...))
	node34 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength3), node3...), append(merkleTree.UintToBytesArray(segmentLength4), node4...)...))

	rootHash := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength12), node12...), append(merkleTree.UintToBytesArray(segmentLength34), node34...)...))

	segments := []merkleTree.InputSegment{
		{segmentStart1, segmentEnd1, d1},
		{segmentStart2, segmentEnd2, d2},
		{segmentStart3, segmentEnd3, d3},
		{segmentStart4, segmentEnd4, d4},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if bytes.Equal(rootHash, tree.RootNode.Segment.Hash) && tree.RootNode.Segment.SegmentLength == rootSegmentLength {
		return "test_4_leafs: true\n"
	} else {
		return "test_4_leafs: false\n"
	}
}

func test_5_leafs() string {
	d1 := []byte("Привет")
	d2 := []byte("Это")
	d3 := []byte("Второй")
	d4 := []byte("Тест")
	d5 := []byte("Let's")

	segmentStart1 := uint32(0)
	segmentEnd1 := uint32(2)
	segmentLength1 := uint32(segmentEnd1 - segmentStart1)
	segmentStart2 := uint32(2)
	segmentEnd2 := uint32(6)
	segmentLength2 := uint32(segmentEnd2 - segmentStart2)
	segmentStart3 := uint32(6)
	segmentEnd3 := uint32(10)
	segmentLength3 := uint32(segmentEnd3 - segmentStart3)
	segmentStart4 := uint32(10)
	segmentEnd4 := uint32(13)
	segmentLength4 := uint32(segmentEnd4 - segmentStart4)
	segmentStart5 := uint32(13)
	segmentEnd5 := uint32(16)
	segmentLength5 := uint32(segmentEnd5 - segmentStart5)

	segmentLength12 := segmentLength1 + segmentLength2
	segmentLength34 := segmentLength3 + segmentLength4
	segmentLength1234 := segmentLength12 + segmentLength34

	rootSegmentLength := segmentLength1234 + segmentLength5

	node1 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...))
	node2 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength2), d2...))
	node3 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength3), d3...))
	node4 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength4), d4...))
	node5 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength5), d5...))

	node12 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1), node1...), append(merkleTree.UintToBytesArray(segmentLength2), node2...)...))
	node34 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength3), node3...), append(merkleTree.UintToBytesArray(segmentLength4), node4...)...))

	node1234 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength12), node12...), append(merkleTree.UintToBytesArray(segmentLength34), node34...)...))

	rootHash := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1234), node1234...), append(merkleTree.UintToBytesArray(segmentLength5), node5...)...))

	segments := []merkleTree.InputSegment{
		{segmentStart1, segmentEnd1, d1},
		{segmentStart2, segmentEnd2, d2},
		{segmentStart3, segmentEnd3, d3},
		{segmentStart4, segmentEnd4, d4},
		{segmentStart5, segmentEnd5, d5},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if bytes.Equal(rootHash, tree.RootNode.Segment.Hash) && tree.RootNode.Segment.SegmentLength == rootSegmentLength {
		return "test_5_leafs: true\n"
	} else {
		return "test_5_leafs: false\n"
	}
}

func test_6_leafs() string {
	d1 := []byte("Привет")
	d2 := []byte("Это")
	d3 := []byte("Тест")
	d4 := []byte("С")
	d5 := []byte("6")
	d6 := []byte("Листками")

	segmentStart1 := uint32(0)
	segmentEnd1 := uint32(2)
	segmentLength1 := uint32(segmentEnd1 - segmentStart1)
	segmentStart2 := uint32(2)
	segmentEnd2 := uint32(6)
	segmentLength2 := uint32(segmentEnd2 - segmentStart2)
	segmentStart3 := uint32(6)
	segmentEnd3 := uint32(10)
	segmentLength3 := uint32(segmentEnd3 - segmentStart3)
	segmentStart4 := uint32(10)
	segmentEnd4 := uint32(13)
	segmentLength4 := uint32(segmentEnd4 - segmentStart4)
	segmentStart5 := uint32(13)
	segmentEnd5 := uint32(14)
	segmentLength5 := uint32(segmentEnd5 - segmentStart5)
	segmentStart6 := uint32(14)
	segmentEnd6 := uint32(24)
	segmentLength6 := uint32(segmentEnd6 - segmentStart6)

	segmentLength12 := segmentLength1 + segmentLength2
	segmentLength34 := segmentLength3 + segmentLength4
	segmentLength56 := segmentLength5 + segmentLength6

	segmentLength1234 := segmentLength12 + segmentLength34

	rootSegmentLength := segmentLength1234 + segmentLength56

	node1 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...))
	node2 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength2), d2...))
	node3 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength3), d3...))
	node4 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength4), d4...))
	node5 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength5), d5...))
	node6 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength6), d6...))

	node12 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1), node1...), append(merkleTree.UintToBytesArray(segmentLength2), node2...)...))
	node34 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength3), node3...), append(merkleTree.UintToBytesArray(segmentLength4), node4...)...))
	node56 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength5), node5...), append(merkleTree.UintToBytesArray(segmentLength6), node6...)...))

	node1234 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength12), node12...), append(merkleTree.UintToBytesArray(segmentLength34), node34...)...))

	rootHash := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1234), node1234...), append(merkleTree.UintToBytesArray(segmentLength56), node56...)...))

	segments := []merkleTree.InputSegment{
		{segmentStart1, segmentEnd1, d1},
		{segmentStart2, segmentEnd2, d2},
		{segmentStart3, segmentEnd3, d3},
		{segmentStart4, segmentEnd4, d4},
		{segmentStart5, segmentEnd5, d5},
		{segmentStart6, segmentEnd6, d6},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if bytes.Equal(rootHash, tree.RootNode.Segment.Hash) && tree.RootNode.Segment.SegmentLength == rootSegmentLength {
		return "test_6_leafs: true\n"
	} else {
		return "test_6_leafs: false\n"
	}
}

func test_7_leafs() string {
	d1 := []byte("Привет")
	d2 := []byte("Это")
	d3 := []byte("Тест")
	d4 := []byte("С")
	d5 := []byte("7")
	d6 := []byte("Листками")
	d7 := []byte("...")

	segmentStart1 := uint32(0)
	segmentEnd1 := uint32(2)
	segmentLength1 := uint32(segmentEnd1 - segmentStart1)
	segmentStart2 := uint32(2)
	segmentEnd2 := uint32(7)
	segmentLength2 := uint32(segmentEnd2 - segmentStart2)
	segmentStart3 := uint32(7)
	segmentEnd3 := uint32(10)
	segmentLength3 := uint32(segmentEnd3 - segmentStart3)
	segmentStart4 := uint32(10)
	segmentEnd4 := uint32(13)
	segmentLength4 := uint32(segmentEnd4 - segmentStart4)
	segmentStart5 := uint32(13)
	segmentEnd5 := uint32(14)
	segmentLength5 := uint32(segmentEnd5 - segmentStart5)
	segmentStart6 := uint32(14)
	segmentEnd6 := uint32(20)
	segmentLength6 := uint32(segmentEnd6 - segmentStart6)
	segmentStart7 := uint32(20)
	segmentEnd7 := uint32(24)
	segmentLength7 := uint32(segmentEnd7 - segmentStart7)

	segmentLength12 := segmentLength1 + segmentLength2
	segmentLength34 := segmentLength3 + segmentLength4
	segmentLength56 := segmentLength5 + segmentLength6

	segmentLength1234 := segmentLength12 + segmentLength34
	segmentLength567 := segmentLength56 + segmentLength7

	rootSegmentLength := segmentLength1234 + segmentLength567

	node1 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...))
	node2 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength2), d2...))
	node3 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength3), d3...))
	node4 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength4), d4...))
	node5 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength5), d5...))
	node6 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength6), d6...))
	node7 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength7), d7...))

	node12 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1), node1...), append(merkleTree.UintToBytesArray(segmentLength2), node2...)...))
	node34 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength3), node3...), append(merkleTree.UintToBytesArray(segmentLength4), node4...)...))
	node56 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength5), node5...), append(merkleTree.UintToBytesArray(segmentLength6), node6...)...))

	node1234 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength12), node12...), append(merkleTree.UintToBytesArray(segmentLength34), node34...)...))
	node567 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength56), node56...), append(merkleTree.UintToBytesArray(segmentLength7), node7...)...))

	rootHash := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1234), node1234...), append(merkleTree.UintToBytesArray(segmentLength567), node567...)...))

	segments := []merkleTree.InputSegment{
		{segmentStart1, segmentEnd1, d1},
		{segmentStart2, segmentEnd2, d2},
		{segmentStart3, segmentEnd3, d3},
		{segmentStart4, segmentEnd4, d4},
		{segmentStart5, segmentEnd5, d5},
		{segmentStart6, segmentEnd6, d6},
		{segmentStart7, segmentEnd7, d7},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)
	if bytes.Equal(rootHash, tree.RootNode.Segment.Hash) && tree.RootNode.Segment.SegmentLength == rootSegmentLength {
		return "test_7_leafs: true\n"
	} else {
		return "test_7_leafs: false\n"
	}
}

func test_15_leafs() string {
	d1 := []byte("Привет")
	d2 := []byte("Это")
	d3 := []byte("Тест")
	d4 := []byte("С")
	d5 := []byte("15")
	d6 := []byte("Листками")
	d7 := []byte("...")
	d8 := []byte("This")
	d9 := []byte("Is")
	d10 := []byte("Test")
	d11 := []byte("for")
	d12 := []byte("15")
	d13 := []byte("Leafs")
	d14 := []byte("...")
	d15 := []byte("...!")

	segmentStart1 := uint32(0)
	segmentEnd1 := uint32(2)
	segmentLength1 := uint32(segmentEnd1 - segmentStart1)
	segmentStart2 := uint32(2)
	segmentEnd2 := uint32(6)
	segmentLength2 := uint32(segmentEnd2 - segmentStart2)
	segmentStart3 := uint32(6)
	segmentEnd3 := uint32(11)
	segmentLength3 := uint32(segmentEnd3 - segmentStart3)
	segmentStart4 := uint32(11)
	segmentEnd4 := uint32(13)
	segmentLength4 := uint32(segmentEnd4 - segmentStart4)
	segmentStart5 := uint32(13)
	segmentEnd5 := uint32(15)
	segmentLength5 := uint32(segmentEnd5 - segmentStart5)
	segmentStart6 := uint32(15)
	segmentEnd6 := uint32(20)
	segmentLength6 := uint32(segmentEnd6 - segmentStart6)
	segmentStart7 := uint32(20)
	segmentEnd7 := uint32(25)
	segmentLength7 := uint32(segmentEnd7 - segmentStart7)
	segmentStart8 := uint32(25)
	segmentEnd8 := uint32(27)
	segmentLength8 := uint32(segmentEnd8 - segmentStart8)
	segmentStart9 := uint32(27)
	segmentEnd9 := uint32(29)
	segmentLength9 := uint32(segmentEnd9 - segmentStart9)
	segmentStart10 := uint32(29)
	segmentEnd10 := uint32(31)
	segmentLength10 := uint32(segmentEnd10 - segmentStart10)
	segmentStart11 := uint32(31)
	segmentEnd11 := uint32(34)
	segmentLength11 := uint32(segmentEnd11 - segmentStart11)
	segmentStart12 := uint32(34)
	segmentEnd12 := uint32(40)
	segmentLength12 := uint32(segmentEnd12 - segmentStart12)
	segmentStart13 := uint32(40)
	segmentEnd13 := uint32(42)
	segmentLength13 := uint32(segmentEnd13 - segmentStart13)
	segmentStart14 := uint32(42)
	segmentEnd14 := uint32(45)
	segmentLength14 := uint32(segmentEnd14 - segmentStart14)
	segmentStart15 := uint32(45)
	segmentEnd15 := uint32(100)
	segmentLength15 := uint32(segmentEnd15 - segmentStart15)

	segmentLength1_2 := segmentLength1 + segmentLength2
	segmentLength34 := segmentLength3 + segmentLength4
	segmentLength56 := segmentLength5 + segmentLength6
	segmentLength78 := segmentLength7 + segmentLength8
	segmentLength9_10 := segmentLength9 + segmentLength10
	segmentLength11_12 := segmentLength11 + segmentLength12
	segmentLength13_14 := segmentLength13 + segmentLength14

	segmentLength1234 := segmentLength1_2 + segmentLength34
	segmentLength5678 := segmentLength56 + segmentLength78
	segmentLength9_10_11_12 := segmentLength9_10 + segmentLength11_12
	segmentLength13_14_15 := segmentLength13_14 + segmentLength15

	segmentLength12345678 := segmentLength1234 + segmentLength5678
	segmentLength9_15 := segmentLength9_10_11_12 + segmentLength13_14_15

	rootSegmentLength := segmentLength12345678 + segmentLength9_15

	node1 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...))
	node2 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength2), d2...))
	node3 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength3), d3...))
	node4 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength4), d4...))
	node5 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength5), d5...))
	node6 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength6), d6...))
	node7 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength7), d7...))
	node8 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength8), d8...))
	node9 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength9), d9...))
	node10 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength10), d10...))
	node11 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength11), d11...))
	node12 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength12), d12...))
	node13 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength13), d13...))
	node14 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength14), d14...))
	node15 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength15), d15...))

	node1_2 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1), node1...), append(merkleTree.UintToBytesArray(segmentLength2), node2...)...))
	node34 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength3), node3...), append(merkleTree.UintToBytesArray(segmentLength4), node4...)...))
	node56 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength5), node5...), append(merkleTree.UintToBytesArray(segmentLength6), node6...)...))
	node78 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength7), node7...), append(merkleTree.UintToBytesArray(segmentLength8), node8...)...))
	node9_10 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength9), node9...), append(merkleTree.UintToBytesArray(segmentLength10), node10...)...))
	node11_12 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength11), node11...), append(merkleTree.UintToBytesArray(segmentLength12), node12...)...))
	node13_14 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength13), node13...), append(merkleTree.UintToBytesArray(segmentLength14), node14...)...))

	node1234 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1_2), node1_2...), append(merkleTree.UintToBytesArray(segmentLength34), node34...)...))
	node5678 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength56), node56...), append(merkleTree.UintToBytesArray(segmentLength78), node78...)...))
	node9_10_11_12 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength9_10), node9_10...), append(merkleTree.UintToBytesArray(segmentLength11_12), node11_12...)...))
	node13_14_15 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength13_14), node13_14...), append(merkleTree.UintToBytesArray(segmentLength15), node15...)...))

	node12345678 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1234), node1234...), append(merkleTree.UintToBytesArray(segmentLength5678), node5678...)...))

	node9_15 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength9_10_11_12), node9_10_11_12...), append(merkleTree.UintToBytesArray(segmentLength13_14_15), node13_14_15...)...))

	rootHash := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength12345678), node12345678...), append(merkleTree.UintToBytesArray(segmentLength9_15), node9_15...)...))

	segments := []merkleTree.InputSegment{
		{segmentStart1, segmentEnd1, d1},
		{segmentStart2, segmentEnd2, d2},
		{segmentStart3, segmentEnd3, d3},
		{segmentStart4, segmentEnd4, d4},
		{segmentStart5, segmentEnd5, d5},
		{segmentStart6, segmentEnd6, d6},
		{segmentStart7, segmentEnd7, d7},
		{segmentStart8, segmentEnd8, d8},
		{segmentStart9, segmentEnd9, d9},
		{segmentStart10, segmentEnd10, d10},
		{segmentStart11, segmentEnd11, d11},
		{segmentStart12, segmentEnd12, d12},
		{segmentStart13, segmentEnd13, d13},
		{segmentStart14, segmentEnd14, d14},
		{segmentStart15, segmentEnd15, d15},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if bytes.Equal(rootHash, tree.RootNode.Segment.Hash) && tree.RootNode.Segment.SegmentLength == rootSegmentLength {
		return "test_15_leafs: true\n"
	} else {
		return "test_15_leafs: false\n"
	}
}

func TestSortSegments() string {

	var testArr []merkleTree.InputSegment

	one := merkleTree.InputSegment{Start: 3, End: 4, Data: []byte("1")}
	two := merkleTree.InputSegment{Start: 7, End: 9, Data: []byte("2")}
	three := merkleTree.InputSegment{Start: 9, End: 10, Data: []byte("3")}
	four := merkleTree.InputSegment{Start: 13, End: 14, Data: []byte("4")}
	five := merkleTree.InputSegment{Start: 29, End: 50, Data: []byte("5")}
	six := merkleTree.InputSegment{Start: 90, End: 92, Data: []byte("6")}
	seven := merkleTree.InputSegment{Start: 92, End: 93, Data: []byte("7")}
	eigth := merkleTree.InputSegment{Start: 12224, End: 28911, Data: []byte("8")}

	testArr = append(testArr, one)
	testArr = append(testArr, two)
	testArr = append(testArr, three)
	testArr = append(testArr, four)
	testArr = append(testArr, five)
	testArr = append(testArr, six)
	testArr = append(testArr, seven)
	testArr = append(testArr, eigth)

	shouldSorted := []merkleTree.InputSegment{
		merkleTree.InputSegment{Start: 0, End: 3, Data: []byte{}},
		one,
		merkleTree.InputSegment{Start: 4, End: 7, Data: []byte{}},
		two,
		three,
		merkleTree.InputSegment{Start: 10, End: 13, Data: []byte{}},
		four,
		merkleTree.InputSegment{Start: 14, End: 29, Data: []byte{}},
		five,
		merkleTree.InputSegment{Start: 50, End: 90, Data: []byte{}},
		six,
		seven,
		merkleTree.InputSegment{Start: 93, End: 12224, Data: []byte{}},
		eigth,
		merkleTree.InputSegment{Start: 28911, End: 16777215, Data: []byte{}},
	}

	sorted := merkleTree.PrepareSegments(testArr)

	if reflect.DeepEqual(sorted, shouldSorted) {
		return "TestSortSegments: true\n"
	} else {
		return "TestSortSegments: false\n"
	}
}

func test_tree_with_sort() string {
	d1 := []byte("Привет")
	d2 := []byte("Это")
	d3 := []byte("Sort")
	d4 := []byte("Тест")
	d5 := []byte("Let's go")
	zeroByteArray := []byte{}
	rootSegmentLength := uint32(16777216 - 1)

	segmentStart1 := uint32(123)
	segmentEnd1 := uint32(234)
	segmentLength1 := uint32(segmentEnd1 - segmentStart1)
	segmentStart2 := uint32(244)
	segmentEnd2 := uint32(444)
	segmentLength2 := uint32(segmentEnd2 - segmentStart2)
	segmentStart3 := uint32(444)
	segmentEnd3 := uint32(555)
	segmentLength3 := uint32(segmentEnd3 - segmentStart3)
	segmentStart4 := uint32(1023)
	segmentEnd4 := uint32(1111)
	segmentLength4 := uint32(segmentEnd4 - segmentStart4)
	segmentStart5 := uint32(1200)
	segmentEnd5 := uint32(3000)
	segmentLength5 := uint32(segmentEnd5 - segmentStart5)

	segmentStartZero0 := uint32(0)
	segmentEndZero0 := segmentStart1
	segmentLengthZero0 := uint32(segmentEndZero0 - segmentStartZero0)

	segmentStartZero1 := segmentEnd1
	segmentEndZero1 := segmentStart2
	segmentLengthZero1 := uint32(segmentEndZero1 - segmentStartZero1)

	segmentStartZero3 := segmentEnd3
	segmentEndZero3 := segmentStart4
	segmentLengthZero3 := uint32(segmentEndZero3 - segmentStartZero3)

	segmentStartZero4 := segmentEnd4
	segmentEndZero4 := segmentStart5
	segmentLengthZero4 := uint32(segmentEndZero4 - segmentStartZero4)

	segmentStartZero5 := segmentEnd5
	segmentEndZero5 := rootSegmentLength
	segmentLengthZero5 := uint32(segmentEndZero5 - segmentStartZero5)

	segmentLength12 := segmentLengthZero0 + segmentLength1
	segmentLength34 := segmentLengthZero1 + segmentLength2
	segmentLength56 := segmentLength3 + segmentLengthZero3
	segmentLength78 := segmentLength4 + segmentLengthZero4
	segmentLength910 := segmentLength5 + segmentLengthZero5

	segmentLength1234 := segmentLength12 + segmentLength34
	segmentLength5678 := segmentLength56 + segmentLength78

	segmentLength12345678 := segmentLength1234 + segmentLength5678

	node1 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLengthZero0), zeroByteArray...))
	node2 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...))
	node3 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLengthZero1), zeroByteArray...))
	node4 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength2), d2...))
	node5 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength3), d3...))
	node6 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLengthZero3), zeroByteArray...))
	node7 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength4), d4...))
	node8 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLengthZero4), zeroByteArray...))
	node9 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength5), d5...))
	node10 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLengthZero5), zeroByteArray...))

	node12 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLengthZero0), node1...), append(merkleTree.UintToBytesArray(segmentLength1), node2...)...))
	node34 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLengthZero1), node3...), append(merkleTree.UintToBytesArray(segmentLength2), node4...)...))
	node56 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength3), node5...), append(merkleTree.UintToBytesArray(segmentLengthZero3), node6...)...))
	node78 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength4), node7...), append(merkleTree.UintToBytesArray(segmentLengthZero4), node8...)...))
	node910 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength5), node9...), append(merkleTree.UintToBytesArray(segmentLengthZero5), node10...)...))

	node1234 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength12), node12...), append(merkleTree.UintToBytesArray(segmentLength34), node34...)...))
	node5678 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength56), node56...), append(merkleTree.UintToBytesArray(segmentLength78), node78...)...))

	node12345678 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1234), node1234...), append(merkleTree.UintToBytesArray(segmentLength5678), node5678...)...))

	rootHash := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength12345678), node12345678...), append(merkleTree.UintToBytesArray(segmentLength910), node910...)...))

	segments := merkleTree.PrepareSegments([]merkleTree.InputSegment{
		{segmentStart1, segmentEnd1, d1},
		{segmentStart2, segmentEnd2, d2},
		{segmentStart3, segmentEnd3, d3},
		{segmentStart4, segmentEnd4, d4},
		{segmentStart5, segmentEnd5, d5},
	})

	shouldSegments := []merkleTree.InputSegment{
		{0, segmentStart1, zeroByteArray},
		{segmentStart1, segmentEnd1, d1},
		{segmentEnd1, segmentStart2, zeroByteArray},
		{segmentStart2, segmentEnd2, d2},
		{segmentStart3, segmentEnd3, d3},
		{segmentEnd3, segmentStart4, zeroByteArray},
		{segmentStart4, segmentEnd4, d4},
		{segmentEnd4, segmentStart5, zeroByteArray},
		{segmentStart5, segmentEnd5, d5},
		{segmentEnd5, rootSegmentLength, zeroByteArray},
	}

	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if bytes.Equal(rootHash, tree.RootNode.Segment.Hash) && tree.RootNode.Segment.SegmentLength == rootSegmentLength && reflect.DeepEqual(segments, shouldSegments) {
		return "test_tree_with_sort: true\n"
	} else {
		return "test_tree_with_sort: false\n"
	}
}

//func test_get_proof() string {
//	d1 := []byte("Привет")
//	d2 := []byte("Это")
//	d3 := []byte("Тест")
//	d4 := []byte("По")
//	d5 := []byte("Получению")
//	d6 := []byte("Доказательства")
//
//	segmentLength1 := uint32(1)
//	segmentLength2 := uint32(1)
//	segmentLength3 := uint32(1)
//	segmentLength4 := uint32(1)
//	segmentLength5 := uint32(1)
//	segmentLength6 := uint32(1)
//
//	tree := merkleTree.NewMerkleTree(
//		[]merkleTree.Segment{
//			merkleTree.Segment{segmentLength1, d1},
//			merkleTree.Segment{segmentLength2, d2},
//			merkleTree.Segment{segmentLength3, d3},
//			merkleTree.Segment{segmentLength4, d4},
//			merkleTree.Segment{segmentLength5, d5},
//			merkleTree.Segment{segmentLength6, d6},
//		},
//		crypto.Keccak256,
//		)
//
//	trueProof := merkleTree.MerkleProof{
//		tree.Levels,
//		tree.RootNode.Segment.Data,
//		tree.RootNode.Segment.SegmentLength,
//		merkleTree.Segment{segmentLength4, d4},
//	}
//
//	proof, _ := tree.GetProof(merkleTree.Segment{segmentLength4, d4})
//
//	if reflect.DeepEqual(&trueProof, proof) {
//		return "test_get_proof: true\n"
//	} else {
//		return "test_get_proof: false\n"
//	}
//}
//
//func test_verify_proof() string {
//	d1 := []byte("Привет")
//	d2 := []byte("Это")
//	d3 := []byte("Тест")
//	d4 := []byte("По")
//	d5 := []byte("Получению")
//	d6 := append([]byte("Доказательства"), []byte("Пруфа")...)
//
//	segmentLength1 := uint32(1)
//	segmentLength2 := uint32(1)
//	segmentLength3 := uint32(1)
//	segmentLength4 := uint32(1)
//	segmentLength5 := uint32(1)
//	segmentLength6 := uint32(2)
//
//	tree := merkleTree.NewMerkleTree(
//		[]merkleTree.Segment{
//			merkleTree.Segment{segmentLength1, d1},
//			merkleTree.Segment{segmentLength2, d2},
//			merkleTree.Segment{segmentLength3, d3},
//			merkleTree.Segment{segmentLength4, d4},
//			merkleTree.Segment{segmentLength5, d5},
//			merkleTree.Segment{segmentLength6, d6},
//		},
//		crypto.Keccak256,
//	)
//
//	proof, _ := tree.GetProof(merkleTree.Segment{segmentLength4, d4})
//
//	verifyProof := merkleTree.Verify(proof, tree.RootNode.Segment.Data)
//
//	if verifyProof == true {
//		return "test_verify_proof: true\n"
//	} else {
//		return "test_verify_proof: false\n"
//	}
//}
