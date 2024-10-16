package main

import (
	"flag"
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

func handleFolder(dir string) bool {

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
