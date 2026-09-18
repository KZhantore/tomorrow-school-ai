package piscine

func ListReverse(l *List) {
	var prev *NodeL
	current := l.Head

	// Сохраняем исходную голову, так как она станет новым хвостом
	oldHead := l.Head

	for current != nil {
		next := current.Next // Запоминаем следующий узел
		current.Next = prev  // Разворачиваем указатель текущего узла назад
		prev = current       // Сдвигаем prev на текущий узел
		current = next       // Переходим к следующему узлу
	}

	// Перенастраиваем указатели Head и Tail самой структуры List
	l.Tail = oldHead
	l.Head = prev
}
