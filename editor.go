package main

import (
	"encoding/json"
	"fmt"

	"github.com/divanvisagie/ui"
)

func translatePhrase(word string) (string, error) {
	translation, err := TranslateText(targetLanguageStore.value, word, apiKey)
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

func createMiddleSection(outputJSONControl *ui.MultilineEntry) (*ui.Box, *ui.Combobox) {
	box := ui.NewVerticalBox()

	combobox := ui.NewCombobox()
	box.Append(combobox, true)
	combobox.OnSelected(func(c *ui.Combobox) {
		itemIndex := c.Selected()

		translationkey := jsonFileStore.file.Keys()[itemIndex]

		fmt.Println("Selected Item", translationkey)
		translatedJSONString := translateJSONWithKey(jsonFileStore.file, translationkey)

		outputJSONControl.SetText(translatedJSONString)
	})
	return box, combobox
}

// CreateEditor creates a box that contains all the json editing related stuff
func CreateEditor() *ui.Box {

	box := ui.NewHorizontalBox()
	box.SetPadded(true)
	inputJSONControl := ui.NewMultilineEntry()
	outputJSONControl := ui.NewMultilineEntry()

	middleBox, combobox := createMiddleSection(outputJSONControl)

	box.Append(inputJSONControl, true)
	box.Append(middleBox, true)
	box.Append(outputJSONControl, true)

	go func() {
		for jsonFile := range jsonFileStore.channel {
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
