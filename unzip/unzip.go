package unzip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func OpenZipFile(zipFile string) (*zip.ReadCloser, error) {
	zipFileReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return nil, err
	}
	return zipFileReader, nil
}

// ExtractFiles extracts files from a zip archive to a destination folder.
func ExtractFiles(zipFilePath, destinationFolder string) error {
	// Open the zip file for reading
	zipFile, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create the destination folder if it doesn't exist
	if err := os.MkdirAll(destinationFolder, os.ModePerm); err != nil {
		return err
	}

	// Extract each file in the zip archive while preserving directory structure
	for _, file := range zipFile.File {
		localFilePath := filepath.Join(destinationFolder, file.Name)
		if file.FileInfo().IsDir() {
			// If the item is a directory, create it locally
			if err := os.MkdirAll(localFilePath, os.ModePerm); err != nil {
				return err
			}
		} else {
			localFile, err := os.Create(localFilePath)
			if err != nil {
				return err
			}
			defer localFile.Close()

			zippedFile, err := file.Open()
			if err != nil {
				return err
			}
			defer zippedFile.Close()

			_, err = io.Copy(localFile, zippedFile)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
