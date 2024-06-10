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
	out := flag.String("output", "", "this is the ouput file")
	flag.Usage = func() { // define usage
		fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n")
		os.Exit(0)
	}
	flag.Parse()

	for _, arg := range os.Args {
		if arg[len(arg)-1] == '=' {
			flag.Usage()
		}
		if arg[0] == '=' {
			flag.Usage()
		}
		if arg == "--output" {
			flag.Usage()
		}
		if arg == "=" {
			flag.Usage()
		}
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") {
			flag.Usage()
		}
		if strings.HasPrefix(arg, "output") {
			flag.Usage()
		}
	}

	var input string
	var file *os.File
	var outputFile *os.File
	var err error

	// Check command-line arguments
	switch len(os.Args) {
	case 2, 3, 4:
		// Valid cases, do nothing
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	if len(flag.Args()) == 0 { // if string is not provided
		flag.Usage()
	}

	if len(os.Args) == 4 {
		input = flag.Args()[0]
		file, err = os.Open(flag.Args()[1] + ".txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(0)
		}
		outputFile, err = os.Create(*out) //create the file to write program's output
		if err != nil {
			fmt.Printf("Unable to create file %s!\n", *out)
			os.Exit(0)
		}
		defer file.Close()

	}
	if len(os.Args) == 3 {
		if os.Args[len(os.Args)-1] == "thinkertoy" || os.Args[len(os.Args)-1] == "standard" || os.Args[len(os.Args)-1] == "shadow" {
			input = os.Args[1]
			file, err = os.Open(os.Args[2] + ".txt")
			if err != nil {
				fmt.Println("Error opening file:", err)
				os.Exit(0)
			}
			defer file.Close()
			outputFile = nil
		} else {
			input = flag.Args()[0]
			file, err = os.Open("standard.txt")
			if err != nil {
				fmt.Println("Error opening file:", err)
				os.Exit(0)
			}
			outputFile, err = os.Create(*out) //create the file to write program's output
			if err != nil {
				fmt.Printf("Unable to create file %s!\n", *out)
				os.Exit(0)
			}
			defer file.Close()
		}
	}

	if len(os.Args) == 2 && os.Args[1] != *out {
		input = os.Args[1]
		file, err = os.Open("standard.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(0)
		}
		defer file.Close()
		outputFile = nil

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
	asciiart.HandleLn(input, characters, outputFile)
}
