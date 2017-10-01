package goutils

import (
	"math/rand"
	"time"
)

const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randNum     = rand.NewSource(time.Now().UnixNano())
)

// GenerateRandomKey returns a random string of len n
func GenerateRandomKey(n int) []byte {
	b := make([]byte, n)
	for i, cache, remain := n-1, randNum.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randNum.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}
