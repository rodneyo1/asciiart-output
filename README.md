# ASCII Art Printer

This Go program prints ASCII characters based on a provided file containing ASCII patterns. The file standard.txt should contain ASCII patterns for characters from space " " to tilde "~". The program takes two arguments to run:

    The name of the file containing the main function (art)
    A string as input

Example command:

bash

go run main.go "Hello" | cat -e

How to Use

    Ensure you have Go installed on your system.
    Place your ASCII art patterns in a file named standard.txt.
    Navigate to the folder, art, containing main.go
    Run the program with the main file name and a string input.

Code Explanation

The program consists of the following main functions:
main()

    Reads the contents of standard.txt into memory.
    Converts the input string argument to the desired format.
    Handles line breaks and prints the corresponding ASCII art.

HandleLn(s string, b [][]string)

    Handles newline characters in the input string.
    Prints ASCII art for each character in the input string.

Printer(s string, b [][]string)

    Prints the ASCII art for each character in the input string.

Usage Notes

    Ensure the standard.txt file contains ASCII patterns for characters from space " " to tilde "~".
    The program may not handle characters outside this range gracefully and will print an error message if encountered.

Feel free to enhance the program to handle a wider range of characters or improve error handling as needed.
