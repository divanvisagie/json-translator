package main

import (
	"fmt"

	"github.com/divanvisagie/ui"
)

func CreateEditor(ch chan *JSONFile) *ui.Box {

	var currentJSONFile *JSONFile

	box := ui.NewHorizontalBox()
	box.SetPadded(true)
	inputJSONControl := ui.NewMultilineEntry()
	outputJSONControl := ui.NewMultilineEntry()

	combobox := ui.NewCombobox()

	box.Append(inputJSONControl, true)
	box.Append(combobox, true)
	box.Append(outputJSONControl, true)

	combobox.OnSelected(func(c *ui.Combobox) {
		itemIndex := c.Selected()
		fmt.Println("Selected Item", itemIndex)
	})

	go func() {
		for jsonFile := range ch {
			currentJSONFile = jsonFile
			inputJSONControl.SetText(jsonFile.ToString())
			parsed, _ := jsonFile.Parse()
			for _, object := range parsed {
				for k, _ := range object {
					// fmt.Println(k, v)
					combobox.Append(k)
				}
			}
		}
	}()

	return box
}
