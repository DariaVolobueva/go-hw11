package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
    file, err := os.Open("1689007676028_text.txt")
    if err != nil {
        fmt.Println("Помилка відкриття файлу:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    vowelConsonantRegex := regexp.MustCompile(`(\s|^)[аеєиіїоуюяАЕЄИІЇОУЮЯ][аеєиіїоуюябвгґджзйклмнпрстфхцчшщьАЕЄИІЇОУЮЯБВГҐДЖЗЙКЛМНПРСТФХЦЧШЩЬ]*[бвгґджзйклмнпрстфхцчшщБВГҐДЖЗЙКЛМНПРСТФХЦЧШЩ][\s,.]`)

    vowelConsonantWords := make(map[string]bool)
    repeatedLetterWords := make(map[string]bool)

    for scanner.Scan() {
        line := scanner.Text()

        vcMatches := vowelConsonantRegex.FindAllString(line, -1)
        for _, word := range vcMatches {
            vowelConsonantWords[strings.ToLower(word)] = true
        }


		words := strings.Fields(scanner.Text())
        for _, word := range words {
            cleanWord := strings.Trim(word, ".,!?()\"':;")
            if hasRepeatedLetters(cleanWord) {
                repeatedLetterWords[cleanWord] = true
            }
        }
    }

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

func hasRepeatedLetters(word string) bool {
    runes := []rune(word)
    for i := 0; i < len(runes)-2; i++ {
        if unicode.IsLetter(runes[i]) && unicode.IsLetter(runes[i+2]) && unicode.ToLower(runes[i]) == unicode.ToLower(runes[i+2]) {
            return true
        }
    }
    return false
}