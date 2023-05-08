package test

import (
	"io/ioutil"
	"testing"
)

func TestShred_ReadOnlyFile(t *testing.T) {
	// Create a temporary file
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	tempFilePath := tempFile.Name()
	tempFile.Close()

	// Write some data to the file
	err = ioutil.WriteFile(tempFilePath, []byte("test data"), 0444)
	if err != nil {
		t.Fatal(err)
	}

	// Call Shred function
	err = Shred(tempFilePath)
	if err == nil {
		t.Error("Shred should have failed for a read-only file")
	}
}
