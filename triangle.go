// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Kind what kind of triangle?
// Notice KindFromSides() returns this type. Pick a suitable data type.
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
// How many sides are the same?
func KindFromSides(a, b, c float64) Kind {

	var k Kind

	// to be a triangle at all, all sides have to be of length > 0
	if a < 0 || b < 0 || c < 0 {
		k = NaT
		return k
	}

	// the sum of the lengths of any two sides must be greater than or equal to
	// the length of the third side
	if a+b < c {
		k = NaT
		return k
	}

	// how many are the same?
	ab1 := a == b
	ac2 := a == c
	bc3 := b == c
	same := 0
	if ab1 && ac2 {
		same++
	}
	if ab1 && bc3 {
		same++
	}
	if ac2 && bc3 {
		same++
	}

	if same == 3 {
		k = Equ
	} else if same == 2 {
		k = Iso
	} else {
		k = Sca
	}

	return k
}
