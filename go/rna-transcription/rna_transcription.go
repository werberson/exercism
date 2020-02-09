// package strand implements the function to transcript DNA to RNA
package strand

import "strings"

// ToRNA given a DNA strand, return its RNA complement (per RNA transcription).
func ToRNA(dna string) string {
	return strings.Map(nucleotideToRNA, dna)
}

// given a DNA nucleotide, return its RNA nucleotide complement.
// if no complement found, returns param itself.
func nucleotideToRNA(nucleotide rune) rune {
	switch nucleotide {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	}
	return nucleotide
}
