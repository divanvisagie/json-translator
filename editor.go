package main

import (
	"encoding/json"
	"fmt"

	"github.com/divanvisagie/ui"
)

func translatePhrase(word string) (string, error) {
	translation, err := TranslateText("fr", word, apiKey)
	if err != nil {
		return "", err

	}
	return translation, nil
}

func translateJSONWithKey(jsonF *JSONFile, key string) string {
	parsed, _ := jsonF.Parse()
	for _, object := range parsed {
		for k, v := range object {
			if k != key {
				continue
			}
			translated, err := translatePhrase(v)
			if err != nil {
				ui.MsgBoxError(window, "Translation Error", err.Error())
				return ""
			}
			fmt.Println("translated:", translated)
			object[k] = translated
		}
	}
	b, _ := json.Marshal(parsed)

	fmt.Println(string(b))
	return string(b)
}

// CreateEditor creates a box that contains all the json editing related stuff
func CreateEditor() *ui.Box {

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
		translatedJSONString := translateJSONWithKey(currentJSONFile, translationkey)

		outputJSONControl.SetText(translatedJSONString)
	})

	go func() {
		for jsonFile := range jsonFileStore.channel {
			currentJSONFile = jsonFile
			inputJSONControl.SetText(jsonFile.ToString())
			parsed, _ := jsonFile.Parse()
			object := parsed[0]
			for k := range object {
				combobox.Append(k)
			}
		}
	}()

	return box
}
