package helpers

import "os"

func GetCurrDir() (string, error) {
	return os.Getwd()
}

func ChangeDir(dir string) error {
	return os.Chdir(dir)
}
