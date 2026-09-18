package piscine

// Определяем структуру food, как указано в условии
type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	// Инициализируем элементы меню с помощью структуры food
	burger := food{preptime: 15}
	chips := food{preptime: 10}
	nuggets := food{preptime: 12}

	// Проверяем входящую строку и возвращаем соответствующее время
	switch order {
	case "burger":
		return burger.preptime
	case "chips":
		return chips.preptime
	case "nuggets":
		return nuggets.preptime
	default:
		// Если товара нет в меню, возвращаем 404
		return 404
	}
}
