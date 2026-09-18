package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}
	// 1. Если список пуст или новый элемент должен быть самым первым
	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}
	// 2. Ищем место для вставки в середине или в конце списка
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}
	// Вставляем новый узел между current и current.Next
	newNode.Next = current.Next
	current.Next = newNode
	return l
}
