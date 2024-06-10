// Libs
package utils

import "strings"

// Data
const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Functions
/*
A method to convert the given text to Caesar Cipher using the provided shift.

## Arguments:
- `text`: The text to be converted.
- `shift`: The shift containing the numbers of alphabet's characters to be shifted.
*/
func ConvertToCesar(text string, shift int) string {
	// Loop over all the text.
	var convertedText string
	for _, char := range strings.Split(text, "") {
		// Treat the selected char.
		isUpper := char == strings.ToUpper(char)

		// Try to find the char in the alphabet.
		index := strings.Index(ALPHABET, strings.ToUpper(char))
		if index == -1 {
			convertedText += char
			continue
		}

		// Use the shift to determine the wanted char.
		index += shift
		if index >= len(ALPHABET) {
			index -= len(ALPHABET)
		}

		// Add the char to the convertedText.
		char = string(rune(ALPHABET[index]))
		if !isUpper {
			char = strings.ToLower(char)
		}
		convertedText += char
	}
	return convertedText
}
