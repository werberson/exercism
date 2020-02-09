// package strand implements the function to transcript DNA to RNA
package strand

// ToRNA given a DNA strand, return its RNA complement (per RNA transcription).
func ToRNA(dna string) string {
	rna := make([]byte, len(dna))
	for i, nucleotideDNA := range dna {
		rna[i] = nucleotideToRNA(byte(nucleotideDNA))
	}
	return string(rna)
}

// given a DNA nucleotide, return its RNA nucleotide complement.
func nucleotideToRNA(nucleotide byte) byte {
	var rna byte
	switch nucleotide {
	case 'G':
		rna = 'C'
	case 'C':
		rna = 'G'
	case 'T':
		rna = 'A'
	case 'A':
		rna = 'U'
	}
	return rna
}
