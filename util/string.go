package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/nats-io/nuid"
	uuid "github.com/satori/go.uuid"
)

//UUID 生成Guid字串
func UUID() string {
	u := uuid.NewV4()
	return u.String()
}

// NUID ID
func NUID() string {
	return nuid.Next()
}

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
