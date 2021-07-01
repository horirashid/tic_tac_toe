package main

import (
	"bufio"
	"fmt"
	"os"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

func (b *Board) put(x, y int, u string) {
	if b.tokens[x+3*y] == 0 {
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
	if b.tokens[x+3*y] == 1 {
		return "o"
	} else if b.tokens[x+3*y] == 2 {
		return "x"
	}
	return "."
}

func (b *Board) print() {
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			if b.tokens[j*3+i] == 1 {
				fmt.Printf("o")
			} else if b.tokens[j*3+i] == 2 {
				fmt.Printf("x")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	p := "Player1"
	s := "o"
	var x, y int
	var r = bufio.NewReader(os.Stdin)
	for true {
		fmt.Printf("%s: Input (x,y) ", p)
		fmt.Fscanf(r, "%d,%d\n", &x, &y)
		b.put(x, y, s)

		//xxxxxxxxxxxxxxxxxxxxxx
		// TODO -try to output the whole board here instead of just current cell
		//fmt.Printf("%s\n", b.get(x, y))
		b.print()

		// swap players
		if s == "o" {
			s = "x"
			p = "Player2"
		} else {
			s = "o"
			p = "Player1"
		}
	}
}

//wo jiaoyixia
