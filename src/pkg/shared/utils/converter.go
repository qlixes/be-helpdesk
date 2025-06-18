package utils

import "strconv"

func StrToInt(s string) int {
	num, err := strconv.Atoi(s)

	if err != nil {
		panic(err.Error())
	}

	return num
}
