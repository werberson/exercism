package raindrops

import "strconv"

func Convert(num int) string {
	var str string
	if num% 3 == 0 {
		str = "Pling"
	}
	if num% 5 == 0 {
		str += "Plang"
	}
	if num% 7 == 0 {
		str += "Plong"
	}
	if str == "" {
		str = strconv.Itoa(num)
	}
	return str
}