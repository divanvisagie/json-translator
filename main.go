package main

import (
	"fmt"
	"os"

	"github.com/divanvisagie/ui"
)

var apiKey string = os.Getenv("GOOGLE_API_KEY")
var sourceFilePath string

func createSourceInputBox(window *ui.Window) *ui.Box {
	sourcePath := ui.NewEntry()
	openSourceButton := ui.NewButton("...")
	sourceBox := ui.NewHorizontalBox()
	sourceBox.SetPadded(false)
	sourceBox.Append(sourcePath, true)
	sourceBox.Append(openSourceButton, false)
	openSourceButton.OnClicked(func(*ui.Button) {
		sourceFilePath = ui.OpenFile(window)
		sourcePath.SetText(sourceFilePath)
	})
	return sourceBox
}

func createDestinationInputBox(window *ui.Window) *ui.Box {
	sourcePath := ui.NewEntry()
	openSourceButton := ui.NewButton("...")
	sourceBox := ui.NewHorizontalBox()
	sourceBox.SetPadded(false)
	sourceBox.Append(sourcePath, true)
	sourceBox.Append(openSourceButton, false)
	openSourceButton.OnClicked(func(*ui.Button) {
		path := ui.SaveFile(window)
		sourcePath.SetText(path)
	})
	return sourceBox
}

func createGoogleTranslateSetupBox() *ui.Box {
	entry := ui.NewEntry()
	entry.SetText(apiKey)

	updateButton := ui.NewButton("Update")
	updateButton.OnClicked(func(*ui.Button) {
		os.Setenv("GOOGLE_API_KEY", entry.Text())
		apiKey = os.Getenv("GOOGLE_API_KEY")
	})

	box := ui.NewVerticalBox()
	box.SetPadded(true)
	box.Append(ui.NewLabel("Google Translate API key:"), false)
	box.Append(entry, true)
	box.Append(updateButton, false)

	return box
}

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("JSON Translator", 500, 500, false)
		box := ui.NewVerticalBox()
		outputLabel := ui.NewLabel("")
		translateButton := ui.NewButton("Translate")

		jsonChannel := make(chan *JSONFile)

		translateButton.OnClicked(func(*ui.Button) {

			jsonFile := ReadJsonFromFile(sourceFilePath)

			jsonChannel <- &jsonFile

			fmt.Println(jsonFile.ToString())

			// word := "Cheese"
			// translation, err := TranslateText("fr", word, apiKey)
			// if err != nil {
			// 	log.Fatalln(err.Error())
			// }
			// outputLabel.SetText("Translated! " + word + " to " + translation)
		})

		editor := CreateEditor(jsonChannel)

		box.Append(createGoogleTranslateSetupBox(), false)

		box.Append(ui.NewLabel("Select Source File:"), false)
		box.Append(createSourceInputBox(window), false)
		box.Append(ui.NewLabel("Select Destination File:"), false)
		box.Append(createDestinationInputBox(window), false)
		box.Append(translateButton, false)
		box.Append(outputLabel, false)
		box.Append(editor, true)

		window.SetChild(box)
		window.SetMargined(true)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			close(jsonChannel)
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
