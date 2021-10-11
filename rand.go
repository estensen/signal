package main

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func RandString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
