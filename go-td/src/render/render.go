package render

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"go-td/src/conf"
	_map "go-td/src/game/map"
	"go-td/src/game/map/tower"
	"go-td/src/game/ui"
	"log"
	"math"
)

func Create() *sdl.Renderer {
	window, err := sdl.CreateWindow(
		"Pure Go Tower Defense",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		conf.WindowWidth,
		conf.WindowHeight,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		log.Fatalf("Failed to create window: %s", err)
		return nil
	}

	r, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalf("Failed to create renderer: %s", err)
		return nil
	}
	ttf.Init()
	return r
}

func Render(renderer *sdl.Renderer, grid *_map.Map, ui ui.Ui) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	renderMap(renderer, grid)
	renderMissiles(renderer, grid.Towers)
	renderEnemy(renderer, grid)
	renderRange(renderer, grid)
	renderUI(renderer, grid, ui)
	renderer.Present()
}

func renderMap(renderer *sdl.Renderer, grid *_map.Map) {
	for x := 0; x < conf.WindowWidth; x += conf.GridSize {
		for y := 0; y < conf.WindowHeight; y += conf.GridSize {
			kind := grid.Cells[x/conf.GridSize][y/conf.GridSize].Kind
			switch kind {
			case 0:
				renderer.SetDrawColor(255, 255, 255, 255)
			case 1:
				if int32(x)/conf.GridSize == grid.Start.X && int32(y)/conf.GridSize == grid.Start.Y {
					renderer.SetDrawColor(0, 255, 0, 255)
				} else if int32(x)/conf.GridSize == grid.End.X-1 && int32(y)/conf.GridSize == grid.End.Y {
					renderer.SetDrawColor(255, 255, 0, 255)
				} else {
					renderer.SetDrawColor(128, 128, 128, 255)
				}
			case 2:
				renderer.SetDrawColor(0, 0, 255, 255)
			case 3:
				renderer.SetDrawColor(255, 0, 0, 255)
			}
			renderer.FillRect(&sdl.Rect{int32(x), int32(y), conf.GridSize, conf.GridSize})
		}
	}
}

func renderRange(renderer *sdl.Renderer, grid *_map.Map) {
	for _, tower := range grid.Towers {
		renderer.SetDrawColor(0, 255, 0, 255)
		drawCircle(renderer, int32(tower.Cord.X*conf.GridSize)+conf.GridSize/2, int32(tower.Cord.Y*conf.GridSize)+conf.GridSize/2, int32(tower.Range))
	}
}

func renderEnemy(renderer *sdl.Renderer, grid *_map.Map) {
	for _, e := range grid.Enemies {
		centerX := int32(e.X*conf.GridSize) + conf.GridSize/2
		centerY := int32(e.Y*conf.GridSize) + conf.GridSize/2
		radius := conf.GridSize / 2

		renderer.SetDrawColor(255, 0, 0, 255) // Красный цвет для круга
		drawFilledCircle(renderer, centerX, centerY, int32(radius-conf.GridSize/4))
	}
}

func renderUI(renderer *sdl.Renderer, grid *_map.Map, ui ui.Ui) {
	// Открываем шрифт
	font, err := ttf.OpenFont("src/conf/Movistar Text Regular.ttf", 16)
	if err != nil {
		panic(err)
	}
	defer font.Close()

	uiTexts := []struct {
		Text string
		X, Y int32
	}{
		{"Enemies: " + fmt.Sprint(len(grid.Enemies)), 10, 10},
		{"Towers: " + fmt.Sprint(len(grid.Towers)), 10, 30},
		{"Money: " + fmt.Sprint(ui.Money), 10, 50},
		{"Wave: " + fmt.Sprint(ui.Wave), 10, 70},
		{"HP: " + fmt.Sprint(ui.HP), 10, 90},
	}

	for _, uiText := range uiTexts {
		renderUIText(font, renderer, uiText.Text, uiText.X, uiText.Y)
	}
}

func renderUIText(font *ttf.Font, renderer *sdl.Renderer, text string, x, y int32) {
	surface, err := font.RenderUTF8Solid(text, sdl.Color{R: 0, G: 0, B: 0, A: 255})
	if err != nil {
		panic(err)
	}
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	rect := &sdl.Rect{
		X: x,
		Y: y,
		W: surface.W,
		H: surface.H,
	}
	renderer.Copy(texture, nil, rect)
}

func drawFilledCircle(renderer *sdl.Renderer, xCenter, yCenter, radius int32) {
	for y := int32(0); y <= radius; y++ {
		xOffset := int32(math.Sqrt(float64(radius*radius - y*y)))
		for x := -xOffset; x <= xOffset; x++ {
			renderer.DrawPoint(xCenter+x, yCenter+y)
			renderer.DrawPoint(xCenter+x, yCenter-y)
		}
	}
}

func drawCircle(renderer *sdl.Renderer, xCenter, yCenter, radius int32) {
	x := radius
	y := int32(0)
	err := int32(1 - x)

	for x >= y {
		drawSymmetricPoints(renderer, xCenter, yCenter, x, y)

		y++
		if err <= 0 {
			err += 2*y + 1
		} else {
			x--
			err += 2*(y-x) + 1
		}
	}
}

func drawSymmetricPoints(renderer *sdl.Renderer, xCenter, yCenter, x, y int32) {
	renderer.DrawPoint(xCenter+x, yCenter+y)
	renderer.DrawPoint(xCenter-x, yCenter+y)
	renderer.DrawPoint(xCenter+x, yCenter-y)
	renderer.DrawPoint(xCenter-x, yCenter-y)
	renderer.DrawPoint(xCenter+y, yCenter+x)
	renderer.DrawPoint(xCenter-y, yCenter+x)
	renderer.DrawPoint(xCenter+y, yCenter-x)
	renderer.DrawPoint(xCenter-y, yCenter-x)
}

func renderMissiles(renderer *sdl.Renderer, towers []tower.Tower) {
	renderer.SetDrawColor(0, 255, 255, 255)
	for _, t := range towers {
		for _, mis := range t.Missiles {
			renderer.FillRect(&sdl.Rect{int32(mis.Cord.X*conf.GridSize) + conf.GridSize/2, int32(mis.Cord.Y*conf.GridSize) + conf.GridSize/2, mis.Size.X * conf.MissileSize, mis.Size.Y * conf.MissileSize})
		}

	}

}
