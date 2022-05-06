package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{nil, i, nil}
}

func (std *BinarySearchTree) Insert(i int) {
	if i <= std.data {
		if std.left == nil {
			std.left = &BinarySearchTree{nil, i, nil}
		} else {
			std.left.Insert(i)
		}
	} else {
		if std.right == nil {
			std.right = &BinarySearchTree{nil, i, nil}
		} else {
			std.right.Insert(i)
		}
	}
}

func (std *BinarySearchTree) SortedData() (result []int) {
	result = make([]int, 0)
	var helper func(*BinarySearchTree, []int) []int
	helper = func(std *BinarySearchTree, result []int) []int {
		if std.left != nil {
			result = helper(std.left, result)
		}
		result = append(result, std.data)
		if std.right != nil {
			result = helper(std.right, result)
		}
		return result
	}
	return helper(std, result)
}
