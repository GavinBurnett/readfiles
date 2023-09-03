// +build linux

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// GetAllFiles: Get all files in the given directory and all sub directories
func GetAllFiles(_startDir string) []string {

	if DEBUG == true {
		fmt.Println(fmt.Sprintf(UI_Arguments, GetFunctionName()), "_startDir: ", _startDir)
	}

	if len(_startDir) > 0 {
		files, err := ioutil.ReadDir(_startDir)
		if err != nil {
			if DEBUG == true {
				fmt.Println(UI_DirectoryReadFail, _startDir)
			}
		} else {
			for _, file := range files {
				fmt.Printf("\rBuilding Files.")
				if file.IsDir() {

					if skipDirs != nil && len(skipDirs) > 0 {

						if SkipDir(file.Name()) == true {
							if DEBUG == true {
								fmt.Println(UI_SkippingDirectory, file.Name())
							}
							skippedDirectoriesCount++
						} else {
							subDir := filepath.Join(_startDir, file.Name())
							fmt.Printf("\rBuilding Files..")
							GetAllFiles(subDir)
						}
					} else {
						subDir := filepath.Join(_startDir, file.Name())
						fmt.Printf("\rBuilding Files..")
						GetAllFiles(subDir)
					}
				} else {
					fmt.Printf("\rBuilding Files...")
					allFiles = append(allFiles, _startDir+"/"+file.Name())
				}
			} // end for
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_startDir: ", _startDir)
	}

	return allFiles
}
