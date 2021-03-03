package main

import (
	"fmt"
	"math/rand"
	"time"
)

var strSlice4RandomStr = [5]string{
	"0123456789",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"abcdefghijklmnopqrstuvwxyz",
	"!@#$%^&*()_-+=[{]}\\|,<.>/?",
	" "}

var randG *rand.Rand

func init() {
	randG = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// generate random string with fixed length
// type is a 5 bit length binary number which represent numbers, upper letters,
// lower letters, special characters and space from lower to higher
func RandomStr(l int, t int) (s string, err error) {
	if l < 1 || t < 1 {
		err = fmt.Errorf("invalid length or type number")
		return
	}

	var (
		i      int
		bts    []byte
		result []byte
	)

	bts = make([]byte, 0, 89)
	for i = 0; i < 5; i++ {
		if t>>i%2 == 1 {
			bts = append(bts, []byte(strSlice4RandomStr[i])...)
		}
	}

	// println(string(bts), len(bts))
	result = make([]byte, l)

	for i = 0; i < l; i++ {
		result[i] = bts[randG.Intn(len(bts))]
	}

	s = string(result)

	return
}

// generate alphanumeric string with fixed length
func RandAlphanumeric(length int) string {
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randG.Intn(len(charset))]
	}

	return string(b)
}
