package reverse_string

func ReverseString(input string) (output string) {
	newStr := make([]rune, len(input))
	index := len(input) - 1

	for _, runeValue := range input {
		newStr[index] = runeValue
		index -= 1
	}
	output = string(newStr)
	return
}
