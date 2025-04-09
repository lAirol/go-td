package enemy

import (
	"go-td/src/game/map/cell"
	"image/color"
	"math"
)

type Enemy struct {
	Health float64
	Speed  float64
	Color  color.Color
	RCord
	Path     []cell.Cell
	nextPath int
}

const epsilon = 0.0001

type RCord struct {
	X, Y float64
}

func NewDefault(r RCord, path []cell.Cell) Enemy {
	return Enemy{
		Health:   10,
		Speed:    0.025,
		Color:    cell.EnemyColor,
		RCord:    r,
		Path:     path,
		nextPath: 1,
	}
}

func (e *Enemy) Update() {
	e.move()
}

func (e *Enemy) move() {
	if len(e.Path) > e.nextPath {
		if math.Abs(e.X-float64(e.Path[e.nextPath].X)) <= epsilon && math.Abs(e.Y-float64(e.Path[e.nextPath].Y)) <= epsilon {
			e.X = float64(e.Path[e.nextPath].X)
			e.Y = float64(e.Path[e.nextPath].Y)
			e.nextPath++
		} else {
			if math.Abs(e.X-float64(e.Path[e.nextPath].X)) > epsilon {
				if float64(e.Path[e.nextPath].X) > e.X {
					e.X += e.Speed
				} else {
					e.X -= e.Speed
				}
			}
			if math.Abs(e.Y-float64(e.Path[e.nextPath].Y)) > epsilon {
				if float64(e.Path[e.nextPath].Y) > e.Y {
					e.Y += e.Speed
				} else {
					e.Y -= e.Speed
				}
			}
		}
	} else {
		e.Health = 0
	}
}
