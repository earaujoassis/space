package utils

import (
	"fmt"
	"os"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func EnsureProjectRoot() error {
	if fileExists("main.go") {
		return nil
	}

	originalDir, _ := os.Getwd()

	for i := 0; i < 5; i++ {
		if err := os.Chdir(".."); err != nil {
			os.Chdir(originalDir)
			return err
		}

		if fileExists("main.go") {
			return nil
		}
	}

	os.Chdir(originalDir)

	return fmt.Errorf("main.go not found in parent directories")
}
