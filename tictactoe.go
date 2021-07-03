package main

import (
	"bufio"
	"fmt"
	"os"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

func (b *Board) put(x, y int, u string) bool {
	if b.tokens[x+3*y] == 0 {
		if u == "o" {
			b.tokens[x+3*y] = 1
		} else if u == "x" {
			b.tokens[x+3*y] = 2
		}
		return true
	} else {
		return false
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
			/*if b.tokens[j*3+i] == 1 {
				fmt.Printf("o")
			} else if b.tokens[j*3+i] == 2 {
				fmt.Printf("x")
			} else {
				fmt.Printf(".")
			}*/
			fmt.Printf("%s",b.get(j,i))
		}
		fmt.Println()
	}
}

func (b *Board) isTripleEqual(x int, y int, z int) bool {
	return (b.tokens[x] == b.tokens[y]) && (b.tokens[x] == b.tokens[z]) && b.tokens[x] != 0
}

func (b *Board) check() bool {
	if b.isTripleEqual(0, 1, 2) {
		return true
	} else if b.isTripleEqual(3, 4, 5) {
		return true
	} else if b.isTripleEqual(6, 7, 8) {
		return true
	} else if b.isTripleEqual(0, 3, 6) {
		return true
	} else if b.isTripleEqual(1, 4, 7) {
		return true
	} else if b.isTripleEqual(2, 5, 8) {
		return true
	} else if b.isTripleEqual(0, 4, 8) {
		return true
	} else if b.isTripleEqual(2, 4, 6) {
		return true
	}
	return false
}

func main() {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	p := "Player1"
	s := "o"
	var x, y int
	var r = bufio.NewReader(os.Stdin)
	var step int
	for step = 0; step < 9; step++ {
		fmt.Printf("%s: Input (x,y) ", p)
		fmt.Fscanf(r, "%d,%d\n", &x, &y)
		success := b.put(y, x, s) //reverse x and y
		b.print()
		// if cell is occupied, ask the same player to make input again
		for (!success) {
			fmt.Println("This cell is occupied! Try again")
			fmt.Printf("%s: Input (x,y) ", p)
			fmt.Fscanf(r, "%d,%d\n", &x, &y)
			success = b.put(y, x, s)
			b.print()
		}

		if b.check() == true {
			fmt.Printf("%s won\n", p)
			break
		}

		// swap players
		if s == "o" {
			s = "x"
			p = "Player2"
		} else {
			s = "o"
			p = "Player1"
		}
	}
	if step == 9 {
		fmt.Println("Draw")
	}
}

//wo jiaoyixia
