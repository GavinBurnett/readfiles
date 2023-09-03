// +build linux

package main

const (
	UI_Help = `readfiles Version 1.2 - by gburnett@outlook.com

Reads every file in a given directory and all its sub directories.
Warns if any file cannot be read from start to end.

Usage:	
readfiles start directory
readfiles start directory --listfiles
readfiles start directory --listfiles --skipdirectories=dir1,dir2,dir3

Examples:
readfiles /home
readfiles /home --listfiles
readfiles / --listfiles --skipdirectories=bin,proc,tmp`
)
