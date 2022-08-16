package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type SaveJsFile struct {
	filePath string
}

func NewSaveJsFile(filepath string) *SaveJsFile {
	return &SaveJsFile{
		filePath: filepath,
	}
}

func (s *SaveJsFile) Save(stations []*TrainStation) {

	jsBuff, _ := json.Marshal(stations)

	buffer := bytes.NewBufferString("var station_names =")
	buffer.Write(jsBuff)
	buffer.WriteString(";")

	ioutil.WriteFile(s.filePath, buffer.Bytes(), 0644)
}
