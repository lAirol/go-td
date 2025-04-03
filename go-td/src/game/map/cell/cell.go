package cell

import "image/color"

type Kind int

var EmptyColor = color.RGBA{R: 255, G: 255, B: 255, A: 255} // Белый
var PathColor = color.RGBA{R: 128, G: 128, B: 128, A: 255}  // Серый
var TowerColor = color.RGBA{R: 0, G: 0, B: 255, A: 255}     // Синий
var EnemyColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}     // Красный

const (
	Empty Kind = iota
	Path
	Tower
	Enemy
)

var EmptyCell = Cell{
	Color: EmptyColor,
	Kind:  Empty,
}
var PathCell = Cell{
	Color: PathColor,
	Kind:  Path,
}

var TowerCell = Cell{
	Color: TowerColor,
	Kind:  Tower,
}

var EnemyCell = Cell{
	Color: EnemyColor,
	Kind:  Enemy,
}

type Cell struct {
	Color color.Color
	Kind
	Cord
}

type Cord struct {
	X int32
	Y int32
}
