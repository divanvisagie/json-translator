package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/divanvisagie/ui"
)

func translatePhrase(word string) string {
	translation, err := TranslateText("fr", word, apiKey)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return translation
}

func translateJsonWithKey(jsonF *JSONFile, key string) string {
	parsed, _ := jsonF.Parse()
	for _, object := range parsed {
		for k, v := range object {
			if k != key {
				continue
			}
			translated := translatePhrase(v)
			fmt.Println("translated:", translated)
			object[k] = translated
		}
	}
	b, _ := json.Marshal(parsed)

	fmt.Println(string(b))
	return string(b)
}

func CreateEditor(ch chan *JSONFile) *ui.Box {

	var currentJSONFile *JSONFile

	box := ui.NewHorizontalBox()
	box.SetPadded(true)
	inputJSONControl := ui.NewMultilineEntry()
	outputJSONControl := ui.NewMultilineEntry()

	combobox := ui.NewCombobox()

	box.Append(inputJSONControl, true)
	box.Append(combobox, true)
	box.Append(outputJSONControl, true)

	combobox.OnSelected(func(c *ui.Combobox) {
		itemIndex := c.Selected()

		translationkey := currentJSONFile.Keys()[itemIndex]

		fmt.Println("Selected Item", translationkey)
		translatedJsonString := translateJsonWithKey(currentJSONFile, translationkey)

		outputJSONControl.SetText(translatedJsonString)
	})

	go func() {
		for jsonFile := range ch {
			currentJSONFile = jsonFile
			inputJSONControl.SetText(jsonFile.ToString())
			parsed, _ := jsonFile.Parse()
			object := parsed[0]
			for k, _ := range object {
				// fmt.Println(k, v)
				combobox.Append(k)
			}
		}
	}()

	return box
}
