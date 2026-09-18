package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// Рекурсивно считаем количество уровней в левом и правом поддеревьях
	leftCount := BTreeLevelCount(root.Left)
	rightCount := BTreeLevelCount(root.Right)
	// Находим максимум из двух поддеревьев и прибавляем 1 (текущий уровень)
	if leftCount > rightCount {
		return leftCount + 1
	}
	return rightCount + 1
}
