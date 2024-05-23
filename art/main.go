package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"asciiart"
)

func main() {
	// create an ouput flag
	out := flag.String("output", "banner.txt", "this is the ouput file")
	flag.Parse()

	outputfile, err := os.Create(*out) //create the file to write program's output
	if err != nil {
		fmt.Println("Unable to create file banner.txt!")
		os.Exit(0)
	}
	var input string
	var file *os.File

	// Check command-line arguments
	switch len(os.Args) {
	case 2, 4:
		// Valid cases, do nothing
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] or \nUsage: go run . [STRING]")
		return
	}

	if len(os.Args) == 4 {
		input = flag.Args()[0]
		file, err = os.Open(flag.Args()[1] + ".txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

	}
	if len(os.Args) == 2 {
		input = os.Args[1]
		file, err = os.Open("standard.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

	}

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

	// check for non-printable characters
	for i := 0; i < len(input); i++ {
		if (input[i] < 32 || input[i] > 126) && input[i] != 10 {
			fmt.Println("Cannot print one or more characters")
			return
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
	asciiart.HandleLn(input, characters, outputfile)
	fmt.Println("Art created in banner.txt successfully!")
}
