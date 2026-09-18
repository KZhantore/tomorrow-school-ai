package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	currentWord := ""

	for i := 0; i < len(str); i++ {
		char := str[i]
		if char == ' ' {
			summary[currentWord]++
			currentWord = ""
		} else {
			currentWord += string(char)
		}
	}
	summary[currentWord]++
	return summary
}
