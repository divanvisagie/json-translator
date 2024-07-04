package language

import (
	"fmt"

	"cloud.google.com/go/translate"
	"github.com/andlabs/ui"
	. "json-translator/internal/state"
	. "json-translator/internal/translation"
)

func populateLanguages(apiKey string) ([]translate.Language, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("apikey is empty")
	}
	languages, err := ListSupportedLanguages(apiKey, "en")
	return languages, err
}

// CreateLanguageSelector creates the combobox that is used to select a languague to translate to
func CreateLanguageSelector(window *ui.Window, apiKey string, state *ApplicationStateStore) *ui.Combobox {

	options, _ := populateLanguages(apiKey)

	validKey := IsAPIKeyValid(apiKey)
	if !validKey {
		ui.MsgBoxError(window, "Error", "API key is not valid, set GOOGLE_API_KEY in your environment variables")
	}

	selector := ui.NewCombobox()

	for _, l := range options {
		selector.Append(l.Tag.String())
	}

	selector.OnSelected(func(c *ui.Combobox) {
		itemIndex := c.Selected()
		languageCode := options[itemIndex].Tag.String()
		state.TargetLanguage.SetValue(languageCode)
		fmt.Println(languageCode)

		// if state.DestinationFilePath.Value == "" {
		// 	_guess := GuessTarget(state.SourceFilePath, state.DestinationFilePath)
		// 	// destinationFilePathStore.SetValue(guess)
		// }
	})
	return selector
}
