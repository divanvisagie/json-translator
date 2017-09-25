package main

import "github.com/divanvisagie/ui"

func CreateDestinationInputBox(destinationFilePathStore *StringStore) *ui.Box {
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
