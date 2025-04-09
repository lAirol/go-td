package missiles

import (
	"fmt"
	"go-td/src/conf"
	"go-td/src/game/map/cell"
	"go-td/src/game/map/enemy"
	"math"
)

type Missile struct {
	Speed  float64
	Size   cell.Cord
	Damage float64
	Cord   *enemy.RCord
	Target *enemy.Enemy
}

func (m *Missile) Launch() bool {
	dx := m.Target.RCord.X - m.Cord.X
	dy := m.Target.RCord.Y - m.Cord.Y

	distance := math.Sqrt(dx*dx + dy*dy)

	nx := dx / distance
	ny := dy / distance

	m.Cord.X += nx * m.Speed * conf.DeltaTime / 1000
	m.Cord.Y += ny * m.Speed * conf.DeltaTime / 1000
	if math.Abs(m.Cord.X-m.Target.X)*conf.GridSize < 1 && math.Abs(m.Cord.Y-m.Target.Y)*conf.GridSize < 1 {
		m.Target.Health -= m.Damage
		fmt.Println(m.Target.Health)
		return true
	}
	return false
}
