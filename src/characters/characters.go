package characters

import (
	"math/rand"
)

var Hiragana = []rune("ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん")
var Katakana = []rune("ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ")
var Kanji = []rune("一二三四五六七八九十百千万山川田人心力日月火水木金土王天気花雨")

func RandomJapaneseCharacter() rune {
	set := rand.Intn(3)
	switch set {
	case 0:
		return Hiragana[rand.Intn(len(Hiragana))]
	case 1:
		return Katakana[rand.Intn(len(Katakana))]
	default:
		return Kanji[rand.Intn(len(Kanji))]
	}
}

func RandomJapaneseCharacterWithSpaces() rune {
	if rand.Intn(10) > 2 {
		return ' '
	}
	return RandomJapaneseCharacter()
}
