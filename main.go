package main

import (
	"github.com/andlabs/ui"
)

func createSourceInputBox(window *ui.Window) *ui.Box {
	sourcePath := ui.NewEntry()
	openSourceButton := ui.NewButton("...")
	sourceBox := ui.NewHorizontalBox()
	sourceBox.Append(sourcePath, true)
	sourceBox.Append(openSourceButton, false)
	openSourceButton.OnClicked(func(*ui.Button) {
		path := ui.OpenFile(window)
		sourcePath.SetText(path)
	})
	return sourceBox
}

func createDestinationInputBox(window *ui.Window) *ui.Box {
	sourcePath := ui.NewEntry()
	openSourceButton := ui.NewButton("...")
	sourceBox := ui.NewHorizontalBox()
	sourceBox.Append(sourcePath, true)
	sourceBox.Append(openSourceButton, false)
	openSourceButton.OnClicked(func(*ui.Button) {
		path := ui.SaveFile(window)
		sourcePath.SetText(path)
	})
	return sourceBox
}

func createGoogleTranslateSetupBox() *ui.Box {
	box := ui.NewVerticalBox()
	box.Append(ui.NewLabel("Google Translate API key:"), false)
	box.Append(ui.NewEntry(), true)
	return box
}

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("JSON Translator", 500, 200, false)
		greeting := ui.NewLabel("")

		box := ui.NewVerticalBox()

		box.Append(createGoogleTranslateSetupBox(), false)

		box.Append(ui.NewLabel("Select Source File:"), false)
		box.Append(createSourceInputBox(window), false)
		box.Append(ui.NewLabel("Select Destination File:"), false)
		box.Append(createDestinationInputBox(window), false)
		box.Append(greeting, false)

		window.SetChild(box)

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
