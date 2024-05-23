package asciiart

import (
	"os"
)

func Printer(s string, b [][]string, outputFile *os.File) {
	for i := 0; i < 9; i++ {
		for _, char := range s {
			toPrint := char - 32
			// Write the character to the file instead of printing
			outputFile.Write([]byte(b[toPrint][i]))
		}
		// Write a newline character to the file
		outputFile.Write([]byte("\n"))
	}
}
