package main

import "math/rand"

// Pouch represents the bag of tokens
type Pouch struct {
	Tokens []TokenColor
}

// Create a new pouch with all tokens
func NewPouch() *Pouch {
	pouch := &Pouch{
		Tokens: make([]TokenColor, 0),
	}

	// Add tokens based on game rules
	for i := 0; i < 23; i++ {
		pouch.Tokens = append(pouch.Tokens, Gray)
	}
	for i := 0; i < 23; i++ {
		pouch.Tokens = append(pouch.Tokens, Blue)
	}
	for i := 0; i < 21; i++ {
		pouch.Tokens = append(pouch.Tokens, Brown)
	}
	for i := 0; i < 19; i++ {
		pouch.Tokens = append(pouch.Tokens, Green)
	}
	for i := 0; i < 19; i++ {
		pouch.Tokens = append(pouch.Tokens, Yellow)
	}
	for i := 0; i < 15; i++ {
		pouch.Tokens = append(pouch.Tokens, Red)
	}

	// Shuffle the tokens
	rand.Shuffle(len(pouch.Tokens), func(i, j int) {
		pouch.Tokens[i], pouch.Tokens[j] = pouch.Tokens[j], pouch.Tokens[i]
	})

	return pouch
}

// Draw tokens from pouch
func (p *Pouch) DrawTokens(count int) []TokenColor {
	if len(p.Tokens) < count {
		count = len(p.Tokens)
	}

	tokens := p.Tokens[:count]
	p.Tokens = p.Tokens[count:]
	return tokens
}
