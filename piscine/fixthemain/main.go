package main

import "github.com/01-edu/z01"

const (
	OPEN  = "OPEN"
	CLOSE = "CLOSE"
)

// Объявляем структуру Door с полем state
type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n') // Добавляем перенос строки для красоты и порядка
}

// Изменяем state у переданного указателя на дверь
func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
	return true
}

// Исправлено: добавлено ptrDoor перед state
func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

// Исправлено: функция должна возвращать bool, а присваивание OPEN перенесено выше
func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	return ptrDoor.state == OPEN
}

// Исправлено: добавлен пропущенный return
func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
