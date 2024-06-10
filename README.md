ASCII Art Banner Generator

This program creates ASCII art banners from a provided string and an optional input file containing character sets.

Features

    Accepts a string as input for the banner text.
    Supports loading character sets from a text file for improved control over the banner's appearance.
    Generates ASCII art banners in the specified output file (default: banner.txt).
    Validates input string to ensure it contains only printable characters (except newline \n).
    Handles basic formatting for newline (\n) and tab (\t) characters within the input string.

Installation

    Make sure you have Go installed (https://go.dev/doc/install).
    Clone or download this repository.
    Navigate to the project directory then into the directory containing the main.go file /art in your terminal.

Usage

There are two ways to use the program:

1. With a String Input Only

`go run . "YOUR_STRING"`

This will create an ASCII art banner for the provided string and use the default character set from the built-in standard.txt file. The banner will be saved in banner.txt.

2. With a String Input and an Optional Character Set File

`go run . -output OUTPUT_FILE_NAME STRING CHARACTER_SET_FILE.txt`

    Replace OUTPUT_FILE_NAME with the desired name for the output file (default: banner.txt).
    STRING is the text you want to convert into ASCII art.
    CHARACTER_SET_FILE.txt is the path to the text file containing your custom character set. This file should have one character per line.

Character Set File Format

Each line in the character set file represents a single character that can be used in the banner. The first 8 lines define the top row of characters, the next 8 lines define the second row, and so on, creating a grid of characters. You can customize these characters to create unique banner styles.

Limitations

    The program currently assumes the character set file is 8 characters wide and has 855 lines (8 x 107 characters). Modifying this layout might cause unexpected results.
    The program only accepts printable characters (except newline and tab) for the input string.

Future Improvements

    Support additional formatting options for the input string.

Author

Rodney Otieno

I hope this README provides a clear and informative guide for using your ASCII art banner generator!
