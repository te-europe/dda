package file_service

import (
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func DeleteFromTemporary(path string) error {
	// Delete the temporary file

	err := os.Remove(path)

	if err != nil {
		return err
	}

	return nil
}

func DownloadFileFromRepo(url string) (string, error) {
	// Download the file from the repo
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// Write the file to a temporary file
	outputPath := os.TempDir() + uuid.New().String() + ".bicep"

	out, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return outputPath, err
}
