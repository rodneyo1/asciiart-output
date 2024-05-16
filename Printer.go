package asciiart

import (
	"fmt"
	"os"
)

func Printer(s string, b [][]string, outputFile *os.File) error {
	// Check if outputFile is valid (optional)
	if outputFile == nil {
		return fmt.Errorf("invalid output file")
	}

	defer outputFile.Close()

	for i := 1; i < 9; i++ {
		for _, char := range s {
			toPrint := char - 32
			// Write the character to the file instead of printing
			_, err := outputFile.Write([]byte(b[toPrint][i]))

			if err != nil {
				return err
			}
		}
		// Write a newline character to the file
		_, err := outputFile.Write([]byte("\n"))

		if err != nil {
			return err
		}
	}
	return nil
}
