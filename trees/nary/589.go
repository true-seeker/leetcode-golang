package nary

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}

	for i := 0; i < len(root.Children); i++ {
		res = append(res, preorder(root.Children[i])...)
	}
	return res
}

func Test589() {
	a := preorder(&Node{
		Val: 1,
		Children: []*Node{{
			Val: 3,
			Children: []*Node{{
				Val:      5,
				Children: nil,
			}, {
				Val:      6,
				Children: nil,
			}},
		}, {
			Val:      2,
			Children: nil,
		}, {
			Val:      4,
			Children: nil,
		}},
	})
	fmt.Println(a)
}
