package helpers

import (
	"bufio"
	"os"
)

func GetFile() ([]byte, error) {
	var path = getPath()
	return os.ReadFile(path)
}

func getPath() string {
	println("write the path to file")
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
