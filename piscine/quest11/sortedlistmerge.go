package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	// Фиктивный (dummy) узел помогает легко построить новый список
	dummy := &NodeI{}
	current := dummy
	// Идем по обоим спискам и выбираем наименьший элемент
	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}
	// Если один из списков закончился, присоединяем оставшуюся часть другого
	if n1 != nil {
		current.Next = n1
	} else {
		current.Next = n2
	}
	return dummy.Next
}
