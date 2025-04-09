package game

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-td/src/conf"
	_map "go-td/src/game/map"
	"go-td/src/game/map/cell"
	"go-td/src/game/map/tower"
	"go-td/src/game/ui"
	"go-td/src/render"
	"time"
)

type Game struct {
	Map       *_map.Map
	UI        ui.Ui
	Renderer  *sdl.Renderer
	running   bool
	DeltaTime time.Duration
}

func (g *Game) prepareMap() {
	g.Map = &_map.Map{}
	g.Map.New()
	g.UI = ui.Ui{
		Wave:  1,
		Money: 100,
	}
}

func (g *Game) Start() {
	defer g.Cleanup()
	g.running = true
	g.prepareMap()
	g.Renderer = render.Create()
	g.DeltaTime = time.Duration(16) * time.Millisecond
	for g.running {
		g.Update()
		g.handleEvents()
		render.Render(g.Renderer, g.Map, g.UI)
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
				if g.Map.Cells[cord.X][cord.Y].Kind == 0 {
					g.Map.SetTower(cord, tower.CreateTower())
				}
			}
		}
	}
}

func (g *Game) Update() {
	g.Map.Update(g.DeltaTime)
}
