package jsonstore

import (
	. "json-translator/pkg/parser"
)

type JSONFileStore struct {
	File    *JSONFile
	Channel chan *JSONFile
}

func (j *JSONFileStore) SetJsonFile(jsonFile *JSONFile) {
	j.File = jsonFile
	j.Channel <- jsonFile
}

func (j *JSONFileStore) Destroy() {
	close(j.Channel)
}

// CreateJSONFileStore Creates a new instance
func CreateJSONFileStore() *JSONFileStore {
	jsonChannel := make(chan *JSONFile)
	return &JSONFileStore{
		nil,
		jsonChannel,
	}
}
