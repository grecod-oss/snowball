package ukrainian

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 1 is the removal of standard suffixes, all of which must
// occur in RV.
//
//
// Search for a PERFECTIVE GERUND ending. If one is found remove it, and
// that is then the end of step 1. Otherwise try and remove a REFLEXIVE
// ending, and then search in turn for (1) an ADJECTIVAL, (2) a VERB or
// (3) a NOUN ending. As soon as one of the endings (1) to (3) is found
// remove it, and terminate step 1.
//
func step1(word *snowballword.SnowballWord) bool {

	// `stop` will be used to signal early termination
	var stop bool

	// Search for a PERFECTIVE GERUND ending
	stop = removePerfectiveGerundEnding(word)
	if stop {
		return true
	}

	// Next remove reflexive endings
	word.RemoveFirstSuffixIn(word.RVstart, "ся", "сь")

	// Next remove adjectival endings
	stop = removeAdjectivalEnding(word)
	if stop {
		return true
	}

	// Next remove verb endings
	stop = removeVerbEnding(word)
	if stop {
		return true
	}

	// Next remove noun endings
	suffix, _ := word.RemoveFirstSuffixIn(word.RVstart,
		"ик", "ник", "івник", "чик", "щик", "ив",
		"альник", "ильник", "ільник", "альність", "аль", "ень",
		"ець", "єць", "ість", "тель", "инн", "иння", "інн", "іння",
		"анн", "ання", "янн", "яння", "енн", "ення", "ен", "еня", "єн", "єня",
		"ечок", "єчок", "ечк", "ечка", "ечк", "ечко", "єчк", "єчко", "ичок",
		"ичка", "ик", "иц", "иця", "енк", "енко", "єнк", "єнко", "еньк", "енько",
		"єньк", "єнько", "иськ", "исько", "їськ", "їсько", "ищ", "ище", "їщ",
		"їще", "ович", "івн", "івна", "ївн", "ївна", "івк", "івка", "ївк",
		"ївка", "овк", "овка", "ок", "ир", "ист", "изм", "ір", "іст", "ізм",
	)
	if suffix != "" {
		return true
	}

	return false
}

// Remove perfective gerund endings and return true if one was removed.
//
func removePerfectiveGerundEnding(word *snowballword.SnowballWord) bool {
	suffix, suffixRunes := word.FirstSuffixIn(word.RVstart, len(word.RS),
		"чи", "вшись", "чись", "ши", "ти", "ий", "вши",
	)
	switch suffix {
	case "в", "вши", "вшись":

		// These are "Group 1" perfective gerund endings.
		// Group 1 endings must follow а (a) or я (ia) in RV.
		if precededByARinRV(word, len(suffixRunes)) == false {
			suffix = ""
		}

	}

	if suffix != "" {
		word.RemoveLastNRunes(len(suffixRunes))
		return true
	}
	return false
}

// Remove adjectival endings and return true if one was removed.
//
func removeAdjectivalEnding(word *snowballword.SnowballWord) bool {

	// Remove adjectival endings.  Start by looking for
	// an adjective ending.
	//
	suffix, _ := word.RemoveFirstSuffixIn(word.RVstart,
		"ого", "ому", "ої", "ою", "їх", "і", "им", "ім", "ими",
		"ій", "а", "іх", "ю", "у", "ій",
	)
	if suffix != "" {

		// We found an adjective ending.  Remove optional participle endings.
		//
		newSuffix, newSuffixRunes := word.FirstSuffixIn(word.RVstart, len(word.RS),
			"ачий", "ячий", "учий", "ючий",
			"тий", "ний", "не", "на", "те",
		)
		switch newSuffix {
		case "тий", "ний", "не", "на", "те":

			// These are "Group 1" participle endings.
			// Group 1 endings must follow а (a) or я (ia) in RV.
			if precededByARinRV(word, len(newSuffixRunes)) == false {
				newSuffix = ""
			}
		}

		if newSuffix != "" {
			word.RemoveLastNRunes(len(newSuffixRunes))
		}
		return true
	}
	return false
}

// Remove verb endings and return true if one was removed.
//
func removeVerbEnding(word *snowballword.SnowballWord) bool {
	suffix, suffixRunes := word.FirstSuffixIn(word.RVstart, len(word.RS),
		"уть", "ють", "ать", "ять", "ти", "іти", "ати", "їти", "и", "імо",
		"ім", "іть", "е", "є", "ї",
	)
	switch suffix {
	case "уть", "ють", "ать", "ять":

		// These are "Group 1" verb endings.
		// Group 1 endings must follow а (a) or я (ia) in RV.
		if precededByARinRV(word, len(suffixRunes)) == false {
			suffix = ""
		}

	}

	if suffix != "" {
		word.RemoveLastNRunes(len(suffixRunes))
		return true
	}
	return false
}

// There are multiple classes of endings that must be
// preceded by а (a) or я (ia) in RV in order to be removed.
//
func precededByARinRV(word *snowballword.SnowballWord, suffixLen int) bool {
	idx := len(word.RS) - suffixLen - 1
	if idx >= word.RVstart && (word.RS[idx] == 'а' || word.RS[idx] == 'я') {
		return true
	}
	return false
}
