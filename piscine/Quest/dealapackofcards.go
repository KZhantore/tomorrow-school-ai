package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	// Количество карт у одного игрока (12 карт / 4 игрока = 3 карты)
	cardsPerPlayer := 3
	for player := 1; player <= 4; player++ {
		// Вычисляем индексы среза для текущего игрока
		startIndex := (player - 1) * cardsPerPlayer
		endIndex := startIndex + cardsPerPlayer
		// Получаем карты текущего игрока
		playerCards := deck[startIndex:endIndex]
		// Выводим имя игрока и его карты в нужном формате
		fmt.Printf("Player %d: %d, %d, %d\n", player, playerCards[0], playerCards[1], playerCards[2])
	}
}
