package game

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-td/src/conf"
	_map "go-td/src/game/map"
	"go-td/src/game/map/cell"
	"go-td/src/game/tower"
	"go-td/src/render"
)

type Game struct {
	Map      *_map.Map
	Renderer *sdl.Renderer
	running  bool
}

func (g *Game) prepareMap() {
	g.Map = &_map.Map{}
	g.Map.New()
}

func (g *Game) Start() {
	defer g.Cleanup()
	g.running = true
	g.prepareMap()
	g.Renderer = render.Create()
	for g.running {
		g.handleEvents()
		render.Render(g.Renderer, g.Map)
		sdl.Delay(16)
	}
}

func (g *Game) Cleanup() {
	g.Renderer.Destroy()
	sdl.Quit()
}

func (g *Game) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			g.running = false
		case *sdl.MouseButtonEvent:
			if e.Button == sdl.BUTTON_LEFT && e.State == sdl.PRESSED {
				cord := cell.Cord{X: e.X / conf.GridSize, Y: e.Y / conf.GridSize}
				g.Map.SetTower(cord, tower.CreateTower())
			}
		}
	}
}
