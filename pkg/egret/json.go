package egret

import (
	"encoding/json"
	"io"
	"os"
)

type JsonFrame struct {
	X int `json:"x"`
	Y int `json:"y"`

	W int `json:"w"`
	H int `json:"h"`

	OffX int `json:"offX"`
	OffY int `json:"offY"`

	SourceW int `json:"sourceW"`
	SourceH int `json:"sourceH"`
}

type JsonData struct {
	File   string               `json:"file"`
	Frames map[string]JsonFrame `json:"frames"`
}

func LoadJsonData(filePath string) *JsonData {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var jsonData JsonData
	_ = json.Unmarshal(byteValue, &jsonData)

	return &jsonData
}
