package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {

		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Left == node {
		root.Left = rplc
		root.Left.Parent = root

	} else if root.Right != nil && root.Right == node {
		root.Right = rplc
		root.Right.Parent = root
	} else {
		root.Left = BTreeTransplant(root.Left, node, rplc)
		root.Right = BTreeTransplant(root.Right, node, rplc)
	}
	return root
}
