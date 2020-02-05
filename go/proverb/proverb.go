// Package proverb implements the proverb generator function
package proverb

import "fmt"

const (
	sentence = "For want of a %s the %s was lost."
	lastSentence = "And all for the want of a %s."
)

// Proverb given an array of rhyme , generate the relevant proverb 'For Want of a Nail'
func Proverb(rhyme []string) []string {
	proverb := make([]string, len(rhyme))
	lastIndex := len(rhyme) - 1
	for i, r := range rhyme {
		var message string
		if i == lastIndex {
			message = fmt.Sprintf(lastSentence, rhyme[0])
		} else {
			message = fmt.Sprintf(sentence, r, rhyme[i+1])
		}
		proverb[i] = message
	}
	return proverb
}
