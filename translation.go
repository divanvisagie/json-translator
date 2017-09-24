package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func TranslateText(targetLanguage, text, apiKey string) (string, error) {
	ctx := context.Background()

	fmt.Println(apiKey)

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", err
	}

	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", err
	}
	return resp[0].Text, nil
}

func ListSupportedLanguages(apiKey string, targetLanguage string) ([]translate.Language, error) {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return nil, err
	}

	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	langs, err := client.SupportedLanguages(ctx, lang)
	if err != nil {
		return nil, err
	}

	var languages []translate.Language
	for _, lang := range langs {

		fmt.Printf("%q: %s\n", lang.Tag, lang.Name)
		languages = append(languages, lang)
	}
	return languages, nil
}
