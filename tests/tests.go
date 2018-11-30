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
}

func test_leaf_to_node1() string {
	d1 := []byte("Лист")

	dataHash := crypto.Keccak256(d1)

	segment := merkleTree.Segment{1, d1}
	hash := merkleTree.LeafToNode(segment, crypto.Keccak256)
	if (bytes.Equal(dataHash, hash.Segment.Data) && hash.Segment.SegmentLength == 1) {
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

	nodes := [][]byte{crypto.Keccak256(d1), crypto.Keccak256(d2), crypto.Keccak256(d3), crypto.Keccak256(d4), crypto.Keccak256(d5), crypto.Keccak256(d6)}

	segments := []merkleTree.Segment{
		{1, d1},
		{1, d2},
		{1, d3},
		{1, d4},
		{1, d5},
	}

	if (len(segments)%2 != 0) {
		segments = append(segments, merkleTree.Segment{0, []byte{}})
	}

	var check = true
	for i, s := range segments {
		hash := merkleTree.LeafToNode(s, crypto.Keccak256)
		if (!bytes.Equal(hash.Segment.Data, nodes[i])) {
			check = false
			break
		}
	}


	if (check == true) {
		return "test_leaf_to_node2: true\n"
	} else {
		return "test_leaf_to_node2: false\n"
	}
}

func test_New_Merkle_Node_with_left_and_right_nodes() string {
	d1 := []byte("Лист1")
	d2 := []byte("Лист2")

	node1 := crypto.Keccak256(d1)
	node2 := crypto.Keccak256(d2)
	node12 := crypto.Keccak256(append(node1, node2...))

	leftNode := merkleTree.MerkleNode{nil, nil, merkleTree.Segment{1, node1}}
	rightNode := merkleTree.MerkleNode{nil, nil, merkleTree.Segment{1, node2}}

	newNode := merkleTree.MerkleNode{&leftNode, &rightNode, merkleTree.Segment{2, node12}}

	node := merkleTree.NewMerkleNode(&leftNode, &rightNode, crypto.Keccak256)
	if (reflect.DeepEqual(node, &newNode)) {
		return "test_New_Merkle_Node_with_left_and_right_nodes: true\n"
	} else {
		return "test_New_Merkle_Node_with_left_and_right_nodes: false\n"
	}
}

func test_New_Merkle_Node_with_nil_right_node() string {
	somed1 := []byte("Некоторые данные1")
	somed2 := []byte("Некоторые данные2")

	somenode1 := crypto.Keccak256(somed1)
	somenode2 := crypto.Keccak256(somed2)

	somenode12 := crypto.Keccak256(append(somenode1, somenode2...))
	somenode12zero := crypto.Keccak256(append(somenode12, zeroHash...))

	leftNode0 := merkleTree.MerkleNode{nil, nil, merkleTree.Segment{1, somenode1}}
	rightNode0 := merkleTree.MerkleNode{nil, nil, merkleTree.Segment{1, somenode2}}
	leftNode := merkleTree.MerkleNode{&leftNode0, &rightNode0, merkleTree.Segment{2, somenode12}}

	newNode := merkleTree.MerkleNode{&leftNode, nil, merkleTree.Segment{2, somenode12zero}}

	node := merkleTree.NewMerkleNode(&leftNode, nil, crypto.Keccak256)
	if (reflect.DeepEqual(node, &newNode)) {
		return "test_New_Merkle_Node_with_nil_right_node: true\n"
	} else {
		return "test_New_Merkle_Node_with_nil_right_node: false\n"
	}
}

func test_3_leafs() string {
	d1 := []byte("Есть")
	d2 := []byte("Три")
	d3 := []byte("Листа")
	d4 := []byte{}

	hash_d1 := crypto.Keccak256(d1)
	hash_d2 := crypto.Keccak256(d2)
	hash_d3 := crypto.Keccak256(d3)
	hash_d4 := crypto.Keccak256(d4)

	hash_d1_d2 := crypto.Keccak256(append(hash_d1, hash_d2...))
	hash_d3_d4 := crypto.Keccak256(append(hash_d3, hash_d4...))

	rootHash := crypto.Keccak256(append(hash_d1_d2, hash_d3_d4...))


	segments := []merkleTree.Segment{
		{1, d1},
		{1, d2},
		{1, d3},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if (bytes.Equal(rootHash, tree.RootNode.Segment.Data) && tree.RootNode.Segment.SegmentLength == 3) {
		return "test_3_leafs: true\n"
	} else {
		return "test_3_leafs: false\n"
	}
}

func test_4_leafs() string {
	d1 := []byte("Привет")
	d2 := []byte("Это")
	d3 := []byte("Первый")
	d4 := []byte("Тест")

	node1 := crypto.Keccak256(d1)
	node2 := crypto.Keccak256(d2)
	node3 := crypto.Keccak256(d3)
	node4 := crypto.Keccak256(d4)

	node12 := crypto.Keccak256(append(node1, node2...))
	node34 := crypto.Keccak256(append(node3, node4...))

	rootHash := crypto.Keccak256(append(node12, node34...))


	segments := []merkleTree.Segment{
		{1, d1},
		{1, d2},
		{1, d3},
		{1, d4},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if (bytes.Equal(rootHash, tree.RootNode.Segment.Data) && tree.RootNode.Segment.SegmentLength == 4) {
		return "test_4_leafs: true\n"
	} else {
		return "test_4_leafs: false\n"
	}
}

func test_5_leafs() string {
	d1 := []byte("Есть")
	d2 := []byte("Пять")
	d3 := []byte("Входных")
	d4 := []byte("Данных")
	d5 := []byte(".")
	d6 := []byte{}

	node1 := crypto.Keccak256(d1)
	node2 := crypto.Keccak256(d2)
	node3 := crypto.Keccak256(d3)
	node4 := crypto.Keccak256(d4)
	node5 := crypto.Keccak256(d5)
	node6 := crypto.Keccak256(d6)

	node12 := crypto.Keccak256(append(node1, node2...))
	node34 := crypto.Keccak256(append(node3, node4...))
	node56 := crypto.Keccak256(append(node5, node6...))

	node1234 := crypto.Keccak256(append(node12, node34...))
	node56zero := crypto.Keccak256(append(node56, zeroHash...))

	rootHash := crypto.Keccak256(append(node1234, node56zero...))


	segments := []merkleTree.Segment{
		{1, d1},
		{1, d2},
		{1, d3},
		{1, d4},
		{1, d5},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if (bytes.Equal(rootHash, tree.RootNode.Segment.Data) && tree.RootNode.Segment.SegmentLength == 5) {
		return "test_5_leafs: true\n"
	} else {
		return "test_5_leafs: false\n"
	}
}

func test_6_leafs() string {
	d1 := []byte("Есть")
	d2 := []byte("Шесть")
	d3 := []byte("Входных")
	d4 := []byte("Данных")
	d5 := []byte("Вот")
	d6 := []byte("....")

	node1 := crypto.Keccak256(d1)
	node2 := crypto.Keccak256(d2)
	node3 := crypto.Keccak256(d3)
	node4 := crypto.Keccak256(d4)
	node5 := crypto.Keccak256(d5)
	node6 := crypto.Keccak256(d6)

	node12 := crypto.Keccak256(append(node1, node2...))
	node34 := crypto.Keccak256(append(node3, node4...))
	node56 := crypto.Keccak256(append(node5, node6...))

	node1234 := crypto.Keccak256(append(node12, node34...))
	node56zero := crypto.Keccak256(append(node56, zeroHash...))

	rootHash := crypto.Keccak256(append(node1234, node56zero...))


	segments := []merkleTree.Segment{
		{5, d1},
		{1, d2},
		{2, d3},
		{1, d4},
		{1, d5},
		{4, d6},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if (bytes.Equal(rootHash, tree.RootNode.Segment.Data) && tree.RootNode.Segment.SegmentLength == 14) {
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

	node1 := crypto.Keccak256(d1)
	node2 := crypto.Keccak256(d2)
	node3 := crypto.Keccak256(d3)
	node4 := crypto.Keccak256(d4)
	node5 := crypto.Keccak256(d5)
	node6 := crypto.Keccak256(d6)
	node7 := crypto.Keccak256(d7)
	node8 := crypto.Keccak256(d8)

	node12 := crypto.Keccak256(append(node1, node2...))
	node34 := crypto.Keccak256(append(node3, node4...))
	node56 := crypto.Keccak256(append(node5, node6...))
	node78 := crypto.Keccak256(append(node7, node8...))

	node1234 := crypto.Keccak256(append(node12, node34...))
	node5678 := crypto.Keccak256(append(node56, node78...))

	rootHash := crypto.Keccak256(append(node1234, node5678...))


	segments := []merkleTree.Segment{
		{1, d1},
		{1, d2},
		{1, d3},
		{1, d4},
		{1, d5},
		{1, d6},
		{1, d7},
		{1, d8},
	}
	tree := merkleTree.NewMerkleTree(segments, crypto.Keccak256)

	if (bytes.Equal(rootHash, tree.RootNode.Segment.Data)) {
		return "test_8_leafs: true\n"
	} else {
		return "test_8_leafs: false\n"
	}
}

