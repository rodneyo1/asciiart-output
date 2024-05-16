package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func HandleLn(s string, b [][]string) {
	if s == "" {
		fmt.Print()
		return
	}
	// Check if all characters in s are exclusively "\n"
	isAllNewline := true
	for _, char := range s {
		if char != '\n' {
			isAllNewline = false
			break
		}
	}

	if isAllNewline {
		count := strings.Count(s, "\n")
		for i := 0; i < count; i++ {
			fmt.Println()
		}
		return
	}

	//Creating the output file that will be passed to the Printer function
	outputFile, err := os.Create("banner.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	err = Printer(s, b, outputFile)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Successfully wrote to banner.txt")
	}

	// logic for processing lines
	str := strings.Split(s, "\n")
	for _, char := range str {
		if char == "" {
			fmt.Println()
			continue
		}
		Printer(char, b, outputFile)
	}
}
