# ASCII Art Output Generator

This is a simple command-line program that generates ASCII art from a given string using a specified banner. The program supports three different banners: `standard`, `shadow`, and `thinkertoy`. Additionally, the output can be saved to a file.
```
                         _    _          _   _          
                        | |  | |        | | | |         
                        | |__| |   ___  | | | |   ___   
                        |  __  |  / _ \ | | | |  / _ \  
                        | |  | | |  __/ | | | | | (_) | 
                        |_|  |_|  \___| |_| |_|  \___/  
                                                        
                                                        
```

## Installation

1. Make sure you have Go installed on your machine. You can download and install Go from [here](https://golang.org/dl/).

2. Clone this repository or download the source code files.

3. Navigate to the project directory.

## Usage

### Build the Program

Before running the program, you need to build it. Use the following command to build the executable:

```sh
go build -o ascii-art-generator
```

This will create an executable file named `ascii-art-generator` in the project directory.

### Run the Program

To run the program, use the following command:

```sh
./ascii-art-generator [OPTION] [STRING] [BANNER]
```

### Examples

1. Generate ASCII art and print it to the console:

    ```sh
    ./ascii-art-generator "Hello, World!" standard
    ```

2. Generate ASCII art and save it to a file:

    ```sh
    ./ascii-art-generator --output=output.txt "Hello, World!" standard
    ```

3. Use a different banner:

    ```sh
    ./ascii-art-generator "Hello, World!" shadow
    ./ascii-art-generator "Hello, World!" thinkertoy
    ```

### Options

- `--output=<fileName.txt>`: Specify the output file where the ASCII art will be saved.

## Banner Files

The program relies on banner files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`) to generate the ASCII art. Make sure these files are in the same directory as the executable.

## License

This project is licensed under the Apache License.