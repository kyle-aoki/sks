package main

import (
	"crypto/rand"
	"math/big"
)

const pool = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generatePassword(n int) string {
	var bytes []byte
	for i := 0; i < n; i++ {
		randomNumber := must(rand.Int(rand.Reader, big.NewInt(int64(len(pool)))))
		bytes = append(bytes, pool[randomNumber.Int64()])
	}
	return string(bytes)
}
