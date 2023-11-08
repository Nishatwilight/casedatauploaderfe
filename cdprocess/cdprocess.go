package cdprocess

import (
	deleteunzip "case-data-upload-to-aws-s3-bucket/deleteunzip"
	"case-data-upload-to-aws-s3-bucket/ppds"
	"case-data-upload-to-aws-s3-bucket/unzip"
	"case-data-upload-to-aws-s3-bucket/zipped"

	"fmt"
	"path/filepath"
)

func Casedataprocessing() {

	// Specify the path to your zip file
	zipFilePath := "C:\\Users\\TLSPC-106\\Downloads\\caseData\\2003.zip"

	// Open the zip file for reading
	zipFile, err := unzip.OpenZipFile(zipFilePath)
	if err != nil {
		fmt.Println("Error opening zip file:", err)
		return
	}
	defer zipFile.Close()

	// Create a directory to extract the files
	extractPath := filepath.Dir(zipFilePath)

	// Extract the files using the unzip package
	if err := unzip.ExtractFiles(zipFilePath, extractPath); err != nil {
		fmt.Println("Error extracting files:", err)
		return
	} else {
		fmt.Println("Case Data unzipped successfully")
	}

	// Example usage of functions from the zip package
	folderPath := "C:\\Users\\TLSPC-106\\Downloads\\caseData\\R2_ORG_newSerializerModifiedVisualProperties\\R2_ORG"
	outputPath := folderPath

	folders := []string{"01_Inputs", "02_Segmentation", "03_Plans", "04_Designs"}

	for _, folder := range folders {
		zipFolderPath := filepath.Join(outputPath, folder+".zip")
		if err := zipped.ZipFolder(filepath.Join(folderPath, folder), zipFolderPath); err != nil {
			fmt.Printf("Error zipping folder %s: %v\n", folder, err)
		} else {
			fmt.Printf("Folder %s zipped successfully.\n", folder)
		}
	}

	// Example usage of functions from the ppds package
	ppds.MoveFilesToPPDSOutputData(folderPath, filepath.Join(folderPath, "PPDS_OUTPUT_DATA"))
	deleteunzip.CleanupUnzippedFiles(outputPath)

}
