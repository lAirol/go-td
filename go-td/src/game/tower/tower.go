package tower

import (
	"go-td/src/game/map/cell"
	"image/color"
	"time"
)

type Tower struct {
	Damage float32
	Price  float32
	Radius float32
	Reload time.Duration
	Cord   cell.Cord
	Color  color.Color
}

func CreateTower() Tower {
	return Tower{
		Damage: 1,
		Price:  10,
		Radius: 10,
		Reload: time.Second,
		Color:  cell.TowerColor,
	}
}
