package snake

type food struct {
	points, x, y int
}

func newFood(x, y int) *food {
	return &food{
		points: 10,
		x:      x,
		y:      y,
	}
}
