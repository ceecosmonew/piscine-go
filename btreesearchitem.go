package piscine

type TreeNode struct {
	Data   interface{}
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func BTreeInsertData(root *TreeNode, data interface{}) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data.(string) < root.Data.(string) {
		child := BTreeInsertData(root.Left, data)
		root.Left = child
		child.Parent = root
	} else {
		child := BTreeInsertData(root.Right, data)
		root.Right = child
		child.Parent = root
	}

	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	data := root.Data.(string)

	if data == elem {
		return root
	}

	if elem < data {
		return BTreeSearchItem(root.Left, elem)
	}

	return BTreeSearchItem(root.Right, elem)
}
