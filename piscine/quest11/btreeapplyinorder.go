package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	// 1. Рекурсивно обходим левое поддерево
	BTreeApplyInorder(root.Left, f)
	// 2. Применяем функцию к данным текущего узла
	f(root.Data)
	// 3. Рекурсивно обходим правое поддерево
	BTreeApplyInorder(root.Right, f)
}
