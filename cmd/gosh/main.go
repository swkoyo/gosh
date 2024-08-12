package main

import (
	"bufio"
	"fmt"
	"gosh/helpers"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	ps1 := helpers.GetPS1()

	fmt.Println("Welcome to Gosh")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		for {
			sig := <-sigChan
			if sig == syscall.SIGINT {
				fmt.Println("\nReceieved SIGINT, exiting gracefully...")
				os.Exit(0)
			}
		}
	}()

	if term.IsTerminal(int(os.Stdin.Fd())) {
		for {
			fmt.Print(ps1)
			line, err := inputReader.ReadString('\n')
			line = strings.TrimSpace(line)
			if err != nil {
				log.Fatal(err)
			}
			handleCommand(line)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			handleCommand(line)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func handleCommand(line string) {
	if strings.Contains(line, ">") || strings.Contains(line, "<") {
		parts := strings.Fields(line)
		var cmd, inputFile, outputFile string

		for i := 0; i < len(parts); i++ {
			if parts[i] == ">" {
				if i+1 < len(parts) {
					outputFile = parts[i+1]
					break
				}
			} else if parts[i] == "<" {
				if i+1 < len(parts) {
					inputFile = parts[i+1]
					break
				}
			} else {
				cmd += parts[i] + " "
			}
		}

		cmd = strings.TrimSpace(cmd)

		if inputFile != "" {
			file, err := os.Open(inputFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				return
			}
			defer file.Close()
			os.Stdin = file
		}

		if outputFile != "" {
			file, err := os.Create(outputFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				return
			}
			defer file.Close()
			os.Stdout = file
		}

		handleCommand(cmd)

		if inputFile != "" {
			os.Stdin = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
		}

		if outputFile != "" {
			os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
		}

		return
	}

	if line == "pwd" {
		pwd, err := helpers.GetCurrDir()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(pwd)
	} else if line == "exit" {
		fmt.Println("Goodbye!")
		os.Exit(0)
	} else if strings.HasPrefix(line, "cd") {
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Invalid cd command\n")
			return
		}
		path := parts[1]
		if path[0] != '/' {
			currentDir, err := helpers.GetCurrDir()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to change dir\n")
				return
			}
			path = fmt.Sprintf("%s/%s", currentDir, path)
		}
		err := helpers.ChangeDir(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			return
		}
	} else if strings.HasPrefix(line, "echo") {
		parts := strings.SplitN(line, " ", 2)
		if len(parts) > 1 {
			fmt.Println(parts[1])
		}
	} else if line == "ls" {
		files, err := os.ReadDir(".")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			return
		}
		var filenames []string
		for _, file := range files {
			filenames = append(filenames, file.Name())
		}
		fmt.Println(strings.Join(filenames, "  "))
	} else if strings.HasPrefix(line, "cat") {
		parts := strings.SplitN(line, " ", 2)
		if len(parts) == 1 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
			}
		} else if len(parts) == 2 {
			filename := parts[1]
			data, err := os.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				return
			}
			fmt.Print(string(data))
		} else {
			fmt.Println("Usage: cat <filename>")
		}
	} else {
		res := fmt.Sprintf("%s is not a valid command", line)
		fmt.Println(res)
	}
}
