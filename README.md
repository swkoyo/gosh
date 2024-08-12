# Gosh

Gosh is a simple shell program written in Go that supports basic command execution, input/output redirection, and signal handling. It's designed to be a lightweight and easy-to-use shell for learning and experimentation.

## Features
### Basic Shell Commands:
- `pwd`: Print the current working directory.
- `exit`: Exit the shell.
- `ls`: List files in the current directory.
- `cd` <directory>: Change the current working directory.
- `echo` <text>: Display a line of text.
- `cat` <filename>: Display the contents of a file.
### STDIN/STDOUT Redirection:
- Redirect input from a file: `cat < input.txt`
- Redirect output to a file: echo `Hello > output.txt`
### Signal Handling:
- `SIGINT` (Ctrl+C): Interrupts the current command but keeps the shell running.
- `SIGTERM`: Gracefully shuts down the shell.


## Installation
Ensure you have Go installed on your system. Clone the repository and build the project:

```bash
git clone https://github.com/swkoyo/gosh.git
cd gosh
make build
```

## Running tests
You can run tests using the following command:
```bash
# Running tests
make test
```

## Running the Program
You can run Gosh using the following command:
```bash
# Running the shell
make run
```

## Usage Examples
Here are some examples of how to use Gosh:

```bash
# Print the current working directory
pwd

# List files in the directory
ls

# Change directory
cd /path/to/directory

# Display text
echo "Hello, Gosh!"

# Display the contents of a file
cat file.txt

# Redirect output to a file
echo "Hello, Gosh!" > output.txt

# Redirect input from a file
cat < input.txt
```

## Error Handling
- Invalid Commands: Gosh will display an error message if you enter an unrecognized command.
- Signal Handling:
    - Ctrl+C (SIGINT): Interrupts the current operation and returns control to the shell.
    - Ctrl+D (EOF): Gracefully exits the shell with a custom message.
