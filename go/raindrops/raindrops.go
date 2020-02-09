// Package raindrops provides function to convert a number into raindrop string sounds
package raindrops

import "strconv"

// Convert convert a number into a string that contains raindrop sounds corresponding to certain potential factors.
func Convert(num int) string {
	var str string
	if num%3 == 0 {
		str = "Pling"
	}
	if num%5 == 0 {
		str += "Plang"
	}
	if num%7 == 0 {
		str += "Plong"
	}
	if str == "" {
		str = strconv.Itoa(num)
	}
	return str
}
