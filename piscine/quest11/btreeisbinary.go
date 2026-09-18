package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBinaryUtil(root, "", "")
}

func isBinaryUtil(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}
	// Если задана минимальная граница, значение текущего узла должно быть больше или равно ей
	// (В зависимости от того, как в вашей школе обрабатываются дубликаты, обычно >= для правых ветвей)
	if min != "" && node.Data < min {
		return false
	}
	// Если задана максимальная граница, значение текущего узла должно быть строго меньше её
	if max != "" && node.Data >= max {
		return false
	}
	// Для левого поддерева обновляем максимальную границу (текущий узел становится максимумом)
	// Для правого поддерева обновляем минимальную границу (текущий узел становится минимумом)
	return isBinaryUtil(node.Left, min, node.Data) && isBinaryUtil(node.Right, node.Data, max)
}
