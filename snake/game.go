// Copyright (c) 2017 Alex Pliutau

package snake

import (
	"log"
	"math"
	"time"
)

var (
	topScoreChan chan int
	topScoreVal  int
)

func init() {
	topScoreChan = make(chan int)
	go func() {
		for {
			s := <-topScoreChan
			if s > topScoreVal {
				topScoreVal = s
			}
		}
	}()
}

// Game type
type Game struct {
	KeyboardEventsChan chan KeyboardEvent
	PointsChan         chan int
	arena              *arena
	score              int
	IsOver             bool
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
				topScoreChan <- g.score
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
	ms := 400 - math.Max(float64(g.score), 100)
	return time.Duration(ms) * time.Millisecond
}

func (g *Game) addPoints(p int) {
	g.score += p
}
