package binary

import (
	"fmt"
	"math"
)

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinBSTValue(root *TreeNode) int {
	leftMin := math.MaxInt
	rightMin := math.MaxInt
	if root.Left != nil {
		leftMin = MinBSTValue(root.Left)
	}
	if root.Right != nil {
		rightMin = MinBSTValue(root.Right)
	}

	return MinInt(root.Val, MinInt(leftMin, rightMin))
}

func MaxBSTValue(root *TreeNode) int {
	leftMax := -math.MaxInt
	rightMax := -math.MaxInt
	if root.Left != nil {
		leftMax = MaxBSTValue(root.Left)
	}
	if root.Right != nil {
		rightMax = MaxBSTValue(root.Right)
	}

	return MaxInt(root.Val, MaxInt(leftMax, rightMax))
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && MaxBSTValue(root.Left) >= root.Val {
		return false
	}

	if root.Right != nil && MinBSTValue(root.Right) <= root.Val {
		return false
	}

	if !isValidBST(root.Left) || !isValidBST(root.Right) {
		return false
	}
	return true
}

func Test98() {
	a := isValidBST(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	})
	fmt.Println(a)

	a = isValidBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	})
	fmt.Println(a)

	a = isValidBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	})
	fmt.Println(a)
}
