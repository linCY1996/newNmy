package util

import (
	"bytes"
	"math/rand"
)

var (
	s  = "abcdefghijklmnopqrstuvwxyz0987654321-"
	l  = len(s)
	s1 = "159357284"
	le = len(s1)
)

func RandStr() string {
	var buf bytes.Buffer
	for i := 0; i < 16; i++ {
		buf.WriteByte(s[rand.Intn(l)])
	}
	return buf.String()
}

func Randone() string {
	var buf bytes.Buffer
	for i := 0; i < 1; i++ {
		buf.WriteByte(s1[rand.Intn(le)])
	}
	return buf.String()
}

func Rand() string {
	var buf bytes.Buffer
	for i := 0; i < 6; i++ {
		buf.WriteByte(s1[rand.Intn(le)])
	}
	return buf.String()
}
