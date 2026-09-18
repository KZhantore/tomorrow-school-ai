package piscine

func LoafOfBread(str string) string {
	// Edge Case: If the input is completely empty, return just a newline
	if len(str) == 0 {
		return "\n"
	}

	// 1. Calculate length of valid characters
	validCount := 0
	for _, r := range str {
		if r != ' ' {
			validCount++
		}
	}

	// Edge Case: If it contains only spaces, it's also considered empty/valid to return "\n"
	if validCount == 0 {
		return "\n"
	}

	// If there are characters but fewer than 5, it's an invalid loaf
	if validCount < 5 {
		return "Invalid Output\n"
	}

	var result []rune
	runes := []rune(str)
	i := 0
	firstBlock := true

	for i < len(runes) {
		var block []rune

		// Gather up to 5 non-space characters
		for i < len(runes) && len(block) < 5 {
			if runes[i] != ' ' {
				block = append(block, runes[i])
			}
			i++
		}

		if len(block) == 0 {
			break
		}

		if !firstBlock {
			result = append(result, ' ')
		}
		firstBlock = false
		result = append(result, block...)

		// Skip exactly one character index position
		i++
	}

	// Clean up any trailing space that might have slipped in before the newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	result = append(result, '\n')
	return string(result)
}
