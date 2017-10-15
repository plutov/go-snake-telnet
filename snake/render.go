// Copyright (c) 2017 Alex Pliutau

package snake

import (
	"fmt"
	"strings"
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
	topScore       = "Top score: "
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
		ascii += strings.Join(row, "") + "\n"
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
	m.cells = append(m.cells, a.horizontalRow)

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
		} else {

			row := []string{verticalLine}
			for i := 0; i < a.width; i++ {
				row = append(row, emptySymbol)
			}
			row = append(row, verticalLine)
			m.cells = append(m.cells, row)
			//m.cells = append(m.cells, a.arenaRow)
		}
	}

	// Add horizontal line on bottom
	m.cells = append(m.cells, a.horizontalRow)
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
	m.renderString(fmt.Sprintf("%s%d", topScore, topScoreVal))
	m.addEmptyRow(a)
	m.cells = append(m.cells, a.inputRow)
}

func (m *matrix) renderTitle(a *arena) {
	m.cells = append(m.cells, a.titleRow)
	m.cells = append(m.cells, a.authorRow)
	m.addEmptyRow(a)
	m.cells = append(m.cells, a.moveRow)
	m.cells = append(m.cells, a.usageRow)
	m.addEmptyRow(a)
}

func (m *matrix) addEmptyRow(a *arena) {
	m.cells = append(m.cells, a.emptyRow)
}

func (m *matrix) renderString(s string) {
	row := renderString(s)
	m.cells = append(m.cells, row)
}

func (a *arena) buildCachedPartials() {
	a.emptyRow = []string{}
	a.horizontalRow = strings.Split(verticalLine+strings.Repeat(horizontalLine, a.width)+verticalLine, "")
	a.arenaRow = strings.Split(verticalLine+strings.Repeat(emptySymbol, a.width)+verticalLine, "")
	a.titleRow = renderString(title)
	a.authorRow = renderString(author)
	a.usageRow = renderString(usage)
	a.moveRow = renderString(move)
	a.inputRow = renderString(input)
}

func renderString(s string) []string {
	return strings.Split(s, "")
}
