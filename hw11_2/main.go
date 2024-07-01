package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
    // Відкриваємо файл
    file, err := os.Open("1689007676028_text.txt")
    if err != nil {
        fmt.Println("Помилка відкриття файлу:", err)
        return
    }
    defer file.Close()

    // Створюємо сканер для читання файлу
    scanner := bufio.NewScanner(file)

    vowelConsonantWords := make(map[string]bool)
    repeatedLetterWords := make(map[string]bool)

    // Читаємо файл рядок за рядком
    for scanner.Scan() {
        words := strings.Fields(scanner.Text())
        for _, word := range words {
            cleanWord := strings.Trim(word, ".,!?()\"':;")
            if startsWithVowelEndsWithConsonant(cleanWord) {
                vowelConsonantWords[strings.ToLower(cleanWord)] = true
            }
            if hasRepeatedLetters(cleanWord) {
                repeatedLetterWords[cleanWord] = true
            }
        }
    }

    // Перевіряємо помилки сканування
    if err := scanner.Err(); err != nil {
        fmt.Println("Помилка читання файлу:", err)
    }

    fmt.Println("Слова, що починаються з голосної і закінчуються на приголосну:")
    for word := range vowelConsonantWords {
        fmt.Println(word)
    }

    fmt.Println("\nСлова з двома однаковими буквами, розділеними будь-яким символом:")
    for word := range repeatedLetterWords {
        fmt.Println(word)
    }
}

func startsWithVowelEndsWithConsonant(word string) bool {
    if len(word) < 2 {
        return false
    }
    runes := []rune(word)
    firstLetter := unicode.ToLower(runes[0])
    lastLetter := unicode.ToLower(runes[len(runes)-1])
    
    vowels := []rune{'а', 'е', 'є', 'и', 'і', 'ї', 'о', 'у', 'ю', 'я'}
    consonants := []rune{'б', 'в', 'г', 'ґ', 'д', 'ж', 'з', 'й', 'к', 'л', 'м', 'н', 'п', 'р', 'с', 'т', 'ф', 'х', 'ц', 'ч', 'ш', 'щ'}

    startsWithVowel := false
    for _, vowel := range vowels {
        if firstLetter == vowel {
            startsWithVowel = true
            break
        }
    }

    endsWithConsonant := false
    for _, consonant := range consonants {
        if lastLetter == consonant {
            endsWithConsonant = true
            break
        }
    }

    return startsWithVowel && endsWithConsonant
}

func hasRepeatedLetters(word string) bool {
    runes := []rune(word)
    for i := 0; i < len(runes)-2; i++ {
        if unicode.IsLetter(runes[i]) && unicode.IsLetter(runes[i+2]) && unicode.ToLower(runes[i]) == unicode.ToLower(runes[i+2]) {
            return true
        }
    }
    return false
}