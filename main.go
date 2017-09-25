package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/divanvisagie/ui"
)

var apiKey = os.Getenv("GOOGLE_API_KEY")
var window *ui.Window

var targetJSONKeyStore *StringStore
var sourceFilePathStore *StringStore
var translatedJSONFileStore *StringStore

func main() {
	err := ui.Main(func() {
		jsonFileStore := CreateJSONFileStore()
		targetLanguageStore := CreateStringStore()
		targetJSONKeyStore = CreateStringStore()
		sourceFilePathStore = CreateStringStore()
		destinationFilePathStore := CreateStringStore()
		translatedJSONFileStore = CreateStringStore()

		window = ui.NewWindow("JSON Translator", 500, 500, false)
		box := ui.NewVerticalBox()

		saveButton := ui.NewButton("Save")

		saveButton.OnClicked(func(b *ui.Button) {
			if destinationFilePathStore.value == "" {
				ui.MsgBoxError(window, "Target path error", "Target path is not defined")
				return
			}

			data := []byte(translatedJSONFileStore.value)
			ioutil.WriteFile(destinationFilePathStore.value, data, 0644)
		})

		editor := CreateEditor(targetLanguageStore, jsonFileStore)
		go func() {
			for jsonFile := range jsonFileStore.channel {
				editor.SetJSON(jsonFile)
			}
		}()

		box.Append(ui.NewLabel("Select Source File:"), false)
		box.Append(CreateSourceInputBox(targetLanguageStore, destinationFilePathStore, jsonFileStore), false)
		box.Append(ui.NewLabel("Select Destination File:"), false)
		box.Append(CreateDestinationInputBox(destinationFilePathStore), false)
		box.Append(CreateLanguageSelector(targetLanguageStore, destinationFilePathStore), false)
		box.Append(editor.box, true)
		box.Append(saveButton, false)

		window.SetChild(box)
		window.SetMargined(true)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			jsonFileStore.Destroy()
			targetLanguageStore.Destroy()
			targetJSONKeyStore.Destroy()
			sourceFilePathStore.Destroy()
			destinationFilePathStore.Destroy()
			translatedJSONFileStore.Destroy()
			return true
		})
		window.Show()

	})
	if err != nil {
		log.Fatalln(err)
	}
}
