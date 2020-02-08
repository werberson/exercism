// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if math.IsInf(a, 1) || math.IsInf(a, -1) ||
		math.IsInf(b, 1) || math.IsInf(b, -1) ||
		math.IsInf(c, 1) || math.IsInf(c, -1) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 || (a+b < c || a+c < b || b+c < a) {
		return NaT
	}

	var k Kind
	if a == b && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
