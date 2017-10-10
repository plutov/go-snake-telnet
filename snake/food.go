package snake

type food struct {
	emoji        rune
	points, x, y int
}

func newFood(x, y int) *food {
	return &food{
		points: 10,
		emoji:  '@',
		x:      x,
		y:      y,
	}
}
