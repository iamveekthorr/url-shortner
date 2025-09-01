// Package utils provides utility function.
package utils

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func generateAlphabets() string {
	var sb strings.Builder
	var alphabets string

	for r := 'a'; r <= 'z'; r++ {
		R := unicode.ToUpper(r)
		sb.WriteRune(r)
		sb.WriteRune(R)

		// Assign the alphabets to the string.
		alphabets = sb.String()
	}

	return alphabets
}

func MakeShortCode(length int) string {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder

	const digits = "0123456789"
	letters := generateAlphabets()

	charset := digits + letters

	for i := 0; i <= length; i++ {
		idx := src.Intn(len(charset))
		sb.WriteByte(charset[idx])
	}

	return sb.String()
}
