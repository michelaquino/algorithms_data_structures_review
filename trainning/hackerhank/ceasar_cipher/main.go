package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(caesarCipher("middle-Outz", 2))
	// okffng-Owvb
}

func caesarCipher(s string, k int32) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	encrypted := strings.Builder{}
	for _, letter := range s {
		lowerLetter := strings.ToLower(string(letter))
		if !strings.Contains(alphabet, lowerLetter) {
			encrypted.WriteRune(letter)
			continue
		}

		positionOnAlphabet := strings.Index(alphabet, lowerLetter)

		// Like round robin
		newLetterPosition := (positionOnAlphabet + int(k)) % len(alphabet)
		newLetter := alphabet[newLetterPosition : newLetterPosition+1]

		if unicode.IsUpper(letter) {
			newLetter = strings.ToUpper(newLetter)
		}

		encrypted.WriteString(newLetter)
	}

	return encrypted.String()
}
