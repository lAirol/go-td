package render

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-td/src/conf"
	_map "go-td/src/game/map"
	"log"
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
	return r
}

func Render(renderer *sdl.Renderer, grid *_map.Map) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	renderMap(renderer, grid)

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
