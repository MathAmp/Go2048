package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

const (
	LEFT int = iota
	RIGHT
	UP
	DOWN
	ESC
	PASS
)

func detectControl() int {
	switch ev := term.PollEvent(); ev.Type {
	case term.EventKey:
		switch ev.Key {
		case term.KeyEsc:
			term.Sync()
			return ESC
		case term.KeyArrowDown:
			term.Sync()
			return DOWN
		case term.KeyArrowUp:
			term.Sync()
			return UP
		case term.KeyArrowRight:
			term.Sync()
			return RIGHT
		case term.KeyArrowLeft:
			term.Sync()
			return LEFT
		default:
			term.Sync()
			return PASS
		}
	case term.EventError:
		panic(ev.Err)
	default:
		return PASS
	}
}

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	ch := make(chan int)
	go func() {
		for {
			ch <- detectControl()
		}
	}()

	for {
		fmt.Println("Press ESC to quit")

		v := <-ch
		if v == ESC {
			break
		}
		if v == PASS {
			continue
		}
		fmt.Println(v)
	}
}
