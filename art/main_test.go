package main

import (
	"asciiart"
	"os"
	"strings"
	"testing"
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

func TestGenerateArt(t *testing.T) {
	for _, testCase := range testCases {
		outputFile := "test_banner.txt"
		err := asciiart.GenerateArt(testCase.input, testCase.bannerFile, outputFile)
		if err != nil {
			t.Fatalf("generateArt failed: %v", err)
		}

		// Read the output from the file
		output, err := os.ReadFile(outputFile)
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
		err = os.Remove(outputFile)
		if err != nil {
			t.Fatalf("Failed to delete output file: %v", err)
		}
	}
}
