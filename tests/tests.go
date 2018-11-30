package main

import (
	"../merkleTree"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"reflect"
)

var zeroHash = crypto.Keccak256([]byte{})

func main()  {
	fmt.Print(test_leaf_to_node1())
	fmt.Print(test_leaf_to_node2())
	fmt.Print(test_New_Merkle_Node_with_left_and_right_nodes())
	fmt.Print(test_New_Merkle_Node_with_nil_right_node())
	fmt.Print(test_3_leafs())
	fmt.Print(test_4_leafs())
	fmt.Print(test_5_leafs())
	fmt.Print(test_6_leafs())
	fmt.Print(test_8_leafs())
	fmt.Print(test_get_proof())
}

func test_leaf_to_node1() string {
	d1 := []byte("Лист")
	segmentLength1 := uint32(1)

	dataHash := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...))

	segment := merkleTree.Segment{segmentLength1, d1}
	hash := merkleTree.LeafToNode(segment, crypto.Keccak256)
	if bytes.Equal(dataHash, hash.Segment.Data) && hash.Segment.SegmentLength == 1 {
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

	segmentLength1 := uint32(1)
	segmentLength2 := uint32(1)
	segmentLength3 := uint32(1)
	segmentLength4 := uint32(1)
	segmentLength5 := uint32(1)
	segmentLength6 := uint32(0)

	nodes := [][]byte{
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...)),
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength2), d2...)),
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength3), d3...)),
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength4), d4...)),
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength5), d5...)),
		crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength6), d6...)),
	}

	segments := []merkleTree.Segment{
		{1, d1},
		{1, d2},
		{1, d3},
		{1, d4},
		{1, d5},
	}

	if len(segments)%2 != 0 {
		segments = append(segments, merkleTree.Segment{0, []byte{}})
	}

	var check = true
	for i, s := range segments {
		hash := merkleTree.LeafToNode(s, crypto.Keccak256)
		if !bytes.Equal(hash.Segment.Data, nodes[i]) {
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

	leftNode := merkleTree.MerkleNode{nil, nil, merkleTree.Segment{someSegmentLenght1, node1}}
	rightNode := merkleTree.MerkleNode{nil, nil, merkleTree.Segment{someSegmentLenght2, node2}}

	newNode := merkleTree.MerkleNode{&leftNode, &rightNode, merkleTree.Segment{2, node12}}

	node := merkleTree.NewMerkleNode(&leftNode, &rightNode, crypto.Keccak256)
	if reflect.DeepEqual(node, &newNode) {
		return "test_New_Merkle_Node_with_left_and_right_nodes: true\n"
	} else {
		return "test_New_Merkle_Node_with_left_and_right_nodes: false\n"
	}
}

func test_New_Merkle_Node_with_nil_right_node() string {
	somed1 := []byte("Некоторые данные1")
	somed2 := []byte("Некоторые данные2")
	someSegmentLenght1 := uint32(1)
	someSegmentLenght2 := uint32(1)
	someSegmentLenght12 := uint32(someSegmentLenght1 + someSegmentLenght2)

	somenode1 := crypto.Keccak256(append(merkleTree.UintToBytesArray(someSegmentLenght1), somed1...))
	somenode2 := crypto.Keccak256(append(merkleTree.UintToBytesArray(someSegmentLenght2), somed2...))

	somenode12 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(someSegmentLenght1), somenode1...), append(merkleTree.UintToBytesArray(someSegmentLenght2), somenode2...)...))
	somenode12Zero := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(someSegmentLenght12), somenode12...), append(merkleTree.UintToBytesArray(0), zeroHash...)...))

	leftNode0 := merkleTree.MerkleNode{nil, nil, merkleTree.Segment{1, somenode1}}
	rightNode0 := merkleTree.MerkleNode{nil, nil, merkleTree.Segment{1, somenode2}}
	leftNode := merkleTree.MerkleNode{&leftNode0, &rightNode0, merkleTree.Segment{2, somenode12}}

	newNode := merkleTree.MerkleNode{&leftNode, nil, merkleTree.Segment{2, somenode12Zero}}

	node := merkleTree.NewMerkleNode(&leftNode, nil, crypto.Keccak256)
	if reflect.DeepEqual(node, &newNode) {
		return "test_New_Merkle_Node_with_nil_right_node: true\n"
	} else {
		return "test_New_Merkle_Node_with_nil_right_node: false\n"
	}
}

func test_3_leafs() string {
	d1 := []byte("Есть")
	d2 := []byte("3")
	d3 := []byte("Листа")
	d4 := []byte{}

	segmentLength1 := uint32(1)
	segmentLength2 := uint32(1)
	segmentLength3 := uint32(1)
	segmentLength4 := uint32(0)

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

	segments := []merkleTree.Segment{
		{segmentLength1, d1},
		{segmentLength2, d2},
		{segmentLength3, d3},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if bytes.Equal(rootHash, tree.RootNode.Segment.Data) && tree.RootNode.Segment.SegmentLength == rootSegmentLength {
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

	segmentLength1 := uint32(1)
	segmentLength2 := uint32(1)
	segmentLength3 := uint32(1)
	segmentLength4 := uint32(1)

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

	segments := []merkleTree.Segment{
		{segmentLength1, d1},
		{segmentLength2, d2},
		{segmentLength3, d3},
		{segmentLength4, d4},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if bytes.Equal(rootHash, tree.RootNode.Segment.Data) && tree.RootNode.Segment.SegmentLength == rootSegmentLength {
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
	d6 := []byte{}

	segmentLength1 := uint32(1)
	segmentLength2 := uint32(1)
	segmentLength3 := uint32(1)
	segmentLength4 := uint32(1)
	segmentLength5 := uint32(1)
	segmentLength6 := uint32(0)

	segmentLength12 := segmentLength1 + segmentLength2
	segmentLength34 := segmentLength3 + segmentLength4
	segmentLength56 := segmentLength5 + segmentLength6

	segmentLength1234 := segmentLength12 + segmentLength34
	segmentLength56Zero := segmentLength56 + 0

	rootSegmentLength := segmentLength1234 + segmentLength56Zero

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
	node56Zero := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength56), node56...), append(merkleTree.UintToBytesArray(0), zeroHash...)...))

	rootHash := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1234), node1234...), append(merkleTree.UintToBytesArray(segmentLength56Zero), node56Zero...)...))


	segments := []merkleTree.Segment{
		{segmentLength1, d1},
		{segmentLength2, d2},
		{segmentLength3, d3},
		{segmentLength4, d4},
		{segmentLength5, d5},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if bytes.Equal(rootHash, tree.RootNode.Segment.Data) && tree.RootNode.Segment.SegmentLength == rootSegmentLength {
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

	segmentLength1 := uint32(1)
	segmentLength2 := uint32(1)
	segmentLength3 := uint32(1)
	segmentLength4 := uint32(1)
	segmentLength5 := uint32(1)
	segmentLength6 := uint32(1)

	segmentLength12 := segmentLength1 + segmentLength2
	segmentLength34 := segmentLength3 + segmentLength4
	segmentLength56 := segmentLength5 + segmentLength6
	segmentLengthZero := uint32(0)

	segmentLength1234 := segmentLength12 + segmentLength34
	segmentLength56zero := segmentLength56 + 0

	rootSegmentLength := segmentLength1234 + segmentLength56zero

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
	node56Zero := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength56), node56...), append(merkleTree.UintToBytesArray(segmentLengthZero), zeroHash...)...))

	rootHash := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1234), node1234...), append(merkleTree.UintToBytesArray(segmentLength56zero), node56Zero...)...))


	segments := []merkleTree.Segment{
		{segmentLength1, d1},
		{segmentLength2, d2},
		{segmentLength3, d3},
		{segmentLength4, d4},
		{segmentLength5, d5},
		{segmentLength6, d6},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if bytes.Equal(rootHash, tree.RootNode.Segment.Data) && tree.RootNode.Segment.SegmentLength == rootSegmentLength {
		return "test_6_leafs: true\n"
	} else {
		return "test_6_leafs: false\n"
	}
}

func test_8_leafs() string {
	d1 := []byte("Привет")
	d2 := []byte("Это")
	d3 := []byte("Второй")
	d4 := []byte("Тест")
	d5 := []byte("Let's")
	d6 := []byte("Check")
	d7 := []byte("It")
	d8 := []byte("!")

	segmentLength1 := uint32(1)
	segmentLength2 := uint32(1)
	segmentLength3 := uint32(1)
	segmentLength4 := uint32(1)
	segmentLength5 := uint32(1)
	segmentLength6 := uint32(1)
	segmentLength7 := uint32(1)
	segmentLength8 := uint32(1)

	segmentLength12 := segmentLength1 + segmentLength2
	segmentLength34 := segmentLength3 + segmentLength4
	segmentLength56 := segmentLength5 + segmentLength6
	segmentLength78 := segmentLength7 + segmentLength8

	segmentLength1234 := segmentLength12 + segmentLength34
	segmentLength5678 := segmentLength56 + segmentLength78

	rootSegmentLength := segmentLength1234 + segmentLength5678

	node1 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength1), d1...))
	node2 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength2), d2...))
	node3 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength3), d3...))
	node4 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength4), d4...))
	node5 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength5), d5...))
	node6 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength6), d6...))
	node7 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength7), d7...))
	node8 := crypto.Keccak256(append(merkleTree.UintToBytesArray(segmentLength8), d8...))

	node12 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1), node1...), append(merkleTree.UintToBytesArray(segmentLength2), node2...)...))
	node34 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength3), node3...), append(merkleTree.UintToBytesArray(segmentLength4), node4...)...))
	node56 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength5), node5...), append(merkleTree.UintToBytesArray(segmentLength6), node6...)...))
	node78 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength7), node7...), append(merkleTree.UintToBytesArray(segmentLength8), node8...)...))

	node1234 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength12), node12...), append(merkleTree.UintToBytesArray(segmentLength34), node34...)...))
	node5678 := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength56), node56...), append(merkleTree.UintToBytesArray(segmentLength78), node78...)...))

	rootHash := crypto.Keccak256(append(append(merkleTree.UintToBytesArray(segmentLength1234), node1234...), append(merkleTree.UintToBytesArray(segmentLength5678), node5678...)...))


	segments := []merkleTree.Segment{
		{segmentLength1, d1},
		{segmentLength2, d2},
		{segmentLength3, d3},
		{segmentLength4, d4},
		{segmentLength5, d5},
		{segmentLength6, d6},
		{segmentLength7, d7},
		{segmentLength8, d8},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if bytes.Equal(rootHash, tree.RootNode.Segment.Data) && tree.RootNode.Segment.SegmentLength == rootSegmentLength {
		return "test_8_leafs: true\n"
	} else {
		return "test_8_leafs: false\n"
	}
}

func test_get_proof() string {
	d1 := []byte("Привет")
	d2 := []byte("Это")
	d3 := []byte("Тест")
	d4 := []byte("По")
	d5 := []byte("Получению")
	d6 := []byte("Доказательства")

	segmentLength1 := uint32(1)
	segmentLength2 := uint32(1)
	segmentLength3 := uint32(1)
	segmentLength4 := uint32(1)
	segmentLength5 := uint32(1)
	segmentLength6 := uint32(1)

	tree := merkleTree.NewMerkleTree(
		[]merkleTree.Segment{
			merkleTree.Segment{segmentLength1, d1},
			merkleTree.Segment{segmentLength2, d2},
			merkleTree.Segment{segmentLength3, d3},
			merkleTree.Segment{segmentLength4, d4},
			merkleTree.Segment{segmentLength5, d5},
			merkleTree.Segment{segmentLength6, d6},
		},
		crypto.Keccak256,
		)

	trueProof := merkleTree.MerkleProof{
		tree.Levels,
		tree.RootNode.Segment.Data,
		merkleTree.Segment{segmentLength4, d4},
	}

	proof, _ := tree.GetProof(merkleTree.Segment{segmentLength4, d4})

	if reflect.DeepEqual(&trueProof, proof) {
		return "test_get_proof: true\n"
	} else {
		return "test_get_proof: false\n"
	}
}