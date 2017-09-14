package main

import (
	"log"
	"os"

	"github.com/andlabs/ui"
)

var API_KEY string = os.Getenv("GOOGLE_API_KEY")
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
	entry.SetText(API_KEY)

	updateButton := ui.NewButton("Update")
	updateButton.OnClicked(func(*ui.Button) {
		os.Setenv("GOOGLE_API_KEY", entry.Text())
		API_KEY = os.Getenv("GOOGLE_API_KEY")
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

		translateButton.OnClicked(func(*ui.Button) {

			ReadJsonFromFile(sourceFilePath)

			word := "Cheese"
			translation, err := TranslateText("fr", word, API_KEY)
			if err != nil {
				log.Fatalln(err.Error())
			}
			outputLabel.SetText("Translated! " + word + " to " + translation)
		})

		box.Append(createGoogleTranslateSetupBox(), false)

		box.Append(ui.NewLabel("Select Source File:"), false)
		box.Append(createSourceInputBox(window), false)
		box.Append(ui.NewLabel("Select Destination File:"), false)
		box.Append(createDestinationInputBox(window), false)
		box.Append(translateButton, false)
		box.Append(outputLabel, false)

		window.SetChild(box)
		window.SetMargined(true)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
