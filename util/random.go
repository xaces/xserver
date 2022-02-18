package util

import (
	"fmt"
	"math/rand"
)

// RandomString 生成指定位数的字符
func RandomString(width int) string {
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
