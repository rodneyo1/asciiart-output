package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func HandleLn(s string, b [][]string, output *os.File) {
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
			output.Write([]byte("\n"))
		}
		return
	}

	// logic for processing lines
	str := strings.Split(s, "\n")
	for _, char := range str {
		if char == "" {
			output.Write([]byte("\n"))
			continue
		}
		Printer(char, b, output)
	}
}
