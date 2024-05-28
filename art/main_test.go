package main

import (
	"os"
	"strings"
	"testing"
	"io/ioutil"
)

type TestCase struct {
	input          string
	bannerFile     string
	expectedOutput string
}

var testCases = []TestCase{
	{
		input:      "hello",
		bannerFile: "standard",
		expectedOutput: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               `,
	},
	{
		input:      "hello",
		bannerFile: "shadow",
		expectedOutput: `                                 
_|                _| _|          
_|_|_|     _|_|   _| _|   _|_|   
_|    _| _|_|_|_| _| _| _|    _| 
_|    _| _|       _| _| _|    _| 
_|    _|   _|_|_| _| _|   _|_|   
                                 
                                 `,
	},
	{
		input:      "hello",
		bannerFile: "thinkertoy",
		expectedOutput: `   
o        o o     
|        | |     
O--o o-o | | o-o 
|  | |-' | | | | 
o  o o-o o o o-o
`,
	},
	// Add more test cases as needed
}

func TestMainFunction(t *testing.T) {
	for _, testCase := range testCases {
		// Set command line arguments
		os.Args = []string{"program_name", testCase.input, testCase.bannerFile}

		// Run the main function
		main()

		// Read the output from the file
		output, err := ioutil.ReadFile("banner.txt")
		if err != nil {
			t.Fatalf("Failed to read output file: %v", err)
		}

		// Compare output with expected string (trimmed)
		expectedOutputTrimmed := strings.TrimSpace(testCase.expectedOutput)
		outputTrimmed := strings.TrimSpace(string(output))
		if outputTrimmed != expectedOutputTrimmed {
			t.Errorf("Output doesn't match for input '%s' with banner '%s'.\nExpected:\n%s\nGot:\n%s", testCase.input, testCase.bannerFile, expectedOutputTrimmed, outputTrimmed)
		}

		// Clean up the output file
		err = os.Remove("banner.txt")
		if err != nil {
			t.Fatalf("Failed to delete output file: %v", err)
		}
	}
}
