package helpers

import "os"

func GetCurrDir() (string, error) {
	return os.Getwd()
}
