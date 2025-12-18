package piscine

// TreeNode structure
type TreeNode struct {
	Data   string
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

// Insert data into the binary search tree (RETURNS root)
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

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

// Search for an item in the tree
func BTreeSearchItem(root *TreeNode, item string) *TreeNode {
	if root == nil || root.Data == item {
		return root
	}

	if item < root.Data {
		return BTreeSearchItem(root.Left, item)
	}
	return BTreeSearchItem(root.Right, item)
}

// Replace subtree rooted at node with rplc
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
