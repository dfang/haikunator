// Package haikunator generates Heroku-like random names to use in your go applications
// Examples:
//   haikunator := haikunator.New()
// 	 haikunator.Haikunate()
package haikunator

import (
	"bytes"
	"math/rand"
	"strings"
	"time"
)

// A Haikunator represents all options needed to use haikunate()
type Haikunator struct {
	Adjectives  []string
	Nouns       []string
	Delimiter   string
	TokenLength int
	TokenHex    bool
	TokenChars  string
	Random      *rand.Rand
}

const (
	numbers = "0123456789"
	hex     = "0123456789abcdef"
)

// New creates a new Haikunator with all default options
func New() *Haikunator {
	return &Haikunator{
		Adjectives:  AdjectiveWords,
		Nouns:       NounWords,
		Delimiter:   "-",
		TokenLength: 4,
		TokenHex:    false,
		TokenChars:  numbers,
		Random:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Haikunate generates a random Heroku-like string
func (h *Haikunator) Haikunate() string {
	adjective := h.randomString(h.Adjectives)
	noun := h.randomString(h.Nouns)
	token := h.buildToken()

	sections := deleteEmpty(adjective, noun, token)
	return strings.Join(sections, h.Delimiter)
}

// buildToken creates and builds random token
func (h *Haikunator) buildToken() string {
	var chars []rune

	if h.TokenHex {
		chars = []rune(hex)
	} else {
		chars = []rune(h.TokenChars)
	}

	size := len(chars)

	if size <= 0 {
		return ""
	}

	var buffer bytes.Buffer

	for i := 0; i < h.TokenLength; i++ {
		index := h.Random.Intn(size)
		buffer.WriteRune(chars[index])
	}

	return buffer.String()
}

// randomString returns random string from slice
func (h *Haikunator) randomString(s []string) string {
	size := len(s)

	if size <= 0 {
		return ""
	}

	return s[h.Random.Intn(size)]
}

// deleteEmpty deletes empty strings from slice
func deleteEmpty(s ...string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
