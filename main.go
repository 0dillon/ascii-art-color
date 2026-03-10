package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Invalid number of arguments\nUsage: go run . [OPTION] [STRING]")
		return
	}

	var colorName string
	var subString string
	var inputText string

	if len(os.Args) == 2 {
		inputText = os.Args[1]
	} else if len(os.Args) == 3 {
		colorName = os.Args[1]
		inputText = os.Args[2]
		subString = inputText
	} else {
		colorName = os.Args[1]
		subString = os.Args[2]
		inputText = os.Args[3]
	}

	if colorName != "" {
		if strings.HasPrefix(colorName, "--color=") {
			colorName = strings.TrimPrefix(colorName, "--color=")
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			return
		}
	}

	if inputText == "" {
		return
	}
	fmt.Print(AsciiArt(inputText, colorName, subString))
}

func AsciiArt(input string, colorName string, subString string) string {
	colors := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"reset":   "\033[0m",
	}

	bannerFile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	AsciiChar := strings.Split(string(bannerFile), "\n")
	words := strings.Split(input, "\\n")
	result := ""

	if strings.ReplaceAll(input, "\\n", "") == "" {
		return strings.Repeat("\n", len(words)-1)
	}

	for _, word := range words {
		if word == "" {
			result += "\n"
			continue
		}

		paintMap := make([]bool, len(word))

		if subString == input {
			for j := 0; j < len(word); j++ {
				paintMap[j] = true
			}
		} else if subString != "" {
			for j := 0; j < len(word); j++ {
				if strings.HasPrefix(word[j:], subString) {
					for k := 0; k < len(subString); k++ {
						paintMap[j+k] = true
					}
				}
			}
		}

		for i := 0; i < 8; i++ {
			for index, char := range word {
				if paintMap[index] && colorName != "" {
					result += colors[colorName] + AsciiChar[i+(int(char-' ')*9)+1] + colors["reset"]
				} else {
					result += AsciiChar[i+(int(char-' ')*9)+1]
				}
			}
			result += "\n"
		}
	}

	return result
}
