package handler

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"Sarva_Assignment/internal/adapter/consensus"
	"Sarva_Assignment/internal/adapter/database"
	"Sarva_Assignment/internal/adapter/logger"
	"Sarva_Assignment/internal/core/fileprocessing"
)

// FileHandler handles HTTP requests related to file operations.
type FileHandler struct {
	fileProcessor    *fileprocessing.FileProcessor
	consensusAdapter *consensus.RaftAdapter
	databaseAdapter  *database.RedisAdapter
	loggerAdapter    *logger.HCLoggerAdapter
}

// NewFileHandler creates a new instance of FileHandler.
func NewFileHandler(fp *fileprocessing.FileProcessor, ca *consensus.RaftAdapter, da *database.RedisAdapter, la *logger.HCLoggerAdapter) *FileHandler {
	return &FileHandler{
		fileProcessor:    fp,
		consensusAdapter: ca,
		databaseAdapter:  da,
		loggerAdapter:    la,
	}
}

// UploadFile handles the file upload HTTP endpoint.
func (fh *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")

	if err != nil {
		fh.loggerAdapter.Log(fmt.Sprintf("Error handling file upload: %s", err.Error()))
		http.Error(w, "Error handling file upload", http.StatusBadRequest)

		return
	}
	defer file.Close()

	// Validate file size and extension
	if err := validateFile(handler); err != nil {
		fh.loggerAdapter.Log(fmt.Sprintf("Invalid file: %s", err.Error()))
		http.Error(w, "Invalid file", http.StatusBadRequest)

		return
	}

	// Calculate file size
	fileSize, err := fh.fileProcessor.CalculateSize(handler.Filename)
	if err != nil {
		fh.loggerAdapter.Log(fmt.Sprintf("Error calculating file size: %s", err.Error()))
		http.Error(w, "Error calculating file size", http.StatusInternalServerError)

		return
	}

	// Send message to RAFT cluster for consensus
	if err := fh.consensusAdapter.SendMessage(fmt.Sprintf("Update Redis: %s - %d bytes", handler.Filename, fileSize)); err != nil {
		fh.loggerAdapter.Log(fmt.Sprintf("Error sending message to RAFT cluster: %s", err.Error()))
		http.Error(w, "Error sending message to RAFT cluster", http.StatusInternalServerError)

		return
	}

	// Save mapping to Redis
	if err := fh.databaseAdapter.SaveMapping(handler.Filename, fileSize); err != nil {
		fh.loggerAdapter.Log(fmt.Sprintf("Error saving mapping to Redis: %s", err.Error()))
		http.Error(w, "Error saving mapping to Redis", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
}

// validateFile checks if the uploaded file is not empty and has a valid extension.
func validateFile(fileHeader *multipart.FileHeader) error {
	if fileHeader.Size <= 0 {
		return fmt.Errorf("file is empty")
	}

	ext := filepath.Ext(fileHeader.Filename)
	if ext == "" || !isValidExtension(ext) {
		return fmt.Errorf("invalid file extension")
	}

	return nil
}

// isValidExtension checks if the file extension is valid.
func isValidExtension(ext string) bool {
	return ext == ".txt"
}
