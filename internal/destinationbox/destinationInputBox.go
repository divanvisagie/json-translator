package destinationbox

import (
	. "json-translator/pkg/storage"

	"github.com/andlabs/ui"
)

func CreateDestinationInputBox(window *ui.Window, destinationFilePathStore *StringStore) *ui.Box {
	sourcePath := ui.NewEntry()
	sourcePath.Disable()
	destinationFilePathStore.SetValue("")

	openDestinationButton := ui.NewButton("...")
	sourceBox := ui.NewHorizontalBox()
	sourceBox.SetPadded(false)
	sourceBox.Append(sourcePath, true)
	sourceBox.Append(openDestinationButton, false)

	go func() {
		for destination := range destinationFilePathStore.Channel {
			sourcePath.SetText(destination)
		}
	}()

	openDestinationButton.OnClicked(func(*ui.Button) {
		destinationFilePath := ui.SaveFile(window)
		sourcePath.SetText(destinationFilePath)
	})
	return sourceBox
}
