// Copyright (c) 2017 Alex Pliutau

package snake

import (
	"fmt"
)

type matrix struct {
	cells [][]string
}

const (
	title          = "Go Snake Telnet v0.2"
	author         = "Author: pliutau.com"
	move           = "Move:"
	usage          = "W,D,S,A <Enter>"
	score          = "Score: "
	input          = "Your input: "
	horizontalLine = "-"
	verticalLine   = "|"
	emptySymbol    = " "
	snakeSymbol    = "*"
	foodSymbol     = "@"
	gameOver       = "Game over!"
	fieldTop       = 7
	fieldLeft      = 1
)

// Render returns game arena as string
func (g *Game) Render() string {
	ascii := ""

	m := g.genMatrix()
	for _, row := range m.cells {
		for _, cell := range row {
			ascii += cell
		}
		ascii += "\n"
	}

	return ascii
}

func (g *Game) genMatrix() *matrix {
	m := new(matrix)
	m.renderTitle(g.arena)
	m.renderArena(g.arena, g)
	if !g.IsOver {
		m.renderFood(g.arena.food.x, g.arena.food.y)
		m.renderSnake(g.arena.snake)
	}

	m.renderScore(g.arena, g.score)
	return m
}

func (m *matrix) renderArena(a *arena, g *Game) {
	// Add horizontal line on top
	horizontal := []string{}
	horizontal = append(horizontal, verticalLine)
	for i := 0; i < a.width; i++ {
		horizontal = append(horizontal, horizontalLine)
	}
	horizontal = append(horizontal, verticalLine)
	m.cells = append(m.cells, horizontal)

	// Render battlefield
	for i := 0; i < a.height; i++ {
		if i == 1 && g.IsOver {
			row := []string{verticalLine, emptySymbol}
			for _, r := range gameOver {
				row = append(row, string(r))
			}
			for j := len(gameOver) + 1; j < a.width; j++ {
				row = append(row, emptySymbol)
			}
			row = append(row, verticalLine)
			m.cells = append(m.cells, row)
			continue
		}

		row := []string{verticalLine}
		for i := 0; i < a.width; i++ {
			row = append(row, emptySymbol)
		}
		row = append(row, verticalLine)
		m.cells = append(m.cells, row)
	}

	// Add horizontal line on bottom
	m.cells = append(m.cells, horizontal)
}

func (m *matrix) renderSnake(s *snake) {
	for _, b := range s.body {
		m.cells[b.x+fieldTop][b.y+fieldLeft] = snakeSymbol
	}
}

func (m *matrix) renderFood(x, y int) {
	m.cells[x+fieldTop][y+fieldLeft] = foodSymbol
}

func (m *matrix) renderScore(a *arena, scoreVal int) {
	m.addEmptyRow(a)
	m.renderString(fmt.Sprintf("%s%d", score, scoreVal))
	m.addEmptyRow(a)
	m.renderString(input)
}

func (m *matrix) renderTitle(a *arena) {
	m.renderString(title)
	m.renderString(author)
	m.addEmptyRow(a)
	m.renderString(move)
	m.renderString(usage)
	m.addEmptyRow(a)
}

func (m *matrix) addEmptyRow(a *arena) {
	empty := []string{}
	for i := 0; i < a.width; i++ {
		empty = append(empty, emptySymbol)
	}
	m.cells = append(m.cells, empty)
}

func (m *matrix) renderString(s string) {
	row := []string{}
	for _, r := range s {
		row = append(row, string(r))
	}
	m.cells = append(m.cells, row)
}
