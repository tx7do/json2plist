package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"json2plist/pkg/ccc"
	"json2plist/pkg/egret"
)

func handleSingleFile(filePath string) bool {
	jsonData := egret.LoadJsonData(filePath)
	if jsonData == nil {
		log.Printf("load egret json data failed: %s\n", filePath)
		return false
	}

	outputFilePath := strings.Replace(filePath, ".json", ".plist", -1)
	outputFilePath = strings.Replace(outputFilePath, ".fnt", ".plist", -1)

	if !ccc.SaveXmlData(outputFilePath, jsonData) {
		log.Printf("save ccc xml data failed: %s\n", outputFilePath)
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

		if filepath.Ext(abs) != ".json" && filepath.Ext(abs) != ".fnt" {
			continue
		}

		currentFiles = append(currentFiles, filepath.Join(path, info.Name()))
	}

	return currentFiles, nil
}

func handleFolder(dir string) bool {
	files, _ := getCurrentDirFiles(dir)
	//fmt.Println("files: ", files)

	log.Println("start to handle folder: ", dir)
	log.Println("total files: ", len(files))

	var count int
	for _, file := range files {
		if handleSingleFile(file) {
			count++
		}
	}

	log.Println("total success files: ", count)

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
