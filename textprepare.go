package countries

import (
	"regexp"
	"strings"
)

// rePrepare - for func textPrepare()
var rePrepare *regexp.Regexp

// stPrepare - for func textPrepare()
var stPrepare *strings.Replacer

func init() {
	rePrepare = regexp.MustCompile("\n|\t|\r|\f|\v|\"|`|'|’|′|´|῾|᾿|″|˝|̏ |̏‴|“|”|‘| ҃|‹|›|«|»|^|ˆ|!|¡|¿|؟|#|№|÷|[|]|{|}|⟨|⟩|\\|/|⁄|¦|;|:|,|·|•|°|º|ﾟ|˚|˳|%|‰|‱|¨|…| | | |-|‐|‑|‒|―|—|–|~|_|¯|@|¶|§|©|®|℠|™|℗")
	stPrepare = strings.NewReplacer("&", "AND", ")", "", "?", "", ".", "", "|", "", "+", "", "*", "")
}

func textPrepare(text string) string {
	text = rePrepare.ReplaceAllString(text, "")
	indx := strings.Index(text, "(")
	if indx > -1 {
		text = text[:indx]
	}
	text = stPrepare.Replace(text)
	return strings.ToUpper(text)
}
