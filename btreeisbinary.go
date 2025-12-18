package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, "", "")
}

func isBST(node *TreeNode, min string, max string) bool {
	if node == nil {
		return true
	}

	value := node.Data.(string)

	if min != "" && value < min {
		return false
	}
	if max != "" && value >= max {
		return false
	}

	return isBST(node.Left, min, value) &&
		isBST(node.Right, value, max)
}
