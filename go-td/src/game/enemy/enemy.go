package enemy

import (
	"go-td/src/game/map/cell"
	"image/color"
)

type Enemy struct {
	Health float32
	Speed  float32
	Color  color.Color
	RCord
}

type RCord struct {
	X, Y float32
}

func NewDefault(r RCord) Enemy {
	return Enemy{
		Health: 10,
		Speed:  0.1,
		Color:  cell.EnemyColor,
		RCord:  r,
	}
}

func (e *Enemy) move() {
	
}
