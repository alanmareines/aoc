package utils

import (
	"os"
)

func InputToString(path string) string {
	data, err := os.ReadFile(path)
	Check(err)
	return string(data)
}
