package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var allowExits = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

var allowMimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

const maxSize = 5 << 20

func ValidateAndSaveFile(fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	// Check extension in filename
	etx := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowExits[etx] {
		return "", errors.New("unsupported file extension")
	}

	// Check size
	if fileHeader.Size > maxSize {
		return "", errors.New("file too large")
	}

	// Check file type
	file, err := fileHeader.Open()
	if err != nil {
		return "", errors.New("error opening file")
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file")
		}
	}(file)

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", errors.New("error reading file")
	}

	mimeType := http.DetectContentType(buffer)
	if !allowMimeTypes[mimeType] {
		return "", fmt.Errorf("invalid mime type: %s", mimeType)
	}

	// Change filename
	filename := fmt.Sprintf("%s%s", uuid.New().String(), etx)

	// Create folder if not exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		return "", errors.New("error creating directory")
	}

	savePath := filepath.Join(uploadDir, filename)
	if err := saveFile(fileHeader, savePath); err != nil {
		return "", errors.New("error saving file")
	}

	return filename, nil
}

func saveFile(fileHeader *multipart.FileHeader, destination string) error {
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			fmt.Println("error closing file")
		}
	}(src)
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Println("error closing file")
		}
	}(out)

	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}

	return err
}
