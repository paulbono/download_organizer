package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	homeDir, err := os.UserHomeDir()
	printErr(err)
	downloadPath := path.Join(homeDir, "/Downloads")
	entries, err := os.ReadDir(downloadPath)
	printErr(err)
	var filetypeFolder string
	var newFileName string
	var filePath string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		extension := filepath.Ext(entry.Name())
		if extension == "" {
			continue
		}
		filetypeFolder = path.Join(downloadPath, "filetype_"+extension[1:])
		folderStat, err := os.Stat(filetypeFolder)
		printErr(err)
		if folderStat == nil {
			continue
		}
		filePath = path.Join(downloadPath, "/", entry.Name())
		fileStat, err := os.Stat(filePath)
		printErr(err)
		if fileStat == nil {
			continue
		}
		if time.Now().Sub(fileStat.ModTime()) > 1*time.Hour {
			newFileName = path.Join(filetypeFolder, "/", entry.Name())
			os.Rename(filePath, newFileName)
		}
	}
}
