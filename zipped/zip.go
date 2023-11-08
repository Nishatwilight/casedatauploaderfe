package zipped

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ZipFolder compresses a folder into a zip file.
func ZipFolder(sourceFolder, zipFile string) error {
	// Open the zip file for writing
	zipf, err := os.Create(zipFile)
	if err != nil {
		return err
	}
	defer zipf.Close()

	// Create a new zip archive
	zipw := zip.NewWriter(zipf)
	defer zipw.Close()

	// Walk through the source folder and add its contents to the zip archive
	err = filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip adding directory entries to the zip archive
		if info.IsDir() {
			return nil
		}

		// Create a header for the file in the zip archive
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set the name of the file within the zip archive
		header.Name, _ = filepath.Rel(sourceFolder, path)

		// Use deflate compression
		header.Method = zip.Deflate

		// Create a writer for the file in the zip archive
		writer, err := zipw.CreateHeader(header)
		if err != nil {
			return err
		}

		// Write the file's contents to the zip archive
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
