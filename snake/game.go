// Copyright (c) 2017 Alex Pliutau

package snake

import (
	"time"
)

// Game type
type Game struct {
	KeyboardEventsChan chan KeyboardEvent
	matrix             *matrix
	arena              *arena
	score              int
	isOver             bool
}

// NewGame returns Game obj
func NewGame() *Game {
	return &Game{
		arena: initialArena(),
		score: initialScore(),
	}
}

// Start game func
func (g *Game) Start() {
	g.KeyboardEventsChan = make(chan KeyboardEvent)

	for {
		select {
		case p := <-pointsChan:
			g.addPoints(p)
		case e := <-g.KeyboardEventsChan:
			d := keyToDirection(e.Key)
			g.arena.snake.changeDirection(d)
		default:
			if !g.isOver {
				if err := g.arena.moveSnake(); err != nil {
					g.end()
				}
			}

			time.Sleep(g.moveInterval())
		}
	}
}

var (
	pointsChan = make(chan int)
	// KeyboardEventsChan - keyboard input will go here

)

func initialSnake() *snake {
	return newSnake(RIGHT, []coord{
		coord{x: 1, y: 1},
		coord{x: 1, y: 2},
		coord{x: 1, y: 3},
		coord{x: 1, y: 4},
	})
}

func initialScore() int {
	return 0
}

func initialArena() *arena {
	return newArena(initialSnake(), pointsChan, 20, 20)
}

func (g *Game) end() {
	g.isOver = true
}

func (g *Game) moveInterval() time.Duration {
	return time.Duration(1) * time.Second
}

func (g *Game) retry() {
	g.arena = initialArena()
	g.score = initialScore()
	g.isOver = false
}

func (g *Game) addPoints(p int) {
	g.score += p
}
