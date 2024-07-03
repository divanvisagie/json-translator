package language

import (
	"fmt"

	"cloud.google.com/go/translate"
	"github.com/andlabs/ui"
	. "json-translator/internal/guesser"
	. "json-translator/internal/translation"
	. "json-translator/pkg/storage"
)

func populateLanguages(apiKey string) ([]translate.Language, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("apikey is empty")
	}
	languages, err := ListSupportedLanguages(apiKey, "en")
	return languages, err
}

// CreateLanguageSelector creates the combobox that is used to select a languague to translate to
func CreateLanguageSelector(window *ui.Window, apiKey string, sourceLanguageStore *StringStore, targetLanguageStore *StringStore, destinationFilePathStore *StringStore) *ui.Combobox {

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
		targetLanguageStore.SetValue(languageCode)
		fmt.Println(languageCode)

		if destinationFilePathStore.Value == "" {
			guess := GuessTarget(sourceLanguageStore, targetLanguageStore)
			destinationFilePathStore.SetValue(guess)
		}
	})
	return selector
}
