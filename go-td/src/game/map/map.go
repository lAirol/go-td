package _map

import (
	"go-td/src/conf"
	"go-td/src/game/enemy"
	"go-td/src/game/map/cell"
	"go-td/src/game/tower"
	"image/color"
	"math/rand"
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

	for i := start; i <= end; i++ {
		m.Cells[iteration][i].Kind = cell.Path
		m.Cells[iteration][i].Color = cell.PathColor
		m.Path = append(m.Path, m.Cells[iteration][i])
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

func (m *Map) addEnemy(cell cell.Cell) {
	m.Enemies = append(m.Enemies, enemy.NewDefault(enemy.RCord{
		X: float32(cell.X * conf.GridSize),
		Y: float32(cell.X * conf.GridSize),
	}))
}
