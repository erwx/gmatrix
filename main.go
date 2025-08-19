package main

import (
	"fmt"
	"golang.org/x/term"
	"math/rand/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func fall(column int, height int, spd time.Duration, maxRow int) {
	for {
		for i := 1; i <= maxRow; i++ {
			if i > 1 {
				fmt.Printf("\033[%d;%dH\033[0;32m%c", i-1, column, rand.IntN(93)+33)
			}
			fmt.Printf("\033[%d;%dH\033[1;37m%c", i, column, rand.IntN(93)+33)

			if i > height {
				fmt.Printf("\033[%d;%dH ", i-height, column)
			}

			time.Sleep(spd)
		}

		fmt.Printf("\033[%d;%dH\033[0;32m%c", maxRow, column, rune(rand.IntN(93)+33))
		
		for i := maxRow - height + 1; i <= maxRow; i++ {
			if i > 0 {
				fmt.Printf("\033[%d;%dH ", i, column)
			}
			time.Sleep(spd)
		}
		time.Sleep(time.Duration(rand.IntN(150)) * time.Millisecond)
	}
}

func main() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGINT)

	go func() {
		<-interruptChan
		fmt.Println("\033[0m")
		fmt.Print("\033[2J\033[1;1H\033[?25l")
		fmt.Print("\033[2J\033[?25h")
		os.Exit(0)
	}()

	ncol, nrow, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Printf("Error getting terminal size: %v\n", err)
		return
	}
	fmt.Print("\033[2J\033[1;1H\033[?25l")
	defer fmt.Print("\n\033[2J\033[?25h")

	for col := 1; col <= ncol; col = col + 2 {
		randomSpeed := time.Duration(rand.IntN(100)+30) * time.Millisecond
		randomHeight := rand.IntN(nrow) + 15
		go fall(col, randomHeight, randomSpeed, nrow)
	}
	select {}
}
