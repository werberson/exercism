//Package hamming implements a function to calculate the Hamming Distance.
package hamming

import "errors"

// Distance given two DNA sequences, return the Hamming Distance between them.
// returns error if the sequences does not have equal lengths
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the sequences a and b must have equal lengths")
	}
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
