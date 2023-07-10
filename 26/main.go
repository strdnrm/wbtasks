package main

import (
	"fmt"
)

// Функция проверки уникальности символов в строке
func isUnique(str string) bool {
	// Создаем map для хранения символов
	charMap := make(map[rune]struct{})

	// Проходим по каждому символу в строке
	for _, char := range str {
		// Проверяем есть ли символ в map
		if _, found := charMap[char]; found {
			return false
		}

		// Добавляем символ в map
		charMap[char] = struct{}{}
	}

	return true
}

func main() {
	str := "abcd"
	fmt.Println(isUnique(str))

	str = "abCdefAaf"
	fmt.Println(isUnique(str))

	str = "aabcd"
	fmt.Println(isUnique(str))
}
