package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	// 1. Удаляем совпадающие элементы из начала списка (сдвигаем Head)
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	// Если после этого список стал пустым, сбрасываем и Tail
	if l.Head == nil {
		l.Tail = nil
		return
	}
	// 2. Удаляем совпадающие элементы из середины и конца списка
	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	// 3. Обновляем указатель Tail на актуальный последний элемент
	l.Tail = current
}
