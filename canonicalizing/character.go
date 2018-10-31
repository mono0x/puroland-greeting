package canonicalizing

import (
	"regexp"

	"golang.org/x/text/width"
)

var characterRules = map[string]string{
	"バットばつ丸":    "バッドばつ丸",
	"ミルク":       "みるく",
	"うさはな":      "ウサハナ",
	"キティ":       "キティ・ホワイト",
	"ポムポムプリン2":  "ポムポムプリン",
	"パティ & ジミィ": "パティ&ジミィ",
	"みんなのたあ坊":   "たあ坊",
	"らららフローラ":   "ラララフローラ",
	"りっぷちゃん":    "りっぷ",
	"もっぷ":       "モップ",
}

var ignoreCostumes = map[string]struct{}{
	"パパ":     {},
	"ママ":     {},
	"おじいちゃん": {},
	"おばあちゃん": {},
}

var characterCostumeRe = regexp.MustCompile(`\A(.+?)\s*(?:\((.+?)\))?\z`)

func CanonicalizeCharacterName(source string) (string, string) {
	folded := width.Fold.String(source)

	var character string
	var costume string

	submatches := characterCostumeRe.FindStringSubmatch(folded)
	if len(submatches) >= 2 {
		character = submatches[1]
		if len(submatches) >= 3 {
			costume = submatches[2]
		}
	} else {
		character = folded
	}

	if replaced, ok := characterRules[character]; ok {
		character = replaced
	}

	if costume != "" {
		if _, ok := ignoreCostumes[costume]; ok {
			costume = ""
		}
	}
	return character, costume
}
