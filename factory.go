// Package memefactory assists in the creation of text memes.
package memefactory

import "strings"

// MakeMeme generates a meme.
func MakeMeme(s string) string {
	var newDespacito string

	for i, letter := range s {
		switch string(letter) {
		case "e":
			letter = '3'
		case "t":
			letter = '7'
		case "o":
			letter = '0'
		case "i":
			letter = '1'
		}

		if i%2 != 0 {
			newDespacito += strings.ToUpper(string(letter))
		} else {
			newDespacito += string(letter)
		}
	}

	return newDespacito
}
