package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
    file, err := os.Open("1689007675141_numbers.txt")
    if err != nil {
        fmt.Println("Помилка відкриття файлу:", err)
        return
    }
    defer file.Close()

    phoneRegex := regexp.MustCompile(`(?:\+?1?[-.\s]?)?\(?[0-9]{3}\)?[-.\s]?[0-9]{3}[-.\s]?[0-9]{4}`)

    scanner := bufio.NewScanner(file)

    fmt.Println("Знайдені телефонні номери:")
    
    for scanner.Scan() {
        line := scanner.Text()

        matches := phoneRegex.FindAllString(line, -1)
        for _, match := range matches {
            fmt.Println(match)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Помилка читання файлу:", err)
    }
}