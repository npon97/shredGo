package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	Shred("./randomfile")
}

func Shred(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	// Check if the file is read-only
	if fileInfo.Mode().Perm()&0200 == 0 {
		return fmt.Errorf("file is read-only: %s", path)
	}

	fileSize := fileInfo.Size()
	data := make([]byte, fileSize)

	for i := 0; i < 3; i++ {
		if _, err := rand.Read(data); err != nil {
			return err
		}

		if err := ioutil.WriteFile(path, data, fileInfo.Mode()); err != nil {
			return err
		}
	}

	return os.Remove(path)
}
