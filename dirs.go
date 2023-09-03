package main

import (
	"fmt"
	"os"
	"strings"
)

var skipDirs []string

// DirExists: Returns true if the given directory exists, false if it does not
func DirExists(_dirPath string) bool {

	dirExists := false

	if DEBUG == true {
		fmt.Println(fmt.Sprintf(UI_Arguments, GetFunctionName()), "_dirPath: ", _dirPath)
	}

	if len(_dirPath) > 0 {

		_, err := os.Stat(_dirPath)
		if err == nil {
			dirExists = true
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_dirPath: ", _dirPath)
	}

	return dirExists
}

// SetSkipDirs: Extract and set list of directories to skip
func SetSkipDirs(_dirs string) {

	if DEBUG == true {
		fmt.Println(fmt.Sprintf(UI_Arguments, GetFunctionName()), "_dirs: ", _dirs)
	}

	if len(_dirs) > 0 {
		splitDirs := strings.Split(_dirs, ",")
		skipDirs = make([]string, len(splitDirs))
		if splitDirs != nil && len(splitDirs) > 0 {
			splitDirs2 := strings.Split(splitDirs[0], "=")
			if splitDirs2 != nil && len(splitDirs2) > 0 {
				skipDirs[0] = splitDirs2[1]

				for counter := 1; counter != len(splitDirs); counter++ {
					skipDirs[counter] = splitDirs[counter]
				}

			}
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_dirs: ", _dirs)
	}
}
