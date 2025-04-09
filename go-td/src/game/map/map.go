package _map

import (
	"go-td/src/conf"
	"go-td/src/game/map/cell"
	"go-td/src/game/map/enemy"
	"go-td/src/game/map/tower"
	"go-td/src/game/map/tower/missiles"
	"image/color"
	"math/rand"
	"time"
)

type Map struct {
	Towers  []tower.Tower
	Enemies []enemy.Enemy
	Cells   [][]cell.Cell
	Path    []cell.Cell
	Start   cell.Cord
	End     cell.Cord
}

func (m *Map) New() {
	m.Towers = make([]tower.Tower, 0)
	m.Enemies = make([]enemy.Enemy, 0)
	m.generateClearMapWithPath(conf.MapXSize, conf.MapYSize)
	m.generateEnemies()
}

func (m *Map) generateClearMap(x, y int) {
	cells := make([][]cell.Cell, x)
	for i := 0; i < x; i++ {
		cells[i] = make([]cell.Cell, y)
		for j := 0; j < y; j++ {
			cells[i][j] = cell.Cell{Color: color.Black, Kind: cell.Empty, Cord: cell.Cord{X: int32(i), Y: int32(j)}}
		}
	}
	m.Cells = cells
}

func (m *Map) generateClearMapWithPath(x, y int) {
	m.generateClearMap(x, y)
	start := m.startPoint()
	m.Start = cell.Cord{X: 0, Y: int32(start)}

	for i := 0; i < conf.MapXSize; i++ {
		start = m.generatePath(start, i)
	}
	m.End = cell.Cord{X: conf.MapXSize, Y: int32(start)}
}

func (m *Map) generateEnemies() {
	var count = 10
	for i := 0; i < count; i++ {
		m.addEnemy()
	}
}

func (m *Map) generatePath(start int, iteration int) int {
	distance := m.genDistance()
	direction := m.genDirection()

	end := start + direction*distance
	if end < 0 {
		end = 0
	} else if end >= conf.MapYSize {
		end = conf.MapYSize - 1
	}
	res := end
	if start > end {
		start, end = end, start

	}

	var tmp []cell.Cell
	for i := start; i <= end; i++ {
		m.Cells[iteration][i].Kind = cell.Path
		m.Cells[iteration][i].Color = cell.PathColor
		tmp = append(tmp, m.Cells[iteration][i])
	}
	if res == end {
		m.Path = append(m.Path, tmp...)
	} else {
		for i := len(tmp) - 1; i >= 0; i-- {
			m.Path = append(m.Path, tmp[i])
		}
	}
	return res
}

func (m *Map) genDistance() int {
	return rand.Intn(2) + 1
}

func (m *Map) genDirection() int {
	dir := rand.Intn(3) - 1
	return dir
}

func (m *Map) startPoint() int {
	return rand.Intn(conf.MapYSize)
}

func (m *Map) updateCell(cord cell.Cord, cell cell.Cell) {
	cell.Cord = cord
	m.Cells[cord.X][cord.Y] = cell
}

func (m *Map) SetTower(cord cell.Cord, tower tower.Tower) {
	tower.Cord = cord
	m.Towers = append(m.Towers, tower)
	m.updateCell(cord, cell.TowerCell)
}

func (m *Map) addEnemy() {
	m.Enemies = append(m.Enemies, enemy.NewDefault(enemy.RCord{
		X: float64(m.Start.X),
		Y: float64(m.Start.Y),
	}, m.Path))
}

func (m *Map) Update(delta time.Duration) {
	m.updateTowers(delta)
	m.updateEnemies()
	m.destroyEnemies()
}

func (m *Map) updateTowers(delta time.Duration) {
	for t := range m.Towers {
		m.Towers[t].Update(m.Enemies, delta)
		newMissiles := make([]missiles.Missile, 0)
		for _, mi := range m.Towers[t].Missiles {
			if !mi.Launch() {
				newMissiles = append(newMissiles, mi)
			}
		}
		m.Towers[t].Missiles = newMissiles
	}
}

func (m *Map) updateEnemies() {
	for e := range m.Enemies {
		m.Enemies[e].Update()
	}
}

func (m *Map) destroyEnemies() {
	pos := 0
	for _, e := range m.Enemies {
		if e.Health > 0 {
			m.Enemies[pos] = e
			pos++
		}
	}
	m.Enemies = m.Enemies[:pos]
}
