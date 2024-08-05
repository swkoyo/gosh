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

	for {
		fmt.Print(ps1)
		line, err := inputReader.ReadString('\n')
		line = line[:len(line)-1]
		if err != nil {
			log.Fatal(err)
		}
		handleCommand(line)
	}
}

func handleCommand(line string) {
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
	} else {
		res := fmt.Sprintf("%s is not a valid command", line)
		fmt.Println(res)
	}
}
