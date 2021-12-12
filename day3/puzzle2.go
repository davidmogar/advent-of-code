package day3

import (
	"strconv"
)

type Tree struct {
	root *Node
}

type Node struct {
	leftBranch  *Node
	rightBranch *Node
	count       int
	value       int
}

func (node *Node) getBranch(lessSignificant bool) *Node {
	switch lessSignificant {
	case node.leftBranch.count <= node.rightBranch.count:
		return node.leftBranch
	case true:
		return node.rightBranch
	case false:
		return node.rightBranch
	default:
		return node.leftBranch
	}
}

func (node *Node) getBranchOrDefault(lessSignificant bool) *Node {
	var branch *Node
	var ok bool

	if branch, ok = node.getDefaultBranch(); !ok {
		branch = node.getBranch(lessSignificant)
	}

	return branch
}

func (node *Node) getDefaultBranch() (*Node, bool) {
	switch {
	case node.leftBranch != nil && node.rightBranch == nil:
		return node.leftBranch, true
	case node.leftBranch == nil && node.rightBranch != nil:
		return node.rightBranch, true
	default:
		return nil, false
	}
}

func addNode(node *Node, leftBranch bool) *Node {
	currentNode := &node.rightBranch
	value := 1
	if leftBranch {
		currentNode = &node.leftBranch
		value = 0
	}

	if *currentNode != nil {
		(*currentNode).count++
	} else {
		*currentNode = &Node{count: 1, value: value}
	}

	return *currentNode
}

func buildBinaryTree() *Tree {
	tree := &Tree{root: new(Node)}

	for _, report := range reports {
		tree.root.count++

		currentNode := tree.root

		for _, value := range report {
			currentNode = addNode(currentNode, value == 0)
		}
	}

	return tree
}

func getRating(root *Node, lessSignificant bool) (int64, error) {
	var rating []rune
	currentNode := root

	for currentNode.leftBranch != nil || currentNode.rightBranch != nil {
		currentNode = currentNode.getBranchOrDefault(lessSignificant)
		rating = append(rating, rune(48 + currentNode.value))
	}

	return strconv.ParseInt(string(rating), 2, 64)
}

func Puzzle2() int64 {
	tree := buildBinaryTree()

	co2Rating, _ := getRating(tree.root, true)
	oxygenRating, _ := getRating(tree.root, false)

	return co2Rating * oxygenRating
}
