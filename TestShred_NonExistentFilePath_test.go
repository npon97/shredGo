package main

import "testing"

func TestShred_NonExistentFilePath(t *testing.T) {
	err := Shred("nonexistentfile")
	if err == nil {
		t.Error("Shred should have failed for a non-existent file")
	}
}
