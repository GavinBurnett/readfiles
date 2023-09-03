package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var allFiles []string

// GetFileSize: Gets the size of the given file in bytes
func GetFileSize(_file string) int64 {

	var fileSize int64
	fileSize = -1

	if DEBUG == true {
		fmt.Println(fmt.Sprintf(UI_Arguments, GetFunctionName()), "_file: ", _file)
	}

	if len(_file) > 0 {

		fileInfo, err := os.Stat(_file)
		if err == nil {
			fileSize = fileInfo.Size()

			if fileSize == 0 {
				if DEBUG == true {
					fmt.Println(UI_EmptyFile)
				}
				fileSize = -1
			}

			if fileSize < 0 {
				if DEBUG == true {
					fmt.Println(UI_InvalidFileSize)
				}
				fileSize = -1
			}

			if fileSize > math.MaxInt64 {
				if DEBUG == true {
					fmt.Println(UI_FileTooBig)
				}
				fileSize = -1
			}

		} else {
			if DEBUG == true {
				fmt.Println(UI_NoFileSize, _file, err.Error())
			}
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
	}

	return fileSize
}

// ReadFile: Reads given file from start to end
func ReadFile(_file string) int64 {

	const bufferSize = 1000

	var totalBytesRead int
	totalBytesRead = -1

	if DEBUG == true {
		fmt.Println(fmt.Sprintf(UI_Arguments, GetFunctionName()), "_file: ", _file)
	}

	if len(_file) > 0 {

		fileHandle, err := os.OpenFile(_file, os.O_RDONLY, 0)

		if err == nil {

			buffer := make([]byte, bufferSize)

			for {
				bytesRead, err := fileHandle.Read(buffer)
				totalBytesRead += bytesRead
				if err != nil {

					if err == io.EOF {
						if DEBUG == true {
							fmt.Println(_file, ":", UI_EOF)
						}
						break
					} else {
						if DEBUG == true {
							fmt.Println(_file, ":", err)
						}
						totalBytesRead = -1
						break
					}
				} else {
					if DEBUG == true {
						fmt.Println(_file, ":", totalBytesRead)
					}
				}
			}
			totalBytesRead += 1
		} else {
			if DEBUG == true {
				fmt.Println(UI_OpenFileError, _file)
			}
			totalBytesRead = -1
		}

		fileHandle.Close()

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file: ", _file)
	}

	return int64(totalBytesRead)
}

// SkipDir: Returns true if given directory is to be skipped, false otherwise
func SkipDir(_dir string) bool {

	var skipDir bool
	skipDir = false

	if DEBUG == true {
		fmt.Println(fmt.Sprintf(UI_Arguments, GetFunctionName()), "_dir: ", _dir, "skipDirs: ", skipDirs)
	}

	if len(_dir) > 0 && skipDirs != nil && len(skipDirs) > 0 {

		for _, currentDir := range skipDirs {
			if currentDir == _dir {
				skipDir = true
				break
			}
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_dir:", _dir, "skipDirs: ", skipDirs)
	}

	return skipDir
}
