package ukrainian

import (
	"github.com/kljensen/snowball/romance"
	"github.com/kljensen/snowball/snowballword"
)

// Checks if a rune is a lowercase Russian vowel.
//
func isLowerVowel(r rune) bool {

	// The Ukrainian vowels are "аеиоуієюя", which
	// are referenced by their unicode code points
	// in the switch statement below.
	switch r {
	case 1072, 1077, 1080, 1086, 1091, 1099, 1101, 1102, 1103:
		return true
	}
	return false
}

// Return `true` if the input `word` is a French stop word.
//
func isStopWord(word string) bool {
	switch word {
	case "і", "в", "що", "він", "на", "я",
		"зі", "як", "а", "то", "всі", "вона", "так", "його",
		"але", "ти", "же", "ви", "за", "би",
		"по", "тільки", "її", "було", "ось", "від",
		"мені", "ще", "ні", "об", "з", "йому", "тепер",
		"коли", "навіть", "ну", "раптом", "чи", "якщо", "вже",
		"або", "не", "бути", "був", "него", "до", "вас",
		"небудь", "знову", "аж", "вам", "адже", "там", "потім",
		"нічого", "їй", "може", "вони", "де",
		"є", "потрібно", "ній", "для", "ми", "тебе", "їх",
		"ніж", "була", "сам", "без", "ніби", "чого",
		"раз", "тоже", "себе", "під", "буде", "тоді",
		"хто", "цей", "того", "цього", "який",
		"зовсім", "ним", "тут", "цьому", "один", "майже",
		"мій", "тим", "щоб", "неї", "зараз", "були", "куди",
		"навіщо", "всіх", "ніколи", "можна", "при", "нарешті",
		"два", "про", "інший", "хоч", "після", "над", "більше",
		"тот", "через", "ці", "нас", "всього", "них",
		"яка", "багато", "хіба", "три", "цю", "моя",
		"втім", "добре", "свою", "цією", "перед", "іноді",
		"краще", "трохи", "тому", "не можна", "такий", "їм", "більш",
		"завжди", "звичайно", "всю", "між":
		return true
	}
	return false
}

// Find the starting point of the regions R1, R2, & RV
//
func findRegions(word *snowballword.SnowballWord) (r1start, r2start, rvstart int) {

	// R1 & R2 are defined in the standard manner.
	r1start = romance.VnvSuffix(word, isLowerVowel, 0)
	r2start = romance.VnvSuffix(word, isLowerVowel, r1start)

	// Set RV, by default, as empty.
	rvstart = len(word.RS)

	// RV is the region after the first vowel, or the end of
	// the word if it contains no vowel.
	//
	for i := 0; i < len(word.RS); i++ {
		if isLowerVowel(word.RS[i]) {
			rvstart = i + 1
			break
		}
	}

	return
}
