package sliceutil

import (
	"math/rand"
	"time"
)

// Shuffle 随机打散切片
func Shuffle(slice []interface{}) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
