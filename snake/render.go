package snake

type cell struct {
	symbol string
}

type matrix struct {
	cells [][]cell
}

const (
	title          = "Go Snake Telnet v0.1"
	move           = "Move:"
	usage          = "w,d,s,a <Enter>"
	score          = "Score: "
	input          = "Your input: "
	horizontalLine = "-"
	verticalLine   = "|"
	emptySymbol    = " "
	snakeSymbol    = "x"
	foodSymbol     = "@"
	fieldTop       = 6
	fieldLeft      = 1
)

// Render returns game arena as string
func (g *Game) Render() string {
	ascii := ""

	m := g.genMatrix()
	for _, row := range m.cells {
		for _, cell := range row {
			ascii += cell.symbol
		}
		ascii += "\n"
	}

	return ascii
}

func (g *Game) genMatrix() *matrix {
	m := new(matrix)
	m.renderTitle(g.arena)
	m.renderArena(g.arena)
	m.renderSnake(g.arena.snake)
	m.renderFood(10, 10, nil)
	m.renderScore(g.arena, g.score)
	return m
}

func (m *matrix) renderArena(a *arena) {
	horizontal := []cell{}
	horizontal = append(horizontal, cell{
		symbol: verticalLine,
	})
	for i := 0; i < a.width; i++ {
		horizontal = append(horizontal, cell{
			symbol: horizontalLine,
		})
	}
	horizontal = append(horizontal, cell{
		symbol: verticalLine,
	})

	m.cells = append(m.cells, horizontal)
	for i := 0; i < a.height; i++ {
		row := []cell{cell{
			symbol: verticalLine,
		}}
		for i := 0; i < a.width; i++ {
			row = append(row, cell{
				symbol: emptySymbol,
			})
		}
		row = append(row, cell{
			symbol: verticalLine,
		})
		m.cells = append(m.cells, row)
	}

	m.cells = append(m.cells, horizontal)
}

func (m *matrix) renderSnake(s *snake) {
	for _, b := range s.body {
		m.cells[b.x+fieldTop][b.y+fieldLeft] = cell{
			symbol: snakeSymbol,
		}
	}
}

func (m *matrix) renderFood(x, y int, f *food) {
	m.cells[x+fieldTop][y+fieldLeft] = cell{
		symbol: foodSymbol,
	}
}

func (m *matrix) renderScore(a *arena, scoreVal int) {
	m.addEmptyRow(a)
	m.renderString(score + string(scoreVal))
	m.addEmptyRow(a)
	m.renderString(input)
}

func (m *matrix) renderTitle(a *arena) {
	m.renderString(title)
	m.addEmptyRow(a)
	m.renderString(move)
	m.renderString(usage)
	m.addEmptyRow(a)
}

func (m *matrix) addEmptyRow(a *arena) {
	empty := []cell{}
	for i := 0; i < a.width; i++ {
		empty = append(empty, cell{
			symbol: emptySymbol,
		})
	}
	m.cells = append(m.cells, empty)
}

func (m *matrix) renderString(s string) {
	row := []cell{}
	for _, r := range s {
		row = append(row, cell{
			symbol: string(r),
		})
	}
	m.cells = append(m.cells, row)
}
