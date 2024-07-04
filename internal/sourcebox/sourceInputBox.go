package sourcebox

import (
	"fmt"

	"github.com/andlabs/ui"
	. "json-translator/internal/guesser"
	. "json-translator/pkg/parser"
	. "json-translator/internal/state"
)

func CreateSourceInputBox(window *ui.Window , state *ApplicationStateStore) *ui.Box {
	sourcePath := ui.NewEntry()
	sourcePath.Disable()

	go func() {
		for sfp := range state.SourceFilePath.Channel {
			sourcePath.SetText(sfp)
			fmt.Println("Load file:", sfp)
			jf := ReadJsonFromFile(sfp)
			state.JSONFileStore.SetJsonFile(&jf)
			fmt.Println(jf.ToString())

			if state.DestinationFilePath.Value == "" {
				guess := GuessTarget(state.SourceFilePath,state.TargetLanguage)
				state.DestinationFilePath.SetValue(guess)
			}
		}
	}()

	state.SourceFilePath.SetValue("")

	openSourceButton := ui.NewButton("...")
	sb := ui.NewHorizontalBox()
	sb.SetPadded(false)
	sb.Append(sourcePath, true)
	sb.Append(openSourceButton, false)
	openSourceButton.OnClicked(func(*ui.Button) {
		sourceFilePath := ui.OpenFile(window)
		state.SourceFilePath.SetValue(sourceFilePath)
	})

	return sb
}
