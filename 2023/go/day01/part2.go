package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func contains_num(s string)int{
	if strings.Contains(s, "1") || strings.Contains(s, "one") {
		return 1
	}else if strings.Contains(s, "2") || strings.Contains(s, "two") {
		return 2
	}else if strings.Contains(s, "3") || strings.Contains(s, "three") {
		return 3
	}else if strings.Contains(s, "4") || strings.Contains(s, "four") {
		return 4
	}else if strings.Contains(s, "5") || strings.Contains(s, "five") {
		return 5
	}else if strings.Contains(s, "6") || strings.Contains(s, "six") {
		return 6
	}else if strings.Contains(s, "7") || strings.Contains(s, "seven") {
		return 7
	}else if strings.Contains(s, "8") || strings.Contains(s, "eight") {
		return 8
	}else if strings.Contains(s, "9") || strings.Contains(s, "nine") {
		return 9
	}else {
		return 0
	}
}

func part2() int{
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		value := 0
		for i := 1; i <= len(text); i++{
			result := contains_num(text[:i])
			if result != 0 {
				value = result * 10
				break
			}
		}
		for i := 1; i <= len(text); i++{
			result := contains_num(text[len(text) - i:])
			if result != 0 {
				value += result
				break
			}
		}
		total += value
	}
	return total
}
