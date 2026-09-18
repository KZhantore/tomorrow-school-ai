package main

import "os"

func main() {
	// os.Args[0] — это путь к самой программе,
	// поэтому реальные аргументы начинаются с индекса 1.
	args := os.Args[1:]
	// Проверяем каждый переданный аргумент
	for _, arg := range args {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			// Если совпадение найдено, выводим сообщение и завершаем программу
			os.Stdout.WriteString("Alert!!!\n")
			return
		}
	}
}
