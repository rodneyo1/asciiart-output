package main

import (
	"asciiart"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Create an output flag
	out := flag.String("output", "banner.txt", "this is the output file")
	flag.Parse()

	// Check command-line arguments
	if len(os.Args) != 2 && len(os.Args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	var input, bannerFile string
	if len(os.Args) == 4 {
		input = flag.Args()[0]
		bannerFile = flag.Args()[1]
	} else if len(os.Args) == 2 {
		input = os.Args[1]
		bannerFile = "standard"
	}

	err := asciiart.GenerateArt(input, bannerFile, *out)
	if err != nil {
		fmt.Println(err)
	}
}
