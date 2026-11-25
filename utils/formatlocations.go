package utils

import (
	"strings"
)

func capitalize(mot string) string {
	if len(mot) == 0 {
		return mot
	}
	premiere := mot[:1]
	reste := mot[1:]
	return strings.ToUpper(premiere) + reste
}

func formatWords(texte string) string {
	texte = strings.ReplaceAll(texte, "_", " ")
	mots := strings.Fields(texte)
	for i, mot := range mots {
		mots[i] = capitalize(mot)
	}
	return strings.Join(mots, " ")
}

func FormatLocation(s string) string {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return s
	}
	var pays string
	sigles := map[string]bool{"usa": true, "uk": true}
	mot := parts[1]
	if sigles[mot] {
		pays = strings.ToUpper(mot)
	} else {
		pays = formatWords(parts[1])
	}
	ville := formatWords(parts[0])
	return ville + ", " + pays
}
