package main

import (
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	os.Args = []string{"program_name", "hello\\tThere12\n3"}

	expectedOutput := ` _              _   _                                   _______   _                                       
| |            | | | |                                 |__   __| | |                           _   ____   
| |__     ___  | | | |   ___                              | |    | |__     ___   _ __    ___  / | |___ \  
|  _ \   / _ \ | | | |  / _ \                             | |    |  _ \   / _ \ | '__|  / _ \ | |   __) | 
| | | | |  __/ | | | | | (_) |                            | |    | | | | |  __/ | |    |  __/ | |  / __/  
|_| |_|  \___| |_| |_|  \___/                             |_|    |_| |_|  \___| |_|     \___| |_| |_____| 
                                                                                                          
                                                                                                          

        
 _____  
|___ /  
  |_ \  
 ___) | 
|____/  
        
        
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
