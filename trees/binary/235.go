package binary

import "fmt"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func Test235() {
	genericNode := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   0,
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
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	a := lowestCommonAncestor(genericNode, &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(a, 6)

	a = lowestCommonAncestor(genericNode, &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(a, 2)

	a = lowestCommonAncestor(genericNode, &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(a, 8)

	a = lowestCommonAncestor(genericNode, &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(a, 2)

	a = lowestCommonAncestor(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}, &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(a, 2)

	a = lowestCommonAncestor(genericNode, &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(a, 4)

	a = lowestCommonAncestor(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}, &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(a, 3)
}
