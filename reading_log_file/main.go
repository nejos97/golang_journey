package main

import (
	"bufio"
	"fmt"
	"os"
)

func printLastNLines(lines []string, num int) []string {
	var textToPrint []string
	for i := len(lines) - num; i < len(lines); i++ {
		textToPrint = append(textToPrint, lines[i])
	}
	return textToPrint
}

func main() {
	filepath := "log.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error occur while opening the file", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file", err)
		return
	}

	result := printLastNLines(lines, 3)
	for _, line := range result {
		fmt.Println(line)
		fmt.Println("**********")
	}
}
