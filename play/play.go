package main

import (
	"fmt"

	"github.com/mathamp/go2048"
	term "github.com/nsf/termbox-go"
)

type Command int

const (
	LEFT Command = iota
	RIGHT
	UP
	DOWN
	ESC
	PASS
)

func detectControl() Command {
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

func processGame(gs *go2048.GameState, command Command) go2048.StatusCode {
	switch command {
	case UP:
		{
			return gs.Process(go2048.UP)
		}
	case DOWN:
		{
			return gs.Process(go2048.DOWN)
		}
	case LEFT:
		{
			return gs.Process(go2048.LEFT)
		}
	case RIGHT:
		{
			return gs.Process(go2048.RIGHT)
		}
	default:
		{
			panic("Invalid Command")
		}
	}
}

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	ch := make(chan Command)
	go func() {
		for {
			ch <- detectControl()
		}
	}()

	gs := go2048.NewGameState()
	gs.InitRandomBlock()

	for {
		fmt.Println("Press ESC to quit")
		fmt.Println(gs.String())

		v := <-ch
		if v == ESC {
			break
		}
		if v == PASS {
			continue
		}

		if processGame(&gs, v) == go2048.TERMINATED {
			fmt.Println(gs.String())
			fmt.Println("Game Over!")
			<-ch
			return
		}
	}
}
