package countries

import (
	"io"
	"strings"
	"unicode"
)

// NOTE: it works very more faster than strings.Replacer and regexp.Regexp
func textPrepare(text string) string {
	indx := strings.Index(text, "(")
	if indx > -1 {
		text = text[:indx]
	}

	var r rune
	var err error
	reader := strings.NewReader(text)
	text = ""

	for {
		r, _, err = reader.ReadRune()
		if err == io.EOF {
			break
		}
		if unicode.IsLetter(r) {
			text += string(r)
		}
	}

	return strings.ToUpper(text)
}
