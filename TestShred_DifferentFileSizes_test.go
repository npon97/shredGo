package main

import (
	"crypto/rand"
	"io/ioutil"
	"os"
	"testing"
)

func TestShred_DifferentFileSizes(t *testing.T) {
	sizes := []int{0, 1, 1024, 1024 * 1024}

	for _, size := range sizes {
		// Create a temporary file
		tempFile, err := ioutil.TempFile("", "testfile")
		if err != nil {
			t.Fatal(err)
		}
		tempFilePath := tempFile.Name()
		tempFile.Close()

		// Write random data to the file
		data := make([]byte, size)
		if _, err := rand.Read(data); err != nil {
			t.Fatal(err)
		}
		err = ioutil.WriteFile(tempFilePath, data, 0644)
		if err != nil {
			t.Fatal(err)
		}

		// Call Shred function
		err = Shred(tempFilePath)
		if err != nil {
			t.Errorf("Shred failed for size %d: %v", size, err)
		}

		// Check if the file was deleted
		_, err = os.Stat(tempFilePath)
		if !os.IsNotExist(err) {
			t.Errorf("File not deleted for size %d: %v", size, err)
		}
	}
}
