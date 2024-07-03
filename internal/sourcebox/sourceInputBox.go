package sourcebox

import (
	"fmt"

	"github.com/andlabs/ui"
	. "json-translator/internal/jsonstore"
	. "json-translator/internal/guesser"
	. "json-translator/pkg/storage"
	. "json-translator/pkg/parser"
)

func CreateSourceInputBox(window *ui.Window ,sourceFilePathStore *StringStore, targetLanguageStore *StringStore, destinationFilePathStore *StringStore, jsonFileStore *JSONFileStore) *ui.Box {
	sourcePath := ui.NewEntry()
	sourcePath.Disable()

	go func() {
		for sourceFilePath := range sourceFilePathStore.Channel {
			sourcePath.SetText(sourceFilePath)
			fmt.Println("Load file:", sourceFilePath)
			jsonFile := ReadJsonFromFile(sourceFilePath)
			jsonFileStore.SetJsonFile(&jsonFile)
			fmt.Println(jsonFile.ToString())

			if destinationFilePathStore.Value == "" {
				guess := GuessTarget(sourceFilePathStore,targetLanguageStore)
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
