package main

import (
	"encoding/json"
	"fmt"

	"github.com/divanvisagie/ui"
)

type EditorView struct {
	box          *ui.Box
	keySelector  *ui.Combobox
	inputTextBox *ui.MultilineEntry
}

func translatePhrase(word string, targetLanguageStore *StringStore) (string, error) {
	translation, err := TranslateText(targetLanguageStore.value, word, apiKey)
	if err != nil {
		return "", err
	}
	return translation, nil
}

func translateJSONWithKey(targetLanguageStore *StringStore, jsonF *JSONFile, key string) string {
	parsed, _ := jsonF.Parse()
	for _, object := range parsed {
		for k, v := range object {
			if k != key {
				continue
			}
			translated, err := translatePhrase(v, targetLanguageStore)
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

func createMiddleSection(targetLanguageStore *StringStore, outputJSONControl *ui.MultilineEntry) (*ui.Box, *ui.Combobox) {
	box := ui.NewVerticalBox()

	button := ui.NewButton("Translate")
	button.OnClicked(func(b *ui.Button) {
		translatedJSONString := translateJSONWithKey(targetLanguageStore, jsonFileStore.file, targetJSONKeyStore.value)

		outputJSONControl.SetText(translatedJSONString)
		translatedJSONFileStore.SetValue(translatedJSONString)
	})

	combobox := ui.NewCombobox()
	box.Append(combobox, false)
	box.Append(button, false)
	combobox.OnSelected(func(c *ui.Combobox) {
		itemIndex := c.Selected()

		translationkey := jsonFileStore.file.Keys()[itemIndex]
		fmt.Println("Selected Item", translationkey)
		targetJSONKeyStore.SetValue(translationkey)

	})
	return box, combobox
}

//SetJSON sets the json file that the editor is working with
func (e *EditorView) SetJSON(jsonFile *JSONFile) {
	e.inputTextBox.SetText(jsonFile.ToString())

	parsed, _ := jsonFile.Parse()
	object := parsed[0]
	for k := range object {
		e.keySelector.Append(k)
	}
}

// CreateEditor creates a box that contains all the json editing related stuff
func CreateEditor(targetLanguageStore *StringStore) EditorView {

	box := ui.NewHorizontalBox()
	box.SetPadded(true)
	inputJSONControl := ui.NewMultilineEntry()
	outputJSONControl := ui.NewMultilineEntry()

	middleBox, combobox := createMiddleSection(targetLanguageStore, outputJSONControl)

	box.Append(inputJSONControl, true)
	box.Append(middleBox, true)
	box.Append(outputJSONControl, true)

	return EditorView{box, combobox, inputJSONControl}
}
