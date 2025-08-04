// Если тема контекста еще не изучена - пропустить задачу
package main

import (
	"context"
	"fmt"
)

// Есть функция generate(), которая генерит числа
// Функция использует канал отмены. Переделать на контекст.
func generate(cancel context.CancelFunc, start int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i++ {
			select {
			case out <- i:
			default:
				cancel()
			}
		}
	}()
	return out
}

func main() {
	_, cancel := context.WithCancel(context.Background())

	generated := generate(cancel, 11)
	for num := range generated {
		fmt.Print(num, " ")
		if num > 14 {
			break
		}
	}
	fmt.Println()
}
