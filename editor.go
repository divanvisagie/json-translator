package main

import (
	"github.com/protonmail/ui"
)

func CreateEditor(ch chan *JSONFile) *ui.Box {

	box := ui.NewHorizontalBox()
	box.SetPadded(true)
	inputJSONControl := ui.NewMultilineEntry()
	outputJSONControl := ui.NewMultilineEntry()
	box.Append(inputJSONControl, true)
	box.Append(outputJSONControl, true)

	go func() {
		for jsonFile := range ch {
			inputJSONControl.SetText(jsonFile.ToString())
		}
	}()

	return box
}
