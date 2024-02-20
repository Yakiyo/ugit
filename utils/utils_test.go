package utils

import (
	"fmt"
	"testing"
)

func TestScanDir(t *testing.T) {
	files, err := ScanDir(".")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(files)
}
