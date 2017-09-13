package main

import (
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		sourcePath := ui.NewEntry()
		openSourceButton := ui.NewButton("...")
		greeting := ui.NewLabel("")

		sourceBox := ui.NewHorizontalBox()
		sourceBox.Append(sourcePath, true)
		sourceBox.Append(openSourceButton, false)

		box := ui.NewVerticalBox()

		box.Append(ui.NewLabel("Select Source File:"), false)
		box.Append(sourceBox, false)
		box.Append(greeting, false)

		window := ui.NewWindow("Hello", 500, 500, false)

		window.SetChild(box)
		openSourceButton.OnClicked(func(*ui.Button) {
			path := ui.OpenFile(window)
			greeting.SetText("Selected: " + path)
			sourcePath.SetText(path)
		})
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
