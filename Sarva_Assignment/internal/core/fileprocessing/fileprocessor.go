package fileprocessing

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-hclog"
)

// FileProcessor handles the logic for processing files.
type FileProcessor struct {
	logger hclog.Logger
}

// NewFileProcessor creates a new instance of FileProcessor with a provided logger.
func NewFileProcessor(logger hclog.Logger) *FileProcessor {
	return &FileProcessor{logger: logger}
}

// CalculateSize calculates the size of the file in bytes.
func (fp *FileProcessor) CalculateSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fp.logger.Error("File not found", "path", filePath)

		return 0, fmt.Errorf("file not found: %s", filePath)
	} else if err != nil {
		fp.logger.Error("Error getting file information", "path", filePath, "error", err)

		return 0, fmt.Errorf("error getting file information: %v", err)
	}

	fp.logger.Debug("File size calculated", "path", filePath, "size", fileInfo.Size())

	return fileInfo.Size(), nil
}
