package helpers

import (
	"fmt"
	"os"
)

func GetCurrDir() (string, error) {
	return os.Getwd()
}

func ChangeDir(dir string) error {
	return os.Chdir(dir)
}

func GetPS1() string {
	ps1 := os.Getenv("PS1")
	if ps1 == "" {
		ps1 = ">"
	}
	return fmt.Sprintf("%s ", ps1)
}
