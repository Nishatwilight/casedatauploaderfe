package deleteunzip

import (
	"fmt"
	"os"
	"path/filepath"
)

func CleanupUnzippedFiles(directoryPath string) {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println("Error reading directory for cleanup:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() && file.Name() != "PPDS_OUTPUT_DATA" {
			// Delete the directory and its contents
			dirPath := filepath.Join(directoryPath, file.Name())
			err := os.RemoveAll(dirPath)
			if err != nil {
				fmt.Printf("Error deleting file %s: %v\n", file.Name(), err)
			} else {
				fmt.Printf("Deleted file %s\n", file.Name())
			}
		}
	}
}
