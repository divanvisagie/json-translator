package editor

import (
	"encoding/json"
	"fmt"
	"os"

	. "json-translator/internal/state"
	. "json-translator/internal/translation"
	. "json-translator/pkg/parser"
	. "json-translator/pkg/storage"

	"github.com/andlabs/ui"
)

type EditorView struct {
	Box          *ui.Box
	KeySelector  *ui.Combobox
	InputTextBox *ui.MultilineEntry
}

func translatePhrase(word string, targetLanguageStore *StringStore) (string, error) {
	var apiKey = os.Getenv("GOOGLE_API_KEY")
	translation, err := TranslateText(targetLanguageStore.Value, word, apiKey)
	if err != nil {
		return "", err
	}
	return translation, nil
}

func translateJSONWithKey(window *ui.Window, targetLanguageStore *StringStore, jsonF *JSONFile, key string) string {
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

func createMiddleSection(window *ui.Window, outputJSONControl *ui.MultilineEntry, state *ApplicationStateStore) (*ui.Box, *ui.Combobox) {
	box := ui.NewVerticalBox()

	button := ui.NewButton("Translate")
	button.OnClicked(func(b *ui.Button) {
		translatedJSONString := translateJSONWithKey(window, state.TargetLanguage, state.JSONFileStore.File, state.TargetJSONKey.Value)

		outputJSONControl.SetText(translatedJSONString)
		state.TranslatedJSONFile.SetValue(translatedJSONString)
	})

	combobox := ui.NewCombobox()
	box.Append(combobox, false)
	box.Append(button, false)
	combobox.OnSelected(func(c *ui.Combobox) {
		itemIndex := c.Selected()

		translationkey := state.JSONFileStore.File.Keys()[itemIndex]
		fmt.Println("Selected Item", translationkey)
		state.TargetJSONKey.SetValue(translationkey)

	})
	return box, combobox
}

// SetJSON sets the json file that the editor is working with
func (e *EditorView) SetJSON(jsonFile *JSONFile) {
	e.InputTextBox.SetText(jsonFile.ToString())

	parsed, _ := jsonFile.Parse()
	object := parsed[0]
	for k := range object {
		e.KeySelector.Append(k)
	}
}

// CreateEditor creates a box that contains all the json editing related stuff
func CreateEditor(window *ui.Window, state *ApplicationStateStore) EditorView {

	box := ui.NewHorizontalBox()
	box.SetPadded(true)
	inputJSONControl := ui.NewMultilineEntry()
	outputJSONControl := ui.NewMultilineEntry()

	middleBox, combobox := createMiddleSection(window, outputJSONControl, state)

	box.Append(inputJSONControl, true)
	box.Append(middleBox, true)
	box.Append(outputJSONControl, true)

	return EditorView{box, combobox, inputJSONControl}
}