package piscine

// PointOne принимает указатель на int и присваивает значение 1 переменной, на которую он указывает.
func PointOne(n *int) {
	*n = 1
}
