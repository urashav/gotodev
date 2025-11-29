package main

import "fmt"

// Что будет выведено и почему? Подумать, после запустить программу. Исправить.

type Figure interface {
	Area() float64
}

type square struct {
	a float64
}

func (s *square) Area() float64 {
	return s.a * s.a
}

func main() {
	var a *square
	figure := Figure(a)
	fmt.Printf("Is figure: %t\n", figure == nil)
}
