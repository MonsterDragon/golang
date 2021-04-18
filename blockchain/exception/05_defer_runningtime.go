package main

import (
	"fmt"
	"time"
)

func main() {
	startAt := time.Now()
	defer func() {fmt.Println(time.Since(startAt))}()

	time.Sleep(time.Second)
}
