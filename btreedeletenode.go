package piscine

// =====================
// Tree Node Definition
// =====================
type TreeNode struct {
	Data   string
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

// =====================
// Insert Data into BST
// =====================
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
	return root
}

// =====================
// Search Item in BST
// =====================
func BTreeSearchItem(root *TreeNode, item string) *TreeNode {
	if root == nil || root.Data == item {
		return root
	}
	if item < root.Data {
		return BTreeSearchItem(root.Left, item)
	}
	return BTreeSearchItem(root.Right, item)
}

// =====================
// Inorder Traversal
// =====================
func BTreeApplyInorder(root *TreeNode, f func(string)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}

// =====================
// Transplant Subtrees
// =====================
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
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

// =====================
// Delete Node from BST
// =====================
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	// Case 1: no left child
	if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	}

	// Case 2: no right child
	if node.Right == nil {
		return BTreeTransplant(root, node, node.Left)
	}

	// Case 3: two children
	successor := node.Right
	for successor.Left != nil {
		successor = successor.Left
	}

	if successor.Parent != node {
		root = BTreeTransplant(root, successor, successor.Right)
		successor.Right = node.Right
		if successor.Right != nil {
			successor.Right.Parent = successor
		}
	}

	root = BTreeTransplant(root, node, successor)
	successor.Left = node.Left
	if successor.Left != nil {
		successor.Left.Parent = successor
	}

	return root
}

// =====================
// Validate BST
// =====================
func BTreeIsBinary(root *TreeNode) bool {
	return isBinary(root, "", "")
}

func isBinary(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	if (min != "" && node.Data <= min) || (max != "" && node.Data >= max) {
		return false
	}

	return isBinary(node.Left, min, node.Data) &&
		isBinary(node.Right, node.Data, max)
}
