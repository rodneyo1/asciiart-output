package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func HandleLn(s string, b [][]string, outputfile *os.File) {
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
	if outputfile == nil {
		if isAllNewline {
			count := strings.Count(s, "\n")
			for i := 0; i < count; i++ {
				fmt.Println()
			}
			return
		}

		// logic for processing lines
		str := strings.Split(s, "\n")
		for _, char := range str {
			if char == "" {
				fmt.Println()
				continue
			}
			Printer(char, b, outputfile)
		}
	} else {
		if isAllNewline {
			count := strings.Count(s, "\n")
			for i := 0; i < count; i++ {
				outputfile.Write([]byte("\n"))
			}
			return
		}

		// logic for processing lines
		str := strings.Split(s, "\n")
		for _, char := range str {
			if char == "" {
				outputfile.Write([]byte("\n"))
				continue
			}
			Printer(char, b, outputfile)
		}
	}
}
