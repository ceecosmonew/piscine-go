package piscine

// TreeNode definition
type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
}

// Insert data into binary search tree
func BTreeInsertData(root *TreeNode, data interface{}) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data.(string) < root.Data.(string) {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}

	return root
}

// Apply function in INORDER traversal
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
