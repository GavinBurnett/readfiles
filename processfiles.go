package main

import (
	"fmt"
)

var fileReadOKCount int64
var skippedFilesCount int64
var skippedDirectoriesCount int64
var fileErrorCount int64
var fileSizeMismatchCount int64

// ProcessFiles: Read all files in the given directory, and all sub directory
func ProcessFiles(_startDir string, _listFiles bool) {

	var bytesRead int64
	bytesRead = -1

	if DEBUG == true {
		fmt.Println(fmt.Sprintf(UI_Arguments, GetFunctionName()), "_startDir: ", _startDir, "_listfiles: ", _listFiles)
	}

	fileReadOKCount = 0
	skippedFilesCount = 0
	skippedDirectoriesCount = 0
	fileErrorCount = 0
	fileSizeMismatchCount = 0

	var loopCount int64
	loopCount = 0

	if len(_startDir) > 0 {
		allFiles = GetAllFiles(_startDir)
		if allFiles != nil && len(allFiles) > 0 {
			for _, file := range allFiles {

				loopCount++
				fileSize := GetFileSize(file)
				if fileSize != -1 {
					bytesRead = ReadFile(file)
					if bytesRead != -1 {
						if bytesRead == fileSize {
							if _listFiles == true {
								fmt.Println(file, UI_OK)
							} else {
								fmt.Printf("\rFile Count: (%v/%v)", CommaFormat(loopCount), CommaFormat((int64(len(allFiles)))))
							}
							fileReadOKCount++
						} else {
							if DEBUG == true || _listFiles == true {
								fmt.Println(UI_FileSizeMismatch, file)
							} else {
								fmt.Printf("\rFile Count: (%v/%v)", CommaFormat(loopCount), CommaFormat((int64(len(allFiles)))))
							}
							fileSizeMismatchCount++
						}
					} else {
						if DEBUG == true || _listFiles == true {
							fmt.Println(UI_ReadFileError, file)
						} else {
							fmt.Printf("\rFile Count: (%v/%v)", CommaFormat(loopCount), CommaFormat((int64(len(allFiles)))))
						}
						fileErrorCount++
					}
				} else {
					if DEBUG == true || _listFiles == true {
						fmt.Println(UI_SkippingFile, file)
					} else {
						fmt.Printf("\rFile Count: (%v/%v)", CommaFormat(loopCount), CommaFormat((int64(len(allFiles)))))
					}
					skippedFilesCount++
				}
			} // end for

			fmt.Println("\n")
			fmt.Println(UI_FilesRead, CommaFormat(fileReadOKCount))
			fmt.Println(UI_FilesSkipped, CommaFormat(skippedFilesCount))
			fmt.Println(UI_DirectoriesSkipped, CommaFormat(skippedDirectoriesCount))
			fmt.Println(UI_FileErrors, CommaFormat(fileErrorCount))
			fmt.Println(UI_FileSizeMismatches, CommaFormat(fileSizeMismatchCount))
		} else {
			fmt.Println(UI_NoFilesToProcess)
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_startDir: ", _startDir, "_listfiles: ", _listFiles)
	}
}
