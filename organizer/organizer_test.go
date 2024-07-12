package organizer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOrganize(t *testing.T) {
	testDir := "./testdir"

	// Clean up test directory before starting the test
	if err := os.RemoveAll(testDir); err != nil {
		t.Fatalf("Failed to remove test directory %s: %v", testDir, err)
	}

	if err := os.Mkdir(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory %s: %v", testDir, err)
	}

	defer os.RemoveAll(testDir)

	createEmptyFile := func(name string) {
		data := []byte("")
		err := os.WriteFile(name, data, 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s and %v", name, err)
		}
	}

	testFiles := []string{"test1.txt", "test2.pdf", "test3.txt"}
	for _, file := range testFiles {
		dir := filepath.Join(testDir, file)
		createEmptyFile(dir)
	}

	if err := Organize(testDir, false, false); err != nil {
		t.Fatalf("Organize file failed %v", err)
	}

	for _, file := range testFiles {
		ext := filepath.Ext(file)
		if ext == "" {
			ext = "no_extension"
		} else {
			ext = ext[1:]
		}

		destDir := filepath.Join(testDir, ext)
		destPath := filepath.Join(destDir, file)
		if _, err := os.Stat(destPath); os.IsNotExist(err) {
			t.Errorf("File %s was not moved to %s", file, destDir)
			t.Fatalf("Error: %v", err)
		}
	}
}
