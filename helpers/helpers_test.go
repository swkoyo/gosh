package helpers

import (
	"os"
	"testing"
)

func TestGetCurrDir(t *testing.T) {
	currDir, err := os.Getwd()
	if err != nil {
		t.Error("Error while getting current directory")
	}
	res, err := GetCurrDir()
	if err != nil {
		t.Error("Error while getting current directory")
	}
	if res != currDir {
		t.Fatalf("Expected %s, got %s", currDir, res)
	}
}
