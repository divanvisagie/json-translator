package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/translate"
	"github.com/divanvisagie/ui"
)

var apiKey = os.Getenv("GOOGLE_API_KEY")
var window *ui.Window

var jsonFileStore *JSONFileStore
var targetLanguageStore *StringStore
var targetJSONKeyStore *StringStore
var sourceFilePathStore *StringStore
var destinationFilePathStore *StringStore
var translatedJSONFileStore *StringStore

func guessTarget() string {
	if sourceFilePathStore.value == "" {
		return ""
	}
	if targetLanguageStore.value == "" {
		return ""
	}

	p := sourceFilePathStore.value
	split := strings.Split(p, ".")
	split[0] += ("-" + targetLanguageStore.value)
	return strings.Join(split, ".")
}

func createSourceInputBox() *ui.Box {
	sourcePath := ui.NewEntry()

	go func() {
		for sourceFilePath := range sourceFilePathStore.channel {
			sourcePath.SetText(sourceFilePath)
			fmt.Println("Load file:", sourceFilePath)
			jsonFile := ReadJsonFromFile(sourceFilePath)
			jsonFileStore.SetJsonFile(&jsonFile)
			fmt.Println(jsonFile.ToString())

			if destinationFilePathStore.value == "" {
				guess := guessTarget()
				destinationFilePathStore.SetValue(guess)
			}
		}
	}()

	sourceFilePathStore.SetValue("")

	openSourceButton := ui.NewButton("...")
	sourceBox := ui.NewHorizontalBox()
	sourceBox.SetPadded(false)
	sourceBox.Append(sourcePath, true)
	sourceBox.Append(openSourceButton, false)
	openSourceButton.OnClicked(func(*ui.Button) {
		sourceFilePath := ui.OpenFile(window)
		sourceFilePathStore.SetValue(sourceFilePath)
	})

	return sourceBox
}

func createDestinationInputBox() *ui.Box {
	sourcePath := ui.NewEntry()
	destinationFilePathStore.SetValue("")

	openDestinationButton := ui.NewButton("...")
	sourceBox := ui.NewHorizontalBox()
	sourceBox.SetPadded(false)
	sourceBox.Append(sourcePath, true)
	sourceBox.Append(openDestinationButton, false)

	go func() {
		for destination := range destinationFilePathStore.channel {
			sourcePath.SetText(destination)
		}
	}()

	openDestinationButton.OnClicked(func(*ui.Button) {
		destinationFilePath := ui.SaveFile(window)
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

func populateLanguages() ([]translate.Language, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("apikey is empty")
	}
	languages, err := ListSupportedLanguages(apiKey, "en")
	return languages, err
}

func createLanguageSelector() *ui.Combobox {

	options, _ := populateLanguages()

	validKey := IsAPIKeyValid(apiKey)
	if !validKey {
		ui.MsgBoxError(window, "Error", "API key is not valid, set GOOGLE_API_KEY in your environment variables")
	}

	selector := ui.NewCombobox()

	for _, l := range options {
		selector.Append(l.Tag.String())
	}

	selector.OnSelected(func(c *ui.Combobox) {
		itemIndex := c.Selected()
		languageCode := options[itemIndex].Tag.String()
		targetLanguageStore.SetValue(languageCode)
		fmt.Println(languageCode)

		if destinationFilePathStore.value == "" {
			guess := guessTarget()
			destinationFilePathStore.SetValue(guess)
		}
	})
	return selector
}

func main() {
	err := ui.Main(func() {
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

		jsonFileStore = CreateJSONFileStore()
		targetLanguageStore = CreateStringStore()
		targetJSONKeyStore = CreateStringStore()
		sourceFilePathStore = CreateStringStore()
		destinationFilePathStore = CreateStringStore()
		translatedJSONFileStore = CreateStringStore()

		editor := CreateEditor()
		go func() {
			for jsonFile := range jsonFileStore.channel {
				editor.SetJSON(jsonFile)
			}
		}()

		box.Append(createGoogleTranslateSetupBox(), false)
		box.Append(ui.NewLabel("Select Source File:"), false)
		box.Append(createSourceInputBox(), false)
		box.Append(ui.NewLabel("Select Destination File:"), false)
		box.Append(createDestinationInputBox(), false)
		box.Append(createLanguageSelector(), false)
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
