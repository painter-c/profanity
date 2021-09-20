package profanity

import (
	"strings"
	"unicode/utf8"
)

// HasProfanity checks if the provided text contains and of the
// words in the blacklist.
func HasProfanity(blacklist []string, text string) bool {

	for _, word := range blacklist {
		if strings.Contains(text, word) {
			return true
		}
	}
	return false
}

// MaskAll takes each instance of blacklisted words found in the
// given string and replaces its characters with a new character
// c.
func MaskAll(blacklist []string, text string, c rune) string {

	for _, w := range blacklist {
		masked := strings.Repeat(string(c), utf8.RuneCountInString(w))
		text = strings.ReplaceAll(text, w, masked)
	}

	return text
}

// MaskMiddle takes each instance of blacklisted words found in
// the given string and replaces the middle characters with a new
// character c, leaving the first and last character how they were
// originally.
func MaskMiddle(blacklist []string, text string, c rune) string {

	for _, w := range blacklist {

		fst, _ := utf8.DecodeRuneInString(w)
		lst, _ := utf8.DecodeLastRuneInString(w)
		rcw := utf8.RuneCountInString(w)

		// If there are less than three characters in the word, there will
		// be nothing to censor
		if rcw < 3 {
			continue
		}

		masked := string(fst) + strings.Repeat(string(c), rcw-2) + string(lst)
		text = strings.ReplaceAll(text, w, masked)
	}

	return text
}

// MaskTail takes each instance of blacklisted words found in the
// given string and replaces all of the characters after the first
// with a new character c.
func MaskTail(blacklist []string, text string, c rune) string {
	return ""
}

// ReplaceRandom replaces each blacklisted word in the given string
// with a random word in the provided words slice.
func ReplaceRandom(blacklist []string, words []string, text string) string {
	return ""
}
