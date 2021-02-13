package main

import (
	"os"
)

func isExistingDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func arrayContains(stack []string, needle string) bool {
	for _, e := range stack {
		if e == needle {
			return true
		}
	}
	return false
}