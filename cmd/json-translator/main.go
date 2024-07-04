package main

import (
	"io/ioutil"
	"log"
	"os"

	. "json-translator/internal/destinationbox"
	. "json-translator/internal/editor"
	. "json-translator/internal/jsonstore"
	. "json-translator/internal/language"
	. "json-translator/internal/sourcebox"
	. "json-translator/internal/state"
	. "json-translator/pkg/storage"

	"github.com/andlabs/ui"
)

func main() {
	var window *ui.Window

	var state = ApplicationStateStore{
		DestinationFilePath: CreateStringStore(),
		SourceFilePath:      CreateStringStore(),
		TargetJSONKey:       CreateStringStore(),
		TargetLanguage:      CreateStringStore(),
		TranslatedJSONFile:  CreateStringStore(),
		JSONFileStore:       CreateJSONFileStore(),
	}

	err := ui.Main(func() {

		var apiKey = os.Getenv("GOOGLE_API_KEY")
		window = ui.NewWindow("JSON Translator", 500, 500, false)
		box := ui.NewVerticalBox()

		saveButton := ui.NewButton("Save")

		saveButton.OnClicked(func(b *ui.Button) {
			if state.DestinationFilePath.Value == "" {
				ui.MsgBoxError(window, "Target path error", "Target path is not defined")
				return
			}

			data := []byte(state.TranslatedJSONFile.Value)
			ioutil.WriteFile(state.DestinationFilePath.Value, data, 0644)
		})

		editor := CreateEditor(window, &state)
		go func() {
			for jsonFile := range state.JSONFileStore.Channel {
				editor.SetJSON(jsonFile)
			}
		}()

		box.Append(ui.NewLabel("Select Source File:"), false)
		box.Append(CreateSourceInputBox(window, &state), false)
		box.Append(ui.NewLabel("Select Destination File:"), false)
		box.Append(CreateDestinationInputBox(window, state.DestinationFilePath), false)
		box.Append(CreateLanguageSelector(window, apiKey, &state), false)
		box.Append(editor.Box, true)
		box.Append(saveButton, false)

		window.SetChild(box)
		window.SetMargined(true)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			// jsonFileStore.Destroy()
			// targetLanguageStore.Destroy()
			// targetJSONKeyStore.Destroy()
			// sourceFilePathStore.Destroy()
			// destinationFilePathStore.Destroy()
			// translatedJSONFileStore.Destroy()
			return true
		})
		window.Show()

	})
	if err != nil {
		log.Fatalln(err)
	}
}
