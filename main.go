package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type caller func()

func main() {
	moveChan := make(chan bool)
	start := time.Now()
	robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("Ran for :", time.Since(start))
		os.Exit(0)
	})
	s := robotgo.EventStart()
	go func() {
		select {
		case <-robotgo.EventProcess(s):

		}
	}()

	fmt.Println("Press ctrl+shift+q to exit")
	ticker := time.NewTicker(5 * time.Second)

	runner := func(fn caller) {
		for {
			select {
			case <-ticker.C:
				fn()
				// Start this up as a go routine soo that it does not block the timer update
				// Blocking channel takes in a single instruction at a time
				go func() {
					moveChan <- true
				}()
			}
		}
	}

	go func() {
		for range moveChan {
			randomX := rand.Intn(100) * (-1 + rand.Intn(3))
			randomY := rand.Intn(100) * (-1 + rand.Intn(3))
			fmt.Println(randomX, randomY)
			robotgo.KeyTap("shift")
			robotgo.MoveRelative(randomX, randomY)
		}
	}()

	runner(func() {})

}
