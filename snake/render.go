package snake

type cell struct {
	symbol string
}

type matrix struct {
	cells [][]cell
}

const (
	title          = "Go Snake Telnet v0.1"
	score          = "Score: "
	horizontalLine = "-"
	verticalLine   = "|"
	emptySymbol    = " "
	snakeSymbol    = "x"
	foodSymbol     = "@"
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
		m.cells[b.x+3][b.y+1] = cell{
			symbol: snakeSymbol,
		}
	}
}

func (m *matrix) renderFood(x, y int, f *food) {
	// +3 - because title, empty line and horizontal line
	// +1 = because vertical line
	m.cells[x+3][y+1] = cell{
		symbol: foodSymbol,
	}
}

func (m *matrix) renderScore(a *arena, scoreVal int) {
	empty := []cell{}
	for i := 0; i < a.width; i++ {
		empty = append(empty, cell{
			symbol: emptySymbol,
		})
	}
	m.cells = append(m.cells, empty)

	scoreRow := []cell{}
	scoreString := score + string(scoreVal)
	for i, r := range scoreString {
		if i < a.width {
			scoreRow = append(scoreRow, cell{
				symbol: string(r),
			})
		}
	}
	m.cells = append(m.cells, scoreRow)
}

func (m *matrix) renderTitle(a *arena) {
	titleRow := []cell{}
	for i, r := range title {
		if i < a.width {
			titleRow = append(titleRow, cell{
				symbol: string(r),
			})
		}
	}
	m.cells = append(m.cells, titleRow)

	empty := []cell{}
	for i := 0; i < a.width; i++ {
		empty = append(empty, cell{
			symbol: emptySymbol,
		})
	}
	m.cells = append(m.cells, empty)
}
