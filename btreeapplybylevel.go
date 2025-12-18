package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	height := treeHeight(root)

	for level := 1; level <= height; level++ {
		applyLevel(root, level, f)
	}
}

func treeHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := treeHeight(node.Left)
	right := treeHeight(node.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

func applyLevel(node *TreeNode, level int, f func(...interface{}) (int, error)) {
	if node == nil {
		return
	}

	if level == 1 {
		f(node.Data)
	} else {
		applyLevel(node.Left, level-1, f)
		applyLevel(node.Right, level-1, f)
	}
}
