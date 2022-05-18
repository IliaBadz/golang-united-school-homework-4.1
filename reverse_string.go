package reverse_string

import "fmt"

func ReverseString(input string) (output string) {
	newStr := []rune(input)
	fmt.Println(newStr)
	index := len(newStr) - 1

	for i, runeValue := range input {
		newStr[index] = runeValue
		fmt.Println(input[i], runeValue)
		index -= 1
	}

	return string(newStr)
}

// runes := []rune(s)
//     for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//         runes[i], runes[j] = runes[j], runes[i]
//     }
//     return string(runes)
