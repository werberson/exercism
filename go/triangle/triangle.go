// Package triangle provides a function to identify the kind of a triangle.
package triangle

import "math"

// Kind represents the kind of triangle.
type Kind int

// triangle kind
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides given the triangles side, returns:
// Equ equilateral if has all three sides the same length.
// Iso isosceles if has at least two sides the same length.
// Sca scalene if has all sides of different lengths.
// NaT Not a triangle if has invalid triangle sides
func KindFromSides(a, b, c float64) Kind {
	if isNotTriangle(a, b, c) {
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

func isNotTriangle(a, b, c float64) bool {
	for _, v := range []float64{a, b, c} {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return true
		}
	}
	if a <= 0 || b <= 0 || c <= 0 || (a+b < c || a+c < b || b+c < a) {
		return true
	}
	return false
}
