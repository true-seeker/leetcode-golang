package binary

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{{root.Val}}

	l := levelOrder(root.Left)
	for i, elem := range l {
		if len(res) < i+2 {
			res = append(res, []int{})
		}
		res[i+1] = append(res[i+1], elem...)
	}

	r := levelOrder(root.Right)
	for i, elem := range r {
		if len(res) < i+2 {
			res = append(res, []int{})
		}
		res[i+1] = append(res[i+1], elem...)
	}

	return res
}

func Test102() {
	a := levelOrder(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
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

	a = levelOrder(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	})
	fmt.Println(a)
}
