package tower

import (
	"go-td/src/conf"
	"go-td/src/game/map/cell"
	"go-td/src/game/map/enemy"
	"go-td/src/game/map/tower/missiles"
	"image/color"
	"math"
	"time"
)

type Tower struct {
	Damage   float64
	Price    float64
	Range    float64
	Reload   time.Duration
	Cord     cell.Cord
	Color    color.Color
	CoolDown float64
	Missiles []missiles.Missile
}

func CreateTower() Tower {
	return Tower{
		Damage:   5,
		Price:    10,
		Range:    130,
		Reload:   time.Second,
		Color:    cell.TowerColor,
		CoolDown: 0,
		Missiles: make([]missiles.Missile, 0),
	}
}

func (t *Tower) Update(e []enemy.Enemy, deltaTime time.Duration) {
	if t.CoolDown > 0 {
		t.CoolDown -= deltaTime.Seconds()
	}

	if t.CoolDown <= 0 {
		t.shoot(e)
	}
}

func (t *Tower) shoot(ens []enemy.Enemy) {
	for i := 0; i < len(ens); i++ {
		dx, dy := float64(t.Cord.X)-ens[i].X, float64(t.Cord.Y)-ens[i].Y
		dist := math.Sqrt(dx*dx+dy*dy) * conf.GridSize
		if dist <= t.Range && t.CoolDown <= 0 {
			missile := missiles.Missile{
				Speed:  3,
				Size:   cell.Cord{X: 1, Y: 1},
				Damage: t.Damage,
				Cord:   &enemy.RCord{X: float64(t.Cord.X), Y: float64(t.Cord.Y)},
				Target: &ens[i],
			}
			t.Missiles = append(t.Missiles, missile)
			t.CoolDown = t.Reload.Seconds()
		}
	}
}
