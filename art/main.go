package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"asciiart"
)

func main() {
	// Open the file
	file, err := os.Open("shadow.txt")

	if len(os.Args) == 4 {
		bannerName := os.Args[3] + ".txt"
		file, err = os.Open(bannerName)
		if err != nil {
			fmt.Println(err, ". Please specify a valid banner file.")
			return
		}
	}

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read lines from the file
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) != 855 {
		fmt.Printf("Can't Print. The banner file %s has been modified\n", file.Name())
		return
	}

	// Split lines into characters
	characters := [][]string{}
	for i := 0; i < len(lines); i += 9 {
		end := i + 9
		if end > len(lines) {
			end = len(lines)
		}
		characters = append(characters, lines[i:end])
	}

	// Check command-line arguments
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	// Check for non-printable characters in the input string
	input := os.Args[1]

	if len(os.Args) == 4 {
		input = os.Args[2]
	}

	for i := 0; i < len(input); i++ {
		if (input[i] < 32 || input[i] > 126) && input[i] != 10 {
			fmt.Println("Cannot print one or more characters")
			return
		}
	}

	// Format ("\n") in input string
	input = strings.Replace(input, "\\n", "\n", -1)
	input = strings.Replace(input, "\\t", "    ", -1)

	// Generate ASCII art
	asciiart.HandleLn(input, characters)
}
