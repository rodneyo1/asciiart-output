package main

import (
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	os.Args = []string{"program_name", "hello"}

	expectedOutput := ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               `

	// Redirect stdout to capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the main function
	main()

	// Reset stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var output []byte
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		output = append(output, buf[:n]...)
		if err != nil {
			break
		}
	}

	// Compare output with expected string (trimmed)
	expectedOutputTrimmed := strings.TrimSpace(expectedOutput)
	outputTrimmed := strings.TrimSpace(string(output))
	if outputTrimmed != expectedOutputTrimmed {
		t.Errorf("Output doesn't match. Expected:\n%s\nGot:\n%s", expectedOutputTrimmed, outputTrimmed)
	}
}
