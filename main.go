package workspace

import "strconv"

func TestFunc(s string) string {
	return s
}

func InMainFunc(num int) string {
	return strconv.Itoa(num)
}
