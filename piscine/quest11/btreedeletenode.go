package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}
	// Сценарий 1: У узла нет левого потомка
	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)
		// Сценарий 2: У узла нет правого потомка
	} else if node.Right == nil {
		root = BTreeTransplant(root, node, node.Left)
		// Сценарий 3: У узла есть оба потомка
	} else {
		// Ищем преемника (минимальный узел в правом поддереве)
		successor := node.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		// Если преемник находится глубже, перестраиваем его связи
		if successor.Parent != node {
			root = BTreeTransplant(root, successor, successor.Right)
			successor.Right = node.Right
			successor.Right.Parent = successor
		}
		// Заменяем удаляемый узел преемником
		root = BTreeTransplant(root, node, successor)
		successor.Left = node.Left
		successor.Left.Parent = successor
	}
	return root
}
