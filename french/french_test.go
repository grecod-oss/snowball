package french

import (
	"github.com/kljensen/snowball/romance"
	"github.com/kljensen/snowball/snowballword"
	"testing"
)

// Test stopWords for things we know should be true
// or false.
//
func Test_stopWords(t *testing.T) {
	testCases := []romance.WordBoolTestCase{
		{"eussiez", true},
		{"machine", false},
	}
	romance.RunWordBoolTest(t, isStopWord, testCases)
}

// Test isLowerVowel for things we know should be true
// or false.
//
func Test_isLowerVowel(t *testing.T) {
	testCases := []romance.WordBoolTestCase{
		// These are all vowels.
		{"aeiouyâàëéêèïîôûù", true},
		// None of these are vowels.
		{"cbfqhkl", false},
	}
	romance.RunRunewiseBoolTest(t, isLowerVowel, testCases)
}

// Test capitalization of vowels acting as non-vowels.
//
func Test_capitalizeYUI(t *testing.T) {
	testCases := []struct {
		wordIn  string
		wordOut string
	}{
		{"jouer", "joUer"},
		{"ennuie", "ennuIe"},
		{"yeux", "Yeux"},
		{"quand", "qUand"},
	}

	for _, testCase := range testCases {
		w := snowballword.New(testCase.wordIn)
		capitalizeYUI(w)
		if w.String() != testCase.wordOut {
			t.Errorf("Expect %v -> %v, but got %v", testCase.wordIn, testCase.wordOut, w.String())
		}
	}
}
func Test_findRegions(t *testing.T) {
	testCases := []romance.FindRegionsTestCase{
		{"iriez", 2, 5, 3},
		{"reçoivent", 3, 6, 2},
		{"rébarbatif", 3, 5, 2},
		{"paraîtrons", 3, 6, 3},
		{"prétendus", 4, 6, 3},
		{"boUilli", 3, 5, 2},
		{"destitué", 3, 6, 2},
		{"bataillons", 3, 6, 2},
		{"buffa", 3, 5, 2},
		{"suffisante", 3, 6, 2},
		{"excepté", 2, 5, 4},
		{"audace", 3, 5, 3},
		{"vertueuses", 3, 8, 2},
		{"écrièrent", 2, 6, 4},
		{"provoqUer", 4, 6, 3},
		{"barbotement", 3, 6, 2},
		{"contribua", 3, 7, 2},
		{"ensuit", 2, 6, 4},
		{"confédéré", 3, 6, 2},
		{"affairé", 2, 6, 4},
		{"incompatibles", 2, 5, 4},
		{"talma", 3, 5, 2},
		{"péchais", 3, 7, 2},
		{"abusé", 2, 4, 3},
		{"plaisir", 5, 7, 3},
		{"foretells", 3, 5, 2},
		{"walbah", 3, 6, 2},
		{"confucius", 3, 6, 2},
		{"attelée", 2, 5, 4},
		{"tirailler", 3, 6, 2},
		{"vin", 3, 3, 2},
		{"toucher", 4, 7, 2},
		{"reprendrons", 3, 6, 2},
		{"hé", 2, 2, 2},
		{"intéressant", 2, 5, 4},
		{"malebar", 3, 5, 2},
		{"alimenter", 2, 4, 3},
		{"inventée", 2, 5, 4},
		{"rechargez", 3, 6, 2},
		{"revêtu", 3, 5, 2},
		{"étaYé", 2, 4, 3},
		{"maladresse", 3, 5, 2},
		{"envié", 2, 5, 4},
		{"secoUaIent", 3, 5, 2},
		{"parler", 3, 6, 3},
		{"marécages", 3, 5, 2},
		{"privilèges", 4, 6, 3},
		{"examinez", 2, 4, 3},
		{"contraria", 3, 7, 2},
		{"sotte", 3, 5, 2},
		{"méchantes", 3, 6, 2},
		{"coffres", 3, 7, 2},
		{"tressaillir", 4, 8, 3},
		{"charlatanisme", 4, 7, 3},
		{"appuYais", 2, 5, 4},
		{"interdis", 2, 5, 4},
		{"baissa", 4, 6, 2},
		{"sanglotant", 3, 7, 2},
		{"rencontrerai", 3, 6, 2},
		{"subis", 3, 5, 2},
		{"empestée", 2, 5, 4},
		{"communiqUa", 3, 6, 2},
		{"huit", 4, 4, 2},
		{"heurter", 4, 7, 2},
		{"premiers", 4, 7, 3},
		{"brusqUe", 4, 7, 3},
		{"inanimé", 2, 4, 3},
		{"congédia", 3, 6, 2},
		{"souffrir", 4, 8, 2},
		{"élévations", 2, 4, 3},
		{"sablé", 3, 5, 2},
		{"salure", 3, 5, 2},
		{"résigna", 3, 5, 2},
		{"compatriotes", 3, 6, 2},
		{"écrient", 2, 6, 4},
		{"chanoine", 4, 7, 3},
		{"conçois", 3, 7, 2},
		{"lançaIent", 3, 6, 2},
		{"pékin", 3, 5, 2},
		{"poneYs", 3, 5, 2},
		{"pratiqUer", 4, 6, 3},
		{"bâtonne", 3, 5, 2},
		{"possibilités", 3, 6, 2},
		{"aiguille", 3, 6, 3},
		{"ténor", 3, 5, 2},
		{"déchirés", 3, 6, 2},
		{"anoblit", 2, 4, 3},
		{"tombât", 3, 6, 2},
		{"paralysé", 3, 5, 3},
		{"dot", 3, 3, 2},
		{"aigre", 3, 5, 3},
		{"ramena", 3, 5, 2},
		{"appartiennent", 2, 5, 4},
		{"premières", 4, 7, 3},
		{"tentez", 3, 6, 2},
		{"pari", 3, 4, 3},
		{"coudes", 4, 6, 2},
		{"étonnerait", 2, 4, 3},
		{"embrunir", 2, 6, 5},
		{"mobile", 3, 5, 2},
	}

	romance.RunFindRegionsTest(t, findRegions, testCases)
}

// Test step1, the removal of standard suffixes.
//
func Test_step1(t *testing.T) {
	testCases := []romance.StepTestCase{
		{"rapidement", 3, 5, 2, true, "rapid", 3, 5, 2},
		{"paresseuse", 3, 5, 3, true, "paress", 3, 5, 3},
		{"prosaïqUement", 4, 7, 3, true, "prosaïqU", 4, 7, 3},
		{"nonchalance", 3, 7, 2, true, "nonchal", 3, 7, 2},
		{"apostoliqUes", 2, 4, 3, true, "apostol", 2, 4, 3},
		{"assiduités", 2, 5, 4, true, "assidu", 2, 5, 4},
		{"philosophiqUement", 4, 6, 3, true, "philosoph", 4, 6, 3},
		{"despotiqUement", 3, 6, 2, true, "despot", 3, 6, 2},
		{"incontestablement", 2, 5, 4, true, "incontest", 2, 5, 4},
		{"diminution", 3, 5, 2, true, "diminu", 3, 5, 2},
		{"séditieuse", 3, 5, 2, true, "séditi", 3, 5, 2},
		{"anonymement", 2, 4, 3, true, "anonym", 2, 4, 3},
		{"conservation", 3, 6, 2, true, "conserv", 3, 6, 2},
		{"fâcheuses", 3, 7, 2, true, "fâcheux", 3, 7, 2},
		{"houleuse", 4, 7, 2, true, "houleux", 4, 7, 2},
		{"historiqUes", 3, 6, 2, true, "histor", 3, 6, 2},
		{"impérieusement", 2, 5, 4, true, "impéri", 2, 5, 4},
		{"complaisances", 3, 8, 2, true, "complais", 3, 8, 2},
		{"confessionnaux", 3, 6, 2, true, "confessionnal", 3, 6, 2},
		{"grandement", 4, 7, 3, true, "grand", 4, 5, 3},
		{"passablement", 3, 6, 2, true, "passabl", 3, 6, 2},
		{"strictement", 5, 8, 4, true, "strict", 5, 6, 4},
		{"physiqUement", 4, 6, 3, true, "physiqU", 4, 6, 3},
		{"serieusement", 3, 7, 2, true, "serieux", 3, 7, 2},
		{"roulement", 4, 6, 2, true, "roul", 4, 4, 2},
		{"appartement", 2, 5, 4, true, "appart", 2, 5, 4},
		{"reconnaissance", 3, 5, 2, true, "reconnaiss", 3, 5, 2},
		{"aigrement", 3, 6, 3, true, "aigr", 3, 4, 3},
		{"impertinences", 2, 5, 4, true, "impertinent", 2, 5, 4},
		{"parlement", 3, 6, 3, true, "parl", 3, 4, 3},
		{"malicieux", 3, 5, 2, true, "malici", 3, 5, 2},
		{"suffisance", 3, 6, 2, true, "suffis", 3, 6, 2},
		{"prémédité", 4, 6, 3, true, "préméd", 4, 6, 3},
		{"métalliqUes", 3, 5, 2, true, "métall", 3, 5, 2},
		{"météorologiste", 3, 6, 2, true, "météorolog", 3, 6, 2},
		{"prononciation", 4, 6, 3, true, "prononci", 4, 6, 3},
		{"nombreuse", 3, 8, 2, true, "nombreux", 3, 8, 2},
		{"extatiqUe", 2, 5, 4, true, "extat", 2, 5, 4},
		{"magnifiqUement", 3, 6, 2, true, "magnif", 3, 6, 2},
		{"gymnastiqUe", 3, 6, 2, true, "gymnast", 3, 6, 2},
		{"dramatiqUe", 4, 6, 3, true, "dramat", 4, 6, 3},
		{"simplicité", 3, 7, 2, true, "simpliqU", 3, 7, 2},
		{"roYalistes", 3, 5, 2, true, "roYal", 3, 5, 2},
		{"fortifications", 3, 6, 2, true, "fortif", 3, 6, 2},
		{"attendrissement", 2, 5, 4, true, "attendr", 2, 5, 4},
		{"respectueusement", 3, 6, 2, true, "respectu", 3, 6, 2},
		{"patriotisme", 3, 7, 2, true, "patriot", 3, 7, 2},
		{"curieuse", 3, 7, 2, true, "curieux", 3, 7, 2},
		{"fascination", 3, 6, 2, true, "fascin", 3, 6, 2},
		{"effectivement", 2, 5, 4, true, "effect", 2, 5, 4},
		{"condoléance", 3, 6, 2, true, "condolé", 3, 6, 2},
		{"malignité", 3, 5, 2, true, "malign", 3, 5, 2},
		{"capricieuse", 3, 6, 2, true, "caprici", 3, 6, 2},
		{"applaudissements", 2, 7, 5, true, "applaud", 2, 7, 5},
		{"praticable", 4, 6, 3, true, "pratic", 4, 6, 3},
		{"rivaux", 3, 6, 2, true, "rival", 3, 5, 2},
		{"augmentation", 3, 6, 3, true, "augment", 3, 6, 3},
		{"ameublement", 2, 5, 3, true, "ameubl", 2, 5, 3},
		{"honorables", 3, 5, 2, true, "honor", 3, 5, 2},
		{"effervescence", 2, 5, 4, true, "effervescent", 2, 5, 4},
		{"excentricité", 2, 5, 4, true, "excentr", 2, 5, 4},
		{"misérable", 3, 5, 2, true, "misér", 3, 5, 2},
		{"capitulation", 3, 5, 2, true, "capitul", 3, 5, 2},
		{"enjoUement", 2, 5, 4, true, "enjoU", 2, 5, 4},
		{"sévérité", 3, 5, 2, true, "sévér", 3, 5, 2},
		{"perplexités", 3, 7, 2, true, "perplex", 3, 7, 2},
		{"consentement", 3, 6, 2, true, "consent", 3, 6, 2},
		{"convocation", 3, 6, 2, true, "convoc", 3, 6, 2},
		{"assurances", 2, 5, 4, true, "assur", 2, 5, 4},
		{"ébloUissement", 2, 5, 4, true, "ébloU", 2, 5, 4},
		{"méridionaux", 3, 5, 2, true, "méridional", 3, 5, 2},
		{"dérangements", 3, 5, 2, true, "dérang", 3, 5, 2},
		{"domination", 3, 5, 2, true, "domin", 3, 5, 2},
		{"incroYable", 2, 6, 5, true, "incroY", 2, 6, 5},
		{"réjoUissances", 3, 5, 2, true, "réjoUiss", 3, 5, 2},
		{"décadence", 3, 5, 2, true, "décadent", 3, 5, 2},
		{"bâillement", 4, 7, 2, true, "bâill", 4, 5, 2},
		{"habillement", 3, 5, 2, true, "habill", 3, 5, 2},
		{"irréparablement", 2, 5, 4, true, "irrépar", 2, 5, 4},
		{"diplomatiqUes", 3, 6, 2, true, "diplomat", 3, 6, 2},
		{"distribution", 3, 7, 2, true, "distribu", 3, 7, 2},
		{"pétulance", 3, 5, 2, true, "pétul", 3, 5, 2},
		{"considérable", 3, 6, 2, true, "considér", 3, 6, 2},
		{"éducation", 2, 4, 3, true, "éduc", 2, 4, 3},
		{"indications", 2, 5, 4, true, "indiqU", 2, 5, 4},
		{"cupidité", 3, 5, 2, true, "cupid", 3, 5, 2},
		{"traîtreusement", 5, 9, 3, true, "traîtreux", 5, 9, 3},
		{"silencieuse", 3, 5, 2, true, "silenci", 3, 5, 2},
		{"pessimisme", 3, 6, 2, true, "pessim", 3, 6, 2},
		{"préoccupation", 5, 8, 3, true, "préoccup", 5, 8, 3},
		// Special cases that should return false despite
		// being changed.  They "don't count".
		{"compliment", 3, 7, 2, false, "compli", 3, 6, 2},
		{"vraiment", 5, 7, 3, false, "vrai", 4, 4, 3},
		{"remercîment", 3, 5, 2, false, "remercî", 3, 5, 2},
		{"puissamment", 4, 7, 2, false, "puissant", 4, 7, 2},
		{"absolument", 2, 5, 4, false, "absolu", 2, 5, 4},
		{"décidément", 3, 5, 2, false, "décidé", 3, 5, 2},
		{"condiments", 3, 6, 2, false, "condi", 3, 5, 2},
	}
	romance.RunStepTest(t, step1, testCases)

}

// the removal of Verb suffixes beginning
// with "i" in the RV region.
// Test step1, the removal of standard suffixes.
//
func Test_step2a(t *testing.T) {
	testCases := []romance.StepTestCase{
		{"épanoUit", 2, 4, 3, true, "épanoU", 2, 4, 3},
		{"faillirent", 4, 7, 2, true, "faill", 4, 5, 2},
		{"acabit", 2, 4, 3, true, "acab", 2, 4, 3},
		{"établissait", 2, 4, 3, true, "établ", 2, 4, 3},
		{"découvrir", 3, 6, 2, true, "découvr", 3, 6, 2},
		{"réjoUissait", 3, 5, 2, true, "réjoU", 3, 5, 2},
		{"trahiront", 4, 6, 3, true, "trah", 4, 4, 3},
		{"maintenir", 4, 7, 2, true, "mainten", 4, 7, 2},
		{"vendit", 3, 6, 2, true, "vend", 3, 4, 2},
		{"repartit", 3, 5, 2, true, "repart", 3, 5, 2},
		{"giletti", 3, 5, 2, true, "gilett", 3, 5, 2},
		{"rienzi", 4, 6, 2, true, "rienz", 4, 5, 2},
		{"punie", 3, 5, 2, true, "pun", 3, 3, 2},
		{"accueillir", 2, 7, 4, true, "accueill", 2, 7, 4},
		{"rétablit", 3, 5, 2, true, "rétabl", 3, 5, 2},
		{"ravis", 3, 5, 2, true, "rav", 3, 3, 2},
		{"xviIi", 4, 5, 3, true, "xviI", 4, 4, 3},
		{"blottie", 4, 7, 3, true, "blott", 4, 5, 3},
		{"approfondie", 2, 6, 5, true, "approfond", 2, 6, 5},
		{"infirmerie", 2, 5, 4, true, "infirmer", 2, 5, 4},
		{"scotti", 4, 6, 3, true, "scott", 4, 5, 3},
		{"adoucissait", 2, 5, 3, true, "adouc", 2, 5, 3},
		{"finissait", 3, 5, 2, true, "fin", 3, 3, 2},
		{"promit", 4, 6, 3, true, "prom", 4, 4, 3},
		{"franchies", 4, 9, 3, true, "franch", 4, 6, 3},
		{"franchissant", 4, 8, 3, true, "franch", 4, 6, 3},
		{"micheli", 3, 6, 2, true, "michel", 3, 6, 2},
		{"éteignit", 2, 5, 3, true, "éteign", 2, 5, 3},
		{"puni", 3, 4, 2, true, "pun", 3, 3, 2},
		{"apoplexie", 2, 4, 3, true, "apoplex", 2, 4, 3},
		{"désira", 3, 5, 2, true, "dés", 3, 3, 2},
		{"étourdi", 2, 5, 3, true, "étourd", 2, 5, 3},
		{"giovanni", 4, 6, 2, true, "giovann", 4, 6, 2},
		{"apprécie", 2, 6, 5, true, "appréc", 2, 6, 5},
		{"poésies", 4, 7, 2, true, "poés", 4, 4, 2},
		{"pairie", 4, 6, 2, true, "pair", 4, 4, 2},
		{"sortit", 3, 6, 2, true, "sort", 3, 4, 2},
		{"subi", 3, 4, 2, true, "sub", 3, 3, 2},
		{"aigrirait", 3, 6, 3, true, "aigr", 3, 4, 3},
		{"assailli", 2, 6, 4, true, "assaill", 2, 6, 4},
		{"bertolotti", 3, 6, 2, true, "bertolott", 3, 6, 2},
		{"recouvrir", 3, 6, 2, true, "recouvr", 3, 6, 2},
		{"visconti", 3, 6, 2, true, "viscont", 3, 6, 2},
		{"surgir", 3, 6, 2, true, "surg", 3, 4, 2},
		{"remercie", 3, 5, 2, true, "remerc", 3, 5, 2},
		{"joUissaIent", 3, 5, 2, true, "joU", 3, 3, 2},
		{"bondissant", 3, 6, 2, true, "bond", 3, 4, 2},
		{"saisi", 4, 5, 2, true, "sais", 4, 4, 2},
		{"missouri", 3, 7, 2, true, "missour", 3, 7, 2},
		{"remplirent", 3, 7, 2, true, "rempl", 3, 5, 2},
		{"envahi", 2, 5, 4, true, "envah", 2, 5, 4},
		{"tandis", 3, 6, 2, true, "tand", 3, 4, 2},
		{"trahit", 4, 6, 3, true, "trah", 4, 4, 3},
		{"trahissaIent", 4, 6, 3, true, "trah", 4, 4, 3},
		{"réunie", 4, 6, 2, true, "réun", 4, 4, 2},
		{"avarie", 2, 4, 3, true, "avar", 2, 4, 3},
		{"dilettanti", 3, 5, 2, true, "dilettant", 3, 5, 2},
		{"raidie", 4, 6, 2, true, "raid", 4, 4, 2},
		{"écuries", 2, 4, 3, true, "écur", 2, 4, 3},
		{"recouvrît", 3, 6, 2, true, "recouvr", 3, 6, 2},
		{"parsis", 3, 6, 3, true, "pars", 3, 4, 3},
		{"monti", 3, 5, 2, true, "mont", 3, 4, 2},
		{"reproduisit", 3, 6, 2, true, "reproduis", 3, 6, 2},
		{"étendit", 2, 4, 3, true, "étend", 2, 4, 3},
		{"suffi", 3, 5, 2, true, "suff", 3, 4, 2},
		{"pillaji", 3, 6, 2, true, "pillaj", 3, 6, 2},
		{"rougir", 4, 6, 2, true, "roug", 4, 4, 2},
		{"désirez", 3, 5, 2, true, "dés", 3, 3, 2},
		{"subit", 3, 5, 2, true, "sub", 3, 3, 2},
		{"fondirent", 3, 6, 2, true, "fond", 3, 4, 2},
		{"coqUineries", 3, 6, 2, true, "coqUiner", 3, 6, 2},
		{"venir", 3, 5, 2, true, "ven", 3, 3, 2},
		{"plaidoirie", 5, 8, 3, true, "plaidoir", 5, 8, 3},
		{"fournissant", 4, 7, 2, true, "fourn", 4, 5, 2},
		{"bonzeries", 3, 6, 2, true, "bonzer", 3, 6, 2},
		{"flétri", 4, 6, 3, true, "flétr", 4, 5, 3},
		{"faillit", 4, 7, 2, true, "faill", 4, 5, 2},
		{"hardie", 3, 6, 2, true, "hard", 3, 4, 2},
		{"compagnie", 3, 6, 2, true, "compagn", 3, 6, 2},
		{"vernis", 3, 6, 2, true, "vern", 3, 4, 2},
		{"attendit", 2, 5, 4, true, "attend", 2, 5, 4},
		{"blanchies", 4, 9, 3, true, "blanch", 4, 6, 3},
		{"choisie", 5, 7, 3, true, "chois", 5, 5, 3},
		{"rafraîchir", 3, 7, 2, true, "rafraîch", 3, 7, 2},
		{"choisir", 5, 7, 3, true, "chois", 5, 5, 3},
		{"nourrisse", 4, 7, 2, true, "nourr", 4, 5, 2},
		{"chancellerie", 4, 7, 3, true, "chanceller", 4, 7, 3},
		{"repartie", 3, 5, 2, true, "repart", 3, 5, 2},
		{"redira", 3, 5, 2, true, "red", 3, 3, 2},
		{"sentira", 3, 6, 2, true, "sent", 3, 4, 2},
		{"surgirait", 3, 6, 2, true, "surg", 3, 4, 2},
		{"cani", 3, 4, 2, true, "can", 3, 3, 2},
		{"gratis", 4, 6, 3, true, "grat", 4, 4, 3},
		{"médît", 3, 5, 2, true, "méd", 3, 3, 2},
		{"avertis", 2, 4, 3, true, "avert", 2, 4, 3},
		{"chirurgie", 4, 6, 3, true, "chirurg", 4, 6, 3},
		{"ironie", 2, 4, 3, true, "iron", 2, 4, 3},
		{"punîtes", 3, 5, 2, true, "pun", 3, 3, 2},
		{"compromis", 3, 7, 2, true, "comprom", 3, 7, 2},
		{"simonie", 3, 5, 2, true, "simon", 3, 5, 2},
	}
	romance.RunStepTest(t, step2a, testCases)
}

// Test the removal of Verb suffixes in RV that
// do not begin with "i".
//
func Test_step2b(t *testing.T) {
	testCases := []romance.StepTestCase{
		{"posée", 3, 5, 2, true, "pos", 3, 3, 2},
		{"contentait", 3, 6, 2, true, "content", 3, 6, 2},
		{"évita", 2, 4, 3, true, "évit", 2, 4, 3},
		{"cantonnées", 3, 6, 2, true, "cantonn", 3, 6, 2},
		{"tender", 3, 6, 2, true, "tend", 3, 4, 2},
		{"survenait", 3, 6, 2, true, "surven", 3, 6, 2},
		{"plongeaIent", 4, 8, 3, true, "plong", 4, 5, 3},
		{"modéra", 3, 5, 2, true, "modér", 3, 5, 2},
		{"copier", 3, 6, 2, true, "copi", 3, 4, 2},
		{"bougez", 4, 6, 2, true, "boug", 4, 4, 2},
		{"déploYaIent", 3, 6, 2, true, "déploY", 3, 6, 2},
		{"entendra", 2, 5, 4, true, "entendr", 2, 5, 4},
		{"blâmer", 4, 6, 3, true, "blâm", 4, 4, 3},
		{"déshonorait", 3, 6, 2, true, "déshonor", 3, 6, 2},
		{"concentrés", 3, 6, 2, true, "concentr", 3, 6, 2},
		{"mangeant", 3, 7, 2, true, "mang", 3, 4, 2},
		{"écouteront", 2, 5, 3, true, "écout", 2, 5, 3},
		{"pressaIent", 4, 7, 3, true, "press", 4, 5, 3},
		{"ébréché", 2, 5, 4, true, "ébréch", 2, 5, 4},
		{"frapper", 4, 7, 3, true, "frapp", 4, 5, 3},
		{"côtoYé", 3, 5, 2, true, "côtoY", 3, 5, 2},
		{"réfugié", 3, 5, 2, true, "réfugi", 3, 5, 2},
		{"jeûnant", 4, 6, 2, true, "jeûn", 4, 4, 2},
		{"succombé", 3, 6, 2, true, "succomb", 3, 6, 2},
		{"irrité", 2, 5, 4, true, "irrit", 2, 5, 4},
		{"danger", 3, 6, 2, true, "dang", 3, 4, 2},
		{"sachant", 3, 6, 2, true, "sach", 3, 4, 2},
		{"reparaissaIent", 3, 5, 2, true, "reparaiss", 3, 5, 2},
		{"reconnaissant", 3, 5, 2, true, "reconnaiss", 3, 5, 2},
		{"faisant", 4, 6, 2, true, "fais", 4, 4, 2},
		{"arrangés", 2, 5, 4, true, "arrang", 2, 5, 4},
		{"emparés", 2, 5, 4, true, "empar", 2, 5, 4},
		{"choqUée", 4, 7, 3, true, "choqU", 4, 5, 3},
		{"gênait", 3, 6, 2, true, "gên", 3, 3, 2},
		{"croissante", 5, 8, 3, true, "croiss", 5, 6, 3},
		{"scié", 4, 4, 3, true, "sci", 3, 3, 3},
		{"reconnaissez", 3, 5, 2, true, "reconnaiss", 3, 5, 2},
		{"pliaIent", 5, 7, 3, true, "pli", 3, 3, 3},
		{"expédia", 2, 5, 4, true, "expédi", 2, 5, 4},
		{"déshabillaIent", 3, 6, 2, true, "déshabill", 3, 6, 2},
		{"appréciée", 2, 6, 5, true, "appréci", 2, 6, 5},
		{"amputés", 2, 5, 4, true, "amput", 2, 5, 4},
		{"dominait", 3, 5, 2, true, "domin", 3, 5, 2},
		{"vexantes", 3, 5, 2, true, "vex", 3, 3, 2},
		{"fabriqUées", 3, 6, 2, true, "fabriqU", 3, 6, 2},
		{"retomber", 3, 5, 2, true, "retomb", 3, 5, 2},
		{"exercer", 2, 4, 3, true, "exerc", 2, 4, 3},
		{"entourait", 2, 6, 4, true, "entour", 2, 6, 4},
		{"voYait", 3, 6, 2, true, "voY", 3, 3, 2},
		{"soupait", 4, 7, 2, true, "soup", 4, 4, 2},
		{"apportiez", 2, 5, 4, true, "apport", 2, 5, 4},
		{"tuée", 4, 4, 2, true, "tu", 2, 2, 2},
		{"proposait", 4, 6, 3, true, "propos", 4, 6, 3},
		{"citations", 3, 5, 2, true, "citat", 3, 5, 2},
		{"distinguée", 3, 6, 2, true, "distingu", 3, 6, 2},
		{"parlerez", 3, 6, 3, true, "parl", 3, 4, 3},
		{"stanislas", 4, 6, 3, true, "stanisl", 4, 6, 3},
		{"enlevée", 2, 5, 4, true, "enlev", 2, 5, 4},
		{"irriguaIent", 2, 5, 4, true, "irrigu", 2, 5, 4},
		{"contenant", 3, 6, 2, true, "conten", 3, 6, 2},
		{"empêchèrent", 2, 5, 4, true, "empêch", 2, 5, 4},
		{"inspirées", 2, 6, 5, true, "inspir", 2, 6, 5},
		{"basée", 3, 5, 2, true, "bas", 3, 3, 2},
		{"consultait", 3, 6, 2, true, "consult", 3, 6, 2},
		{"retardait", 3, 5, 2, true, "retard", 3, 5, 2},
		{"enlevât", 2, 5, 4, true, "enlev", 2, 5, 4},
		{"convenaIent", 3, 6, 2, true, "conven", 3, 6, 2},
		{"portât", 3, 6, 2, true, "port", 3, 4, 2},
		{"admirée", 2, 5, 4, true, "admir", 2, 5, 4},
		{"copiée", 3, 6, 2, true, "copi", 3, 4, 2},
		{"démenaIent", 3, 5, 2, true, "démen", 3, 5, 2},
		{"fortifiées", 3, 6, 2, true, "fortifi", 3, 6, 2},
		{"apercevrait", 2, 4, 3, true, "apercevr", 2, 4, 3},
		{"risqUer", 3, 7, 2, true, "risqU", 3, 5, 2},
		{"réclamer", 3, 6, 2, true, "réclam", 3, 6, 2},
		{"tremblaIent", 4, 8, 3, true, "trembl", 4, 6, 3},
		{"calomnier", 3, 5, 2, true, "calomni", 3, 5, 2},
		{"réclamée", 3, 6, 2, true, "réclam", 3, 6, 2},
		{"déposât", 3, 5, 2, true, "dépos", 3, 5, 2},
		{"filé", 3, 4, 2, true, "fil", 3, 3, 2},
		{"déchirée", 3, 6, 2, true, "déchir", 3, 6, 2},
		{"prononça", 4, 6, 3, true, "prononç", 4, 6, 3},
		{"précédé", 4, 6, 3, true, "précéd", 4, 6, 3},
		{"asseYait", 2, 5, 4, true, "asseY", 2, 5, 4},
		{"emploYés", 2, 6, 5, true, "emploY", 2, 6, 5},
		{"chagriner", 4, 7, 3, true, "chagrin", 4, 7, 3},
		{"dévorât", 3, 5, 2, true, "dévor", 3, 5, 2},
		{"remonté", 3, 5, 2, true, "remont", 3, 5, 2},
		{"emploYant", 2, 6, 5, true, "emploY", 2, 6, 5},
		{"redoublait", 3, 6, 2, true, "redoubl", 3, 6, 2},
		{"marchant", 3, 7, 2, true, "march", 3, 5, 2},
		{"pétrifiée", 3, 6, 2, true, "pétrifi", 3, 6, 2},
		{"enlevées", 2, 5, 4, true, "enlev", 2, 5, 4},
		{"donnassent", 3, 6, 2, true, "donn", 3, 4, 2},
		{"recomptait", 3, 5, 2, true, "recompt", 3, 5, 2},
		{"masqUait", 3, 8, 2, true, "masqU", 3, 5, 2},
		{"renouvelèrent", 3, 6, 2, true, "renouvel", 3, 6, 2},
		{"recoucher", 3, 6, 2, true, "recouch", 3, 6, 2},
		{"abrégea", 2, 5, 4, true, "abrég", 2, 5, 4},
		{"flattait", 4, 8, 3, true, "flatt", 4, 5, 3},
	}
	romance.RunStepTest(t, step2b, testCases)
}

// Test the cleaning up of "Y" and "ç" suffixes.
//
func Test_step3(t *testing.T) {
	testCases := []romance.StepTestCase{
		{"ennuY", 5, 5, 5, true, "ennui", 5, 5, 5},
		{"envoY", 5, 5, 4, true, "envoi", 5, 5, 4},
		{"aboY", 4, 4, 3, true, "aboi", 4, 4, 3},
		{"essaY", 5, 5, 4, true, "essai", 5, 5, 4},
		{"effroY", 6, 6, 6, true, "effroi", 6, 6, 6},
		{"désennuY", 8, 8, 8, true, "désennui", 8, 8, 8},
		{"renvoY", 6, 6, 6, true, "renvoi", 6, 6, 6},
		{"prononç", 7, 7, 3, true, "prononc", 7, 7, 3},
		{"asseY", 5, 5, 5, true, "assei", 5, 5, 5},
		{"croY", 4, 4, 3, true, "croi", 4, 4, 3},
		{"asseY", 5, 5, 4, true, "assei", 5, 5, 4},
		{"plaç", 4, 4, 3, true, "plac", 4, 4, 3},
		{"ennuY", 5, 5, 5, true, "ennui", 5, 5, 5},
		{"impaY", 5, 5, 5, true, "impai", 5, 5, 5},
		{"déploY", 6, 6, 2, true, "déploi", 6, 6, 2},
		{"avanç", 5, 5, 3, true, "avanc", 5, 5, 3},
		{"recommenç", 9, 9, 2, true, "recommenc", 9, 9, 2},
		{"pitoY", 5, 5, 5, true, "pitoi", 5, 5, 5},
		{"renvoY", 6, 6, 6, true, "renvoi", 6, 6, 6},
		{"choY", 4, 4, 4, true, "choi", 4, 4, 4},
		{"effroY", 6, 6, 6, true, "effroi", 6, 6, 6},
		{"forç", 4, 4, 2, true, "forc", 4, 4, 2},
		{"envoY", 5, 5, 5, true, "envoi", 5, 5, 5},
		{"paY", 3, 3, 3, true, "pai", 3, 3, 3},
		{"bunhY", 5, 5, 2, true, "bunhi", 5, 5, 2},
	}
	romance.RunStepTest(t, step3, testCases)
}

// Test
//
func Test_step4(t *testing.T) {
	testCases := []romance.StepTestCase{
		{"défendues", 3, 5, 2, true, "défendu", 3, 5, 2},
		{"mormones", 3, 6, 2, true, "mormon", 3, 6, 2},
		{"souvienne", 4, 7, 2, true, "souvienn", 4, 7, 2},
		{"poumons", 4, 6, 2, true, "poumon", 4, 6, 2},
		{"relâche", 3, 5, 2, true, "relâch", 3, 5, 2},
		{"ressource", 3, 7, 2, true, "ressourc", 3, 7, 2},
		{"petits", 3, 5, 2, true, "petit", 3, 5, 2},
		{"obstacles", 2, 6, 5, true, "obstacl", 2, 6, 5},
		{"voisine", 4, 6, 2, true, "voisin", 4, 6, 2},
		{"tunnels", 3, 6, 2, true, "tunnel", 3, 6, 2},
		{"politesse", 3, 5, 2, true, "politess", 3, 5, 2},
		{"obéisse", 2, 5, 3, true, "obéiss", 2, 5, 3},
		{"brûlons", 4, 6, 3, true, "brûlon", 4, 6, 3},
		{"tâchons", 3, 6, 2, true, "tâchon", 3, 6, 2},
		{"gothiqUes", 3, 6, 2, true, "gothiqU", 3, 6, 2},
		{"acqUise", 2, 6, 5, true, "acqUis", 2, 6, 5},
		{"pigeons", 3, 6, 2, true, "pigeon", 3, 6, 2},
		{"focs", 3, 4, 2, true, "foc", 3, 3, 2},
		{"profondeurs", 4, 6, 3, true, "profondeur", 4, 6, 3},
		{"mettrons", 3, 7, 2, true, "mettron", 3, 7, 2},
		{"bavards", 3, 5, 2, true, "bavard", 3, 5, 2},
		{"nigauds", 3, 6, 2, true, "nigaud", 3, 6, 2},
		{"déesse", 4, 6, 2, true, "déess", 4, 5, 2},
		{"libraires", 3, 7, 2, true, "librair", 3, 7, 2},
		{"sentimentales", 3, 6, 2, true, "sentimental", 3, 6, 2},
		{"libre", 3, 5, 2, true, "libr", 3, 4, 2},
		{"matérielles", 3, 5, 2, true, "matériell", 3, 5, 2},
		{"habitudes", 3, 5, 2, true, "habitud", 3, 5, 2},
		{"blushes", 4, 7, 3, true, "blush", 4, 5, 3},
		{"suppose", 3, 6, 2, true, "suppos", 3, 6, 2},
		{"décrépitude", 3, 6, 2, true, "décrépitud", 3, 6, 2},
		{"incluse", 2, 6, 5, true, "inclus", 2, 6, 5},
		{"files", 3, 5, 2, true, "fil", 3, 3, 2},
		{"côtes", 3, 5, 2, true, "côt", 3, 3, 2},
		{"spirales", 4, 6, 3, true, "spiral", 4, 6, 3},
		{"bamboches", 3, 6, 2, true, "bamboch", 3, 6, 2},
		{"qUête", 4, 5, 3, true, "qUêt", 4, 4, 3},
		{"siècles", 4, 7, 2, true, "siècl", 4, 5, 2},
		{"glisse", 4, 6, 3, true, "gliss", 4, 5, 3},
		{"carrosses", 3, 6, 2, true, "carross", 3, 6, 2},
		{"supprime", 3, 7, 2, true, "supprim", 3, 7, 2},
		{"officielle", 2, 5, 4, true, "officiell", 2, 5, 4},
		{"vifs", 3, 4, 2, true, "vif", 3, 3, 2},
		{"adresses", 2, 5, 4, true, "adress", 2, 5, 4},
		{"hussards", 3, 6, 2, true, "hussard", 3, 6, 2},
		{"colle", 3, 5, 3, true, "coll", 3, 4, 3},
		{"amendes", 2, 4, 3, true, "amend", 2, 4, 3},
		{"qUeUe", 4, 5, 3, true, "qUeU", 4, 4, 3},
		{"écharpe", 2, 5, 4, true, "écharp", 2, 5, 4},
		{"débute", 3, 5, 2, true, "début", 3, 5, 2},
		{"refuse", 3, 5, 2, true, "refus", 3, 5, 2},
		{"légers", 3, 5, 2, true, "léger", 3, 5, 2},
		{"entrailles", 2, 7, 5, true, "entraill", 2, 7, 5},
		{"écarlate", 2, 4, 3, true, "écarlat", 2, 4, 3},
		{"manufacturières", 3, 5, 2, true, "manufacturi", 3, 5, 2},
		{"instruire", 2, 8, 6, true, "instruir", 2, 8, 6},
		{"danses", 3, 6, 2, true, "dans", 3, 4, 2},
		{"lits", 3, 4, 2, true, "lit", 3, 3, 2},
		{"cours", 4, 5, 2, true, "cour", 4, 4, 2},
		{"belgirate", 3, 6, 2, true, "belgirat", 3, 6, 2},
		{"délire", 3, 5, 2, true, "délir", 3, 5, 2},
		{"offenses", 2, 5, 4, true, "offens", 2, 5, 4},
		{"athènes", 2, 5, 4, true, "athèn", 2, 5, 4},
		{"alphabets", 2, 6, 5, true, "alphabet", 2, 6, 5},
		{"ascagne", 2, 5, 4, true, "ascagn", 2, 5, 4},
		{"lièvre", 4, 6, 2, true, "lièvr", 4, 5, 2},
		{"hercule", 3, 6, 2, true, "hercul", 3, 6, 2},
		{"casqUe", 3, 6, 2, true, "casqU", 3, 5, 2},
		{"cachons", 3, 6, 2, true, "cachon", 3, 6, 2},
		{"herbe", 3, 5, 2, true, "herb", 3, 4, 2},
		{"banqUette", 3, 7, 2, true, "banqUett", 3, 7, 2},
		{"actuelles", 2, 6, 4, true, "actuell", 2, 6, 4},
		{"intercession", 2, 5, 4, true, "intercess", 2, 5, 4},
		{"pêle", 3, 4, 2, true, "pêl", 3, 3, 2},
		{"grossières", 4, 8, 3, true, "grossi", 4, 6, 3},
		{"qUelle", 4, 6, 3, true, "qUell", 4, 5, 3},
		{"séduits", 3, 6, 2, true, "séduit", 3, 6, 2},
		{"vengeance", 3, 7, 2, true, "vengeanc", 3, 7, 2},
		{"indécentes", 2, 5, 4, true, "indécent", 2, 5, 4},
		{"bergères", 3, 6, 2, true, "bergèr", 3, 6, 2},
		{"fenestrelles", 3, 5, 2, true, "fenestrell", 3, 5, 2},
		{"croupe", 5, 6, 3, true, "croup", 5, 5, 3},
		{"légitime", 3, 5, 2, true, "légitim", 3, 5, 2},
		{"ferrare", 3, 6, 2, true, "ferrar", 3, 6, 2},
		{"briqUe", 4, 6, 3, true, "briqU", 4, 5, 3},
		{"étrangère", 2, 5, 4, true, "étrangèr", 2, 5, 4},
		{"arqUés", 2, 6, 5, true, "arqUé", 2, 5, 5},
		{"guèbres", 4, 7, 2, true, "guèbr", 4, 5, 2},
		{"partons", 3, 6, 3, true, "parton", 3, 6, 3},
		{"distingue", 3, 6, 2, true, "distingu", 3, 6, 2},
		{"paratonnerres", 3, 5, 3, true, "paratonnerr", 3, 5, 3},
		{"anonyme", 2, 4, 3, true, "anonym", 2, 4, 3},
		{"volutes", 3, 5, 2, true, "volut", 3, 5, 2},
		{"décence", 3, 5, 2, true, "décenc", 3, 5, 2},
		{"coupure", 4, 6, 2, true, "coupur", 4, 6, 2},
		{"avarice", 2, 4, 3, true, "avaric", 2, 4, 3},
		{"sensible", 3, 6, 2, true, "sensibl", 3, 6, 2},
		{"cramponne", 4, 7, 3, true, "cramponn", 4, 7, 3},
		{"sympathise", 3, 6, 2, true, "sympathis", 3, 6, 2},
		{"assidue", 2, 5, 4, true, "assidu", 2, 5, 4},
	}
	romance.RunStepTest(t, step4, testCases)
}

// Test a large set of words for which we know
// the correct stemmed form.
//
func Test_FrenchVocabulary(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"battements", "batt"},
		{"mélangé", "mélang"},
		{"impériales", "impérial"},
		{"paragraphe", "paragraph"},
		{"charité", "charit"},
		{"reproche", "reproch"},
		{"belvédère", "belvéder"},
		{"illisible", "illisibl"},
		{"pleurs", "pleur"},
		{"passait", "pass"},
		{"heaviest", "heaviest"},
		{"correspondance", "correspond"},
		{"c", "c"},
		{"profitable", "profit"},
		{"remontrance", "remontr"},
		{"ramasseraient", "ramass"},
		{"arrivera", "arriv"},
		{"canta", "cant"},
		{"évanouie", "évanou"},
		{"bleuâtres", "bleuâtr"},
		{"achetées", "achet"},
		{"bazars", "bazar"},
		{"affections", "affect"},
		{"luttent", "luttent"},
		{"recouvra", "recouvr"},
		{"regorgent", "regorgent"},
		{"pruderie", "pruder"},
		{"entomologique", "entomolog"},
		{"jansénisme", "jansen"},
		{"tourne", "tourn"},
		{"tuer", "tu"},
		{"concluantes", "conclu"},
		{"subi", "sub"},
		{"agent", "agent"},
		{"instantanément", "instantan"},
		{"gustave", "gustav"},
		{"colossales", "colossal"},
		{"nothing", "nothing"},
		{"quantièmes", "quantiem"},
		{"aidez", "aid"},
		{"horlogerie", "horloger"},
		{"ranimer", "ranim"},
		{"landau", "landau"},
		{"mêler", "mêl"},
		{"scrupuleusement", "scrupul"},
		{"poitrail", "poitrail"},
		{"chaudement", "chaud"},
		{"impiété", "impiet"},
		{"redoublaient", "redoubl"},
		{"punira", "pun"},
		{"proposa", "propos"},
		{"envolés", "envol"},
		{"réparer", "répar"},
		{"inventer", "invent"},
		{"précision", "précis"},
		{"déguisa", "déguis"},
		{"plantations", "plantat"},
		{"appliqua", "appliqu"},
		{"plat", "plat"},
		{"préfète", "préfet"},
		{"baisers", "baiser"},
		{"calmèrent", "calm"},
		{"tressé", "tress"},
		{"consulta", "consult"},
		{"dédaigneux", "dédaign"},
		{"dithyrambe", "dithyramb"},
		{"obligera", "oblig"},
		{"nommés", "nomm"},
		{"mousseux", "mousseux"},
		{"pusillanimes", "pusillanim"},
		{"richissime", "richissim"},
		{"weber", "web"},
		{"groupes", "group"},
		{"rentra", "rentr"},
		{"persécuté", "persécut"},
		{"nuiraient", "nuir"},
		{"ayant", "ayant"},
		{"joueraient", "jou"},
		{"attenante", "atten"},
		{"formait", "form"},
		{"encombrées", "encombr"},
		{"sifflait", "siffl"},
		{"lire", "lir"},
		{"faciliter", "facilit"},
		{"casse", "cass"},
		{"remit", "rem"},
		{"profond", "profond"},
		{"sortez", "sort"},
		{"boiteux", "boiteux"},
		{"flatteuses", "flatteux"},
		{"plafonds", "plafond"},
		{"trahît", "trah"},
		{"lesquelles", "lesquel"},
		{"fantaisies", "fantais"},
		{"séduite", "séduit"},
		{"consolée", "consol"},
		{"estomac", "estomac"},
		{"adverbe", "adverb"},
		{"promenés", "promen"},
		{"côte", "côt"},
		{"flegme", "flegm"},
		{"végétaient", "véget"},
		{"annoncerait", "annonc"},
		{"quais", "quais"},
		{"hissa", "hiss"},
		{"protection", "protect"},
		{"destine", "destin"},
		{"justice", "justic"},
		{"fili", "fil"},
		{"conduite", "conduit"},
		{"narra", "narr"},
		{"torturé", "tortur"},
		{"couloirs", "couloir"},
		{"bronché", "bronch"},
		{"oeuvres", "oeuvr"},
		{"retire", "retir"},
		{"laisserai", "laiss"},
		{"rassura", "rassur"},
		{"leipsick", "leipsick"},
		{"gâte", "gât"},
		{"désormais", "désorm"},
		{"pain", "pain"},
		{"pianos", "pianos"},
		{"opérée", "oper"},
		{"effrayèrent", "effrai"},
		{"sachez", "sach"},
		{"répétées", "répet"},
		{"time", "tim"},
		{"golgonda", "golgond"},
		{"occupèrent", "occup"},
		{"embrasserais", "embrass"},
		{"dévorante", "dévor"},
		{"soutenant", "souten"},
		{"voluptueuse", "voluptu"},
		{"vicomtes", "vicomt"},
		{"constante", "const"},
		{"admirable", "admir"},
		{"déroger", "dérog"},
		{"survit", "surv"},
		{"manquerais", "manqu"},
		{"remontrer", "remontr"},
		{"exercent", "exercent"},
		{"outrageantes", "outrag"},
		{"dépôt", "dépôt"},
		{"engagées", "engag"},
		{"rouvray", "rouvray"},
		{"comprenez", "compren"},
		{"imprudentes", "imprudent"},
		{"billards", "billard"},
		{"tremblante", "trembl"},
		{"impie", "impi"},
		{"peu", "peu"},
		{"indigène", "indigen"},
		{"social", "social"},
		{"consigne", "consign"},
		{"emporterait", "emport"},
		{"rocky", "rocky"},
		{"cosmopolite", "cosmopolit"},
		{"police", "polic"},
		{"jeun", "jeun"},
		{"lourdes", "lourd"},
		{"extraordinaire", "extraordinair"},
		{"dérangeait", "dérang"},
		{"long", "long"},
		{"empressées", "empress"},
		{"capitulation", "capitul"},
		{"giration", "girat"},
		{"guidés", "guid"},
		{"bourbiers", "bourbi"},
		{"provisions", "provis"},
		{"dois", "dois"},
		{"squelette", "squelet"},
		{"extravagante", "extravag"},
		{"bruns", "brun"},
		{"considérerais", "consider"},
		{"entièrement", "entier"},
		{"suffocations", "suffoc"},
		{"diminue", "diminu"},
		{"froissants", "froiss"},
		{"avalé", "aval"},
		{"détacher", "détach"},
		{"remplace", "remplac"},
		{"exagérait", "exager"},
		{"élévations", "élev"},
		{"exagérant", "exager"},
		{"promenaient", "promen"},
		{"antidatée", "antidat"},
		{"touchait", "touch"},
		{"aimerait", "aim"},
		{"lope", "lop"},
		{"tranchait", "tranch"},
		{"environnent", "environnent"},
		{"inondation", "inond"},
		{"frayeur", "frayeur"},
		{"solaire", "solair"},
		{"oysters", "oyster"},
		{"rêveuse", "rêveux"},
		{"concession", "concess"},
		{"existé", "exist"},
		{"promener", "promen"},
	}
	for _, testCase := range testCases {
		result := Stem(testCase.in, true)
		if result != testCase.out {
			t.Errorf("Expected %v -> %v, but got %v", testCase.in, testCase.out, result)
		}
	}
}
