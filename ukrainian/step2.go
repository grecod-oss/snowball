package ukrainian

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 2 is the removal of the "і" suffix.
//
func step2(word *snowballword.SnowballWord) bool {
	suffix, _ := word.RemoveFirstSuffixIn(word.RVstart, "і")
	if suffix != "" {
		return true
	}
	return false
}
