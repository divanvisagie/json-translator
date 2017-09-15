package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type JSONFile struct {
	raw  []byte
	path string
}

func (j *JSONFile) ToString() string {
	return string(j.raw)
}

func ReadJsonFromFile(file string) JSONFile {
	fmt.Println("reading", file)
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(raw)

	translationArray := []map[string]string{}
	unmarErr := json.Unmarshal(raw, &translationArray)

	if unmarErr != nil {
		log.Fatalln(unmarErr.Error())
	}

	for _, object := range translationArray {
		for k, v := range object {
			fmt.Println(k, v)
		}
	}

	return JSONFile{
		raw,
		file,
	}
}
