package asciiart

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GenerateArt(input, bannerFile, outputFile string) error {
	file, err := os.Open(bannerFile + ".txt")
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	outputfile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("unable to create file %s: %v", outputFile, err)
	}
	defer outputfile.Close()

	// Read lines from the file
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) != 855 {
		return fmt.Errorf("can't Print. The banner file %s has been modified", file.Name())
	}

	// Check for non-printable characters
	for i := 0; i < len(input); i++ {
		if (input[i] < 32 || input[i] > 126) && input[i] != 10 {
			return fmt.Errorf("cannot print one or more characters")
		}
	}

	// Format ("\n") in input string
	input = strings.Replace(input, "\\n", "\n", -1)
	input = strings.Replace(input, "\\t", "    ", -1)

	// Split lines into characters of 8 lines
	characters := [][]string{}
	for i := 1; i < len(lines); i += 9 {
		end := i + 9
		if end > len(lines) {
			end = len(lines)
		}
		characters = append(characters, lines[i:end])
	}

	// Generate ASCII art
	HandleLn(input, characters, outputfile)
	fmt.Println("Art created in banner.txt successfully!")
	return nil
}
