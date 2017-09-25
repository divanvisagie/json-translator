package main

import (
	"fmt"

	"github.com/divanvisagie/ui"
)

func CreateSourceInputBox(targetLanguageStore *StringStore, destinationFilePathStore *StringStore, jsonFileStore *JSONFileStore) *ui.Box {
	sourcePath := ui.NewEntry()
	sourcePath.Disable()

	go func() {
		for sourceFilePath := range sourceFilePathStore.channel {
			sourcePath.SetText(sourceFilePath)
			fmt.Println("Load file:", sourceFilePath)
			jsonFile := ReadJsonFromFile(sourceFilePath)
			jsonFileStore.SetJsonFile(&jsonFile)
			fmt.Println(jsonFile.ToString())

			if destinationFilePathStore.value == "" {
				guess := GuessTarget(targetLanguageStore)
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
