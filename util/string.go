package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func StringToIntSlice(str, sep string) []uint64 {
	strv := strings.Split(str, sep)
	var intv []uint64
	for _, v := range strv {
		if v == "" {
			continue
		}
		val, _ := strconv.Atoi(v)
		intv = append(intv, uint64(val))
	}
	return intv
}

// StringRandom 生成指定位数的字符
func StringRandom(width int) string {
	// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	// b := make([]rune, width)
	// for i := range b {
	// 	b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterRunes))]
	// }
	// return string(b)
	randBytes := make([]byte, width/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)

}
