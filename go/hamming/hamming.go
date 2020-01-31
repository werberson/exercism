//Package hamming implements a function to calculate the Hamming Distance.
package hamming

import "errors"

// Distance given two DNA sequence, return the Hamming Distance between them.
// returns error if the sequences does not have equal lengths
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the sequences must have equal lengths")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
