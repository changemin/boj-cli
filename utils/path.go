package utils

import "strconv"

func GetRangeOfProb(num int) string {
	strNum := strconv.Itoa(num)
	return strNum[:len(strNum)-2] + "00번~" + strNum[:len(strNum)-2] + "99번"
}
