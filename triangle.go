// Package triangle can determine the kind of triangle that we have sides for
package triangle

import "math"

// Kind what kind of triangle?
// equilateral(3), isosisosceles(2), scalene(1), or not(0)
type Kind int

const (
	// NaT not a triangle
	NaT Kind = 0
	// Equ equilateral triangle has all three sides the same length.
	Equ Kind = 3
	// Iso isosisosceles triangle has at least two sides the same length
	Iso Kind = 2
	// Sca scalene triangle has all sides of different lengths
	Sca Kind = 1
)

// KindFromSides what kind of triangle do the 3 sides form?
// eg How many sides are the same?
func KindFromSides(a, b, c float64) Kind {

	var k Kind

	// to be a triangle at all, all sides have to be of length > 0
	if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
		return k
	}
	// handle Nan and Inf
	// https://play.golang.org/p/qHXIoj-JMAw 0 checks for -Inf and +Inf
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		k = NaT
		return k
	}

	// the sum of the lengths of any two sides must be greater than or equal to
	// the length of the third side
	if a+b < c || a+c < b || b+c < a {
		k = NaT
		return k
	}

	// how many are the same?
	ab1 := a == b
	ac2 := a == c
	bc3 := b == c

	if ab1 && ac2 && bc3 {
		k = Equ
	} else if ab1 || ac2 || bc3 {
		k = Iso
	} else {
		k = Sca
	}

	return k
}
