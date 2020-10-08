package string

import (
	"math/rand"
	"strings"
	"time"
)

// 生成随机字符串
func RandomStr(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func SnakeToCamel(input string) string {
	var result []string
	for _, s := range strings.Split(input, "_") {
		result = append(result, strings.Title(s))
	}
	return strings.Join(result, "")
}
