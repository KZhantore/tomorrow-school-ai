package piscine

func ListMerge(l1 *List, l2 *List) {
	// Если второй список пустой, объединять нечего
	if l2.Head == nil {
		return
	}
	// Если первый список пустой, он просто становится копией второго
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}
	// Соединяем хвост первого списка с головой второго списка
	l1.Tail.Next = l2.Head
	// Обновляем хвост первого списка на хвост второго списка
	l1.Tail = l2.Tail
}
