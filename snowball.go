package snowball

import (
	"fmt"

	"github.com/grecod-oss/snowball/english"
	"github.com/grecod-oss/snowball/french"
	"github.com/grecod-oss/snowball/norwegian"
	"github.com/grecod-oss/snowball/russian"
	"github.com/grecod-oss/snowball/spanish"
	"github.com/grecod-oss/snowball/swedish"
	"github.com/grecod-oss/snowball/ukrainian"
)

const (
	VERSION string = "v0.0.1"
)

// Stem a word in the specified language.
//
func Stem(word, language string, stemStopWords bool) (stemmed string, err error) {

	var f func(string, bool) string
	switch language {
	case "english":
		f = english.Stem
	case "spanish":
		f = spanish.Stem
	case "french":
		f = french.Stem
	case "russian":
		f = russian.Stem
	case "swedish":
		f = swedish.Stem
	case "ukrainian":
		f = ukrainian.Stem
	case "norwegian":
		f = norwegian.Stem
	default:
		err = fmt.Errorf("Unknown language: %s", language)
		return
	}
	stemmed = f(word, stemStopWords)
	return

}
