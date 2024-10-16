package ccc

import (
	"fmt"
	"path"
	"testing"

	"json2plist/pkg/egret"
)

func TestSaveXmlData(t *testing.T) {
	jsonData := egret.LoadJsonData("font_clock.json")

	filename := "font_clock.plist"

	if !SaveXmlData(filename, jsonData) {
		t.Errorf("Error saving XML data")
	}
}

func TestFile(t *testing.T) {
	fmt.Println(path.Base("ddd.json"))

}
