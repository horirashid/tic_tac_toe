package main

import "fmt"

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

func (b *Board) put(x, y int, u string) {
	if b.tokens[x + 3*y] == 0 {
		if u == "o" {
			b.tokens[x+3*y] = 1
		} else if u == "x" {
			b.tokens[x+3*y] = 2
		}
	} else {
		fmt.Println("This cell is occupied!")
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x + 3*y] == 1 {
		return "o"
	} else if b.tokens[x + 3*y] == 2 {
		return "x"
	}
	return "."
}