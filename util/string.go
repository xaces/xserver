package util

import (
	"strconv"
	"strings"
)

func StringToIntSlice(str, sep string) []int {
	strv := strings.Split(str, sep)
	var intv []int
	for _, v := range strv {
		if v == "" {
			continue
		}
		val, _ := strconv.Atoi(v)
		intv = append(intv, val)
	}
	return intv
}
