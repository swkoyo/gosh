package main

import (
	"fmt"
	"gosh/helpers"
)

func main() {
	var input string

	fmt.Println("Welcome to Gosh")

	for {
		fmt.Print("> ")
		fmt.Scanln(&input)
		if input == "pwd" {
			pwd, err := helpers.GetCurrDir()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(pwd)
		} else {
			res := fmt.Sprintf("%s is not a valid command", input)
			fmt.Println(res)
		}
	}
}
