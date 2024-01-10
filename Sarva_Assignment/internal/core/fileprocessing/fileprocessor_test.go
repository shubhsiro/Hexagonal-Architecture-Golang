// fileprocessor_test.go
package fileprocessing

import (
	"os"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func TestFileProcessor_CalculateSize(t *testing.T) {
	logger := hclog.NewNullLogger()

	// Create a test file for calculating size
	testFilePath := "testfile.txt"
	createTestFile(testFilePath, 1024) // 1024 bytes

	// Create a FileProcessor instance with the test logger
	fileProcessor := NewFileProcessor(logger)

	// Test case 1: Calculate size of an existing file
	size, err := fileProcessor.CalculateSize(testFilePath)

	assert.NoError(t, err, "Should not return an error")
	assert.Equal(t, int64(1024), size, "File size should match")

	// Test case 2: Calculate size of a non-existing file
	nonExistingFilePath := "nonexistingfile.txt"
	size, err = fileProcessor.CalculateSize(nonExistingFilePath)

	assert.Error(t, err, "Should return an error for non-existing file")
	assert.Zero(t, size, "Size should be 0 for non-existing file")

	// Clean up test files
	removeTestFile(testFilePath)
}

func createTestFile(filePath string, size int64) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.Write(make([]byte, size))
	if err != nil {
		panic(err)
	}
}

func removeTestFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		panic(err)
	}
}
