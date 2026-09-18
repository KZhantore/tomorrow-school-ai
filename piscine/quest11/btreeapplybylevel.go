package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	// Инициализируем очередь, помещая туда корень дерева
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// Извлекаем первый элемент из очереди
		current := queue[0]
		queue = queue[1:]
		// Применяем функцию к данным текущего узла
		f(current.Data)
		// Если есть левый потомок, добавляем его в очередь
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		// Если есть правый потомок, добавляем его в очередь
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}
