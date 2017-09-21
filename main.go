package main

import (
	"fmt"
	"log"
	"os"

	"github.com/divanvisagie/ui"
)

var apiKey string = os.Getenv("GOOGLE_API_KEY")
var sourceFilePath string
var destinationFilePath string
var window *ui.Window

var jsonFileStore *JSONFileStore
var targetLanguageStore *StringStore
var targetJSONKeyStore *StringStore
var sourceFilePathStore *StringStore

func createSourceInputBox() *ui.Box {
	sourcePath := ui.NewEntry()

	go func() {
		for sourceFilePath := range sourceFilePathStore.channel {
			sourcePath.SetText(sourceFilePath)
			fmt.Println("Load file:", sourceFilePath)
			jsonFile := ReadJsonFromFile(sourceFilePath)
			jsonFileStore.SetJsonFile(&jsonFile)
			fmt.Println(jsonFile.ToString())
		}
	}()

	sourceFilePathStore.SetValue("")

	openSourceButton := ui.NewButton("...")
	sourceBox := ui.NewHorizontalBox()
	sourceBox.SetPadded(false)
	sourceBox.Append(sourcePath, true)
	sourceBox.Append(openSourceButton, false)
	openSourceButton.OnClicked(func(*ui.Button) {
		sourceFilePath = ui.OpenFile(window)
		sourceFilePathStore.SetValue(sourceFilePath)
	})

	return sourceBox
}

func createDestinationInputBox() *ui.Box {
	sourcePath := ui.NewEntry()
	openDestinationButton := ui.NewButton("...")
	sourceBox := ui.NewHorizontalBox()
	sourceBox.SetPadded(false)
	sourceBox.Append(sourcePath, true)
	sourceBox.Append(openDestinationButton, false)
	openDestinationButton.OnClicked(func(*ui.Button) {
		destinationFilePath = ui.SaveFile(window)
		sourcePath.SetText(destinationFilePath)
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

func createLanguageSelector() *ui.Combobox {

	options := []string{
		"fr",
		"es",
	}

	selector := ui.NewCombobox()

	for _, l := range options {
		selector.Append(l)
	}

	selector.OnSelected(func(c *ui.Combobox) {
		itemIndex := c.Selected()
		languageCode := options[itemIndex]
		targetLanguageStore.SetValue(languageCode)
		fmt.Println(languageCode)
	})
	return selector
}

func main() {
	err := ui.Main(func() {
		window = ui.NewWindow("JSON Translator", 500, 500, false)
		box := ui.NewVerticalBox()

		saveButton := ui.NewButton("Save")

		jsonFileStore = CreateJSONFileStore()
		targetLanguageStore = CreateStringStore()
		targetJSONKeyStore = CreateStringStore()
		sourceFilePathStore = CreateStringStore()

		editor := CreateEditor()

		box.Append(createGoogleTranslateSetupBox(), false)
		box.Append(ui.NewLabel("Select Source File:"), false)
		box.Append(createSourceInputBox(), false)
		box.Append(ui.NewLabel("Select Destination File:"), false)
		box.Append(createDestinationInputBox(), false)
		box.Append(createLanguageSelector(), false)
		box.Append(editor, true)
		box.Append(saveButton, false)

		window.SetChild(box)
		window.SetMargined(true)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			jsonFileStore.Destroy()
			targetLanguageStore.Destroy()
			targetJSONKeyStore.Destroy()
			sourceFilePathStore.Destroy()
			return true
		})
		window.Show()
	})
	if err != nil {
		log.Fatalln(err)
	}
}
