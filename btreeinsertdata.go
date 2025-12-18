package piscine

// TreeNode definition (REQUIRED)
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	// If tree is empty, create root
	if root == nil {
		return &TreeNode{Data: data}
	}

	// Insert to the left
	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{
				Data:   data,
				Parent: root,
			}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		// Insert to the right
		if root.Right == nil {
			root.Right = &TreeNode{
				Data:   data,
				Parent: root,
			}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}

	return root
}
