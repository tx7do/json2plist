package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"json2plist/pkg/ccc"
	"json2plist/pkg/egret"
)

func handleSingleFile(filePath string) bool {
	jsonData := egret.LoadJsonData(filePath)
	if jsonData == nil {
		return false
	}

	outputFilePath := strings.Replace(filePath, ".json", ".plist", -1)

	if !ccc.SaveXmlData(outputFilePath, jsonData) {
		return false
	}

	return true
}

func getCurrentDirFiles(path string) ([]string, error) {
	infos, err := os.ReadDir(path)
	if err != nil {
		return []string{}, nil
	}
	var currentFiles []string
	for _, info := range infos {
		abs, err := filepath.Abs(info.Name())
		if err != nil {
			return []string{}, nil
		}

		if info.IsDir() {
			continue
		}

		if filepath.Ext(abs) != ".json" {
			continue
		}

		currentFiles = append(currentFiles, abs)
	}

	return currentFiles, nil
}

func handleFolder(dir string) bool {
	files, _ := getCurrentDirFiles(dir)
	fmt.Println("files: ", files)

	for _, file := range files {
		handleSingleFile(file)
	}

	return true
}

func main() {
	includeFile := flag.String("f", "", "输入文件")
	includeDir := flag.String("d", "", "输入文件夹")
	flag.Parse()
	//args := flag.Args()

	if *includeFile != "" {
		handleSingleFile(*includeFile)
	} else if *includeDir != "" {
		handleFolder(*includeDir)
	}
}
