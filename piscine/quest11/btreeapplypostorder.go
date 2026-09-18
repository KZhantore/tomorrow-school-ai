package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	// 1. Рекурсивно обходим левое поддерево
	BTreeApplyPostorder(root.Left, f)
	// 2. Рекурсивно обходим правое поддерево
	BTreeApplyPostorder(root.Right, f)
	// 3. Применяем функцию к данным текущего узла (после обработки детей)
	f(root.Data)
}
