package ccc

import (
	"fmt"
	"os"

	"json2plist/pkg/egret"
	"json2plist/pkg/xmlfmt"
)

const (
	XmlHeader = `<?xml version="1.0" encoding="UTF-8"?>` +
		"\n" +
		`<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">` +
		"\n"
)

func writeFrames(strXml *string, data *egret.JsonData) {
	writeKeyNode(strXml, "frames")

	writeDictNodeBegin(strXml)
	{
		for key, value := range data.Frames {
			writeFrameNode(strXml, key, &value)
		}
	}
	writeDictNodeEnd(strXml)
}

func writeFrameNode(strXml *string, frameName string, frame *egret.JsonFrame) {
	writeKeyNode(strXml, frameName)

	writeDictNodeBegin(strXml)
	{
		writeKeyNode(strXml, "frame")
		writeStringNode(strXml, fmt.Sprintf("{{%d,%d},{%d,%d}}", frame.X, frame.Y, frame.W, frame.H))

		writeKeyNode(strXml, "offset")
		writeStringNode(strXml, fmt.Sprintf("{%d,%d}", frame.OffX, frame.OffY))

		writeKeyNode(strXml, "rotated")
		writeFalseNode(strXml)

		writeKeyNode(strXml, "sourceColorRect")
		writeStringNode(strXml, fmt.Sprintf("{{0,0},{%d,%d}}", frame.SourceW, frame.SourceH))

		writeKeyNode(strXml, "sourceSize")
		writeStringNode(strXml, fmt.Sprintf("{%d,%d}", frame.SourceW, frame.SourceH))
	}
	writeDictNodeEnd(strXml)
}

func writeMetadata(strXml *string, data *egret.JsonData) {
	writeKeyNode(strXml, "metadata")

	writeDictNodeBegin(strXml)
	{
		writeKeyNode(strXml, "format")
		writeIntegerNode(strXml, 2)

		writeKeyNode(strXml, "realTextureFileName")
		writeStringNode(strXml, data.File)

		//writeKeyNode(strXml, "size")
		//writeStringNode(strXml, "{512,512}")

		writeKeyNode(strXml, "textureFileName")
		writeStringNode(strXml, data.File)
	}
	writeDictNodeEnd(strXml)
}

func writePlistNodeBegin(strXml *string) {
	*strXml += "<plist version=\"1.0\">\n"
}

func writePlistNodeEnd(strXml *string) {
	*strXml += "</plist>"
}

func writeDictNodeBegin(strXml *string) {
	*strXml += "<dict>\n"
}

func writeDictNodeEnd(strXml *string) {
	*strXml += "</dict>\n"
}

func writeKeyNode(strXml *string, value string) {
	*strXml += fmt.Sprintf("<key>%s</key>\n", value)
}

func writeStringNode(strXml *string, value string) {
	*strXml += fmt.Sprintf("<string>%s</string>\n", value)
}

func writeIntegerNode(strXml *string, value int) {
	*strXml += fmt.Sprintf("<integer>%d</integer>\n", value)
}

func writeTrueNode(strXml *string) {
	*strXml += "<true/>\n"
}

func writeFalseNode(strXml *string) {
	*strXml += "<false/>\n"
}

func SaveXmlData(filename string, data *egret.JsonData) bool {
	xmlFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return false
	}
	defer xmlFile.Close()

	var strXml string

	strXml = XmlHeader

	writePlistNodeBegin(&strXml)

	{
		writeDictNodeBegin(&strXml)
		{
			writeFrames(&strXml, data)
			writeMetadata(&strXml, data)
		}
		writeDictNodeEnd(&strXml)
	}

	writePlistNodeEnd(&strXml)

	strXml = xmlfmt.FormatXML(strXml, "", "  ", true)
	_, _ = xmlFile.WriteString(strXml)

	return true
}
