package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Канал завершения
// Есть функция, которая произносит текст пословно (с некоторыми задержками):

func say(id int, text string) {
	for _, word := range strings.Fields(text) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

// Запускаем несколько одновременных воркеров, по одной на каждую фразу:

func main() {
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}
	for idx, phrase := range phrases {
		go say(idx+1, phrase)
	}
}

// Программа ничего не печатает — функция main() завершается до того, как отработает хотя бы один воркер:
// Использовать канал для завершения. Пролистай, если нужна подсказка.
//
//
//
//
//
//
// say(done chan<- struct{}, id int, phrase string).
