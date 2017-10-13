// Copyright (c) 2017 Alex Pliutau

package snake

import (
	"log"
	"time"
)

// Game type
type Game struct {
	KeyboardEventsChan chan KeyboardEvent
	PointsChan         chan int
	matrix             *matrix
	arena              *arena
	score              int
	IsOver             bool
}

// NewGame returns Game obj
func NewGame() *Game {
	g := Game{
		arena: initialArena(),
		score: initialScore(),
	}
	g.arena.buildCachedPartials()
	return &g
}

// Start game func
func (g *Game) Start() {
	g.KeyboardEventsChan = make(chan KeyboardEvent)
	g.PointsChan = make(chan int)
	g.arena.pointsChan = g.PointsChan

	for {
		select {
		case p := <-g.PointsChan:
			g.addPoints(p)
		case e := <-g.KeyboardEventsChan:
			d := keyToDirection(e.Key)
			if d > 0 {
				g.arena.snake.changeDirection(d)
			}
		default:
			if g.IsOver {
				log.Printf("Game over, score: %d\n", g.score)
				return
			}

			if err := g.arena.moveSnake(); err != nil {
				g.IsOver = true
			}

			time.Sleep(g.moveInterval())
		}
	}
}

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
	return newArena(initialSnake(), 20, 20)
}

func (g *Game) moveInterval() time.Duration {
	return time.Duration(600) * time.Millisecond
}

func (g *Game) addPoints(p int) {
	g.score += p
}
