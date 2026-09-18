package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	current := root
	// Идем до упора влево
	for current.Left != nil {
		current = current.Left
	}
	return current
}
