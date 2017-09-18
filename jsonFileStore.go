package main

type JSONFileStore struct {
	file    *JSONFile
	channel chan *JSONFile
}

func (j *JSONFileStore) SetJsonFile(jsonFile *JSONFile) {
	j.file = jsonFile
	j.channel <- jsonFile
}

func (j *JSONFileStore) Destroy() {
	close(j.channel)
}

// CreateJSONFileStore Creates a new instance
func CreateJSONFileStore() *JSONFileStore {
	jsonChannel := make(chan *JSONFile)
	return &JSONFileStore{
		nil,
		jsonChannel,
	}
}
