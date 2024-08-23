package main

import (
	day_5 "advent_of_code_2023/day_5/puzzle_2"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)
	go func() {
		day_5.Puzzle_2()
		done <- true
	}()

	select {
	case sig := <-sigs:
		fmt.Println("Received signal:", sig)
	case <-done:
		fmt.Println("Puzzle_2 completed")
	}
}

//day_1.TrebuchetGo()
//day_2.Puzzle_1()
//day_2.Puzzle_2()
//day_3.Puzzle_1()
//day_3.Puzzle_2()
//day_4.Puzzle_2()
