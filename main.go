// readfiles project main.go
package main

import (
	"fmt"
	"os"
	"strings"
)

// Main: Main entry point
func main() {

	exitCode := 0

	if os.Args != nil {

		args := os.Args
		validArgs := false

		//percent := 0
		//files := make([]string, len(args))

		// args[0] = readfiles
		// args[1] = start directory
		// args[1,2] = start directory --listfiles
		// args[1,2] = start directory --skipdirectories=dir1,dir2,dir3
		// args[1,2,3] = start directory --listfiles --skipdirectories=dir1,dir2,dir3

		if len(args) == 1 {
			// no user arguments given - display help
			fmt.Println(UI_Help)
			validArgs = true
		}
		if len(args) == 2 {
			if IsStringHelpArgument(args[1]) {
				// user has given help argument - display help
				fmt.Println(UI_Help)
				validArgs = true
			} else {
				// user has given only one argument that is not a help argument - argument must be directory name
				if DirExists(args[1]) {
					// process the files
					ProcessFiles(args[1], false)
					validArgs = true
				} else {
					fmt.Println(UI_DirectoryNotFound)
				}
			}
		}
		if len(args) == 3 {
			// user has given two arguments - must be directory name and list files/skip directories arguments
			if DirExists(args[1]) {
				listFilesArg := strings.Compare(args[2], LIST_FILES_ARG)
				if listFilesArg == 0 {
					// process the files
					ProcessFiles(args[1], true)
					validArgs = true
				} else {
					// list files argument not given
				}

				skipFilesArg := strings.Contains(args[2], SKIP_DIRS_ARG+"=")
				if skipFilesArg == true {
					SetSkipDirs(args[2])
					// process the files
					ProcessFiles(args[1], false)
					validArgs = true
				} else {
					// skip argument not given
				}
			} else {
				fmt.Println(UI_DirectoryNotFound)
			}
		}
		if len(args) == 4 {
			// user has given three arguments - must be directory name and list files/skip directories arguments
			if DirExists(args[1]) {

				skipFilesArg := strings.Contains(args[3], SKIP_DIRS_ARG+"=")
				if skipFilesArg == true {
					SetSkipDirs(args[3])

					listFilesArg := strings.Compare(args[2], LIST_FILES_ARG)
					if listFilesArg == 0 {
						// list files argument given
						// process the files
						ProcessFiles(args[1], true)
						validArgs = true
					} else {
						// list files argument not given
						// process the files
						ProcessFiles(args[1], false)
						validArgs = true
					}

				} else {
					// skip argument not given

					listFilesArg := strings.Compare(args[2], LIST_FILES_ARG)
					if listFilesArg == 0 {
						// list files argument given
						// process the files
						ProcessFiles(args[1], true)
						validArgs = true
					} else {
						// list files argument not given
						// process the files
						ProcessFiles(args[1], false)
						validArgs = true
					}
				}

			} else {
				fmt.Println(UI_DirectoryNotFound)
			}
		}

		if validArgs == false {
			exitCode = -1
			fmt.Println(UI_InvalidArgs)
		}

	} else {
		exitCode = -1
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
	}

	os.Exit(exitCode)
}
