package main

import (
	"github.com/andlabs/ui"
)

func CreateEditor(ch chan string) *ui.Box {
	box := ui.NewHorizontalBox()
	label := ui.NewLabel("Test")

	box.Append(label, false)

	go func() {
		for text := range ch {
			label.SetText(text)
		}
	}()

	return box
}
