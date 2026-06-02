package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	// Создаём пустую карту для подсчёта слов
	counts := make(map[string]int)
	
	// Разбиваем строку на слова по пробелам
	words := strings.Fields(s)
	
	// Подсчитываем каждое слово
	for _, word := range words {
		counts[word]++
	}
	
	return counts
}

// Собственная функция для тестирования (аналог wc.Test)
func testWordCount() {
	testCases := []string{
		"I am learning Go",
		"Go Go Go",
		"hello world hello",
		"",
		"one",
		"the quick brown fox jumps over the lazy dog",
	}
	
	for _, test := range testCases {
		result := WordCount(test)
		fmt.Printf("Input: %q\n", test)
		fmt.Printf("Output: %v\n\n", result)
	}
}

func main() {
	fmt.Println("=== Тестирование WordCount ===\n")
	testWordCount()
	
	// Интерактивный режим
	fmt.Println("Введите строку для подсчёта слов (или 'exit' для выхода):")
	var input string
	for {
		fmt.Print("> ")
		fmt.Scanln(&input)
		
		if input == "exit" {
			break
		}
		
		result := WordCount(input)
		fmt.Printf("Результат: %v\n\n", result)
	}
}