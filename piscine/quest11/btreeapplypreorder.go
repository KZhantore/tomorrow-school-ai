package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	// 1. Применяем функцию к данным текущего узла (сначала корень)
	f(root.Data)
	// 2. Рекурсивно обходим левое поддерево
	BTreeApplyPreorder(root.Left, f)
	// 3. Рекурсивно обходим правое поддерево
	BTreeApplyPreorder(root.Right, f)
}
