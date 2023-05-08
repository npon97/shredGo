package main

import (
	"io/ioutil"
	"testing"
)

func TestShred_Directory(t *testing.T) {
	// Create a temporary directory
	tempDir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatal(err)
	}

	// Call Shred function
	err = Shred(tempDir)
	if err == nil {
		t.Error("Shred should have failed for a directory")
	}
}
