#!/bin/bash

set -e

# Define new directory structure
mkdir -p cmd/json-translator
mkdir -p internal/editor internal/guesser internal/jsonstore internal/language internal/sourcebox internal/destinationbox internal/translation
mkdir -p pkg/parser pkg/storage
mkdir -p docs/images

# Move files into new structure
mv main.go cmd/json-translator/

mv editor.go internal/editor/
mv guesser.go internal/guesser/
mv jsonFileStore.go internal/jsonstore/
mv languageSelector.go internal/language/
mv sourceInputBox.go internal/sourcebox/
mv destinationInputBox.go internal/destinationbox/
mv translation.go internal/translation/

mv parser.go pkg/parser/
mv stringStore.go pkg/storage/

# Move existing docs/images directory and screenshot.png
if [ -d "docs/images" ]; then
    mv docs/images/screenshot.png docs/images/
else
    mv docs/images/screenshot.png docs/
fi

# Remove old docs directory if empty
if [ -d "docs" ] && [ ! "$(ls -A docs)" ]; then
    rmdir docs
fi


echo "Project files have been rearranged successfully."
