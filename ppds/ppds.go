package ppds

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// MoveFilesToPPDSOutputData moves files to the "PPDS_OUTPUT_DATA" folder.
func MoveFilesToPPDSOutputData(sourcePath, outputPath string) error {
	// Create the "PPDS_OUTPUT_DATA" folder
	if err := os.Mkdir(outputPath, 0755); err != nil {
		fmt.Printf("Error creating 'PPDS_OUTPUT_DATA' folder: %v\n", err)
		return err
	}

	// Get a list of files and folders in the sourcePath
	contents, err := ioutil.ReadDir(sourcePath)
	if err != nil {
		fmt.Printf("Error reading source path: %v\n", err)
		return err
	}

	// Iterate over the files and folders
	for _, item := range contents {
		itemPath := filepath.Join(sourcePath, item.Name())

		// Check if the item is a file (not a directory)
		if !item.IsDir() {
			// Move the file to the "PPDS_OUTPUT_DATA" folder
			newPath := filepath.Join(outputPath, item.Name())
			if err := os.Rename(itemPath, newPath); err != nil {
				fmt.Printf("Error moving %s to 'PPDS_OUTPUT_DATA' folder: %v\n", item.Name(), err)
			} else {
				fmt.Printf("Moved %s to 'PPDS_OUTPUT_DATA' folder.\n", item.Name())
			}
		}
	}

	return nil
}
