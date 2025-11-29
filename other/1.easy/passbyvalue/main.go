package main

import "fmt"

// Что будет выведено и почему? Подумать, после запустить программу. Исправить.

func main() {
	var value int = 3
	p := &value

	fmt.Println(*p)

	changePointer(p)
	fmt.Println(*p)
}

func changePointer(p *int) {
	v := 1
	p = &v
}
