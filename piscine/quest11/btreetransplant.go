package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}
	// 1. Если заменяемый узел — это корень дерева
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		// 2. Если узел является левым ребенком своего родителя
		node.Parent.Left = rplc
	} else {
		// 3. Если узел является правым ребенком своего родителя
		node.Parent.Right = rplc
	}
	// Обновляем ссылку на родителя у вставляемого поддерева, если оно не nil
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
