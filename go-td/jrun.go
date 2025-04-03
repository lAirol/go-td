package main

//
//import (
//	"fmt"
//	"log"
//	"math"
//
//	"github.com/veandco/go-sdl2/sdl"
//)
//
//const (
//	WindowWidth  = 800
//	WindowHeight = 600
//	GridSize     = 40
//)
//
//type Game struct {
//	window   *sdl.Window
//	renderer *sdl.Renderer
//	towers   []Tower
//	enemies  []Enemy
//	grid     [][]int // 0 - пусто, 1 - путь, 2 - башня
//	money    int
//	wave     int
//	running  bool
//}
//
//type Tower struct {
//	X, Y   int
//	Range  float64
//	Damage int
//}
//
//type Enemy struct {
//	X, Y   float64
//	HP     int
//	Speed  float64
//	Path   []sdl.Point
//	PathID int
//}
//
//func NewGame() *Game {
//	// Инициализация SDL
//	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
//		log.Fatal("SDL init error:", err)
//	}
//
//	// Создание окна
//	window, err := sdl.CreateWindow(
//		"Pure Go Tower Defense",
//		sdl.WINDOWPOS_CENTERED,
//		sdl.WINDOWPOS_CENTERED,
//		WindowWidth,
//		WindowHeight,
//		sdl.WINDOW_SHOWN,
//	)
//	if err != nil {
//		log.Fatal("Window creation error:", err)
//	}
//
//	// Создание рендерера
//	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
//	if err != nil {
//		log.Fatal("Renderer creation error:", err)
//	}
//
//	// Инициализация сетки
//	grid := make([][]int, WindowWidth/GridSize)
//	for i := range grid {
//		grid[i] = make([]int, WindowHeight/GridSize)
//	}
//
//	// Задаем путь для врагов (1 - путь)
//	path := []sdl.Point{
//		{0, 1}, {1, 1}, {2, 1}, {2, 2}, {3, 2}, {3, 3}, {4, 3}, {4, 4}, {5, 4}, {5, 5},
//	}
//	for _, p := range path {
//		grid[p.X][p.Y] = 1
//	}
//
//	return &Game{
//		window:   window,
//		renderer: renderer,
//		grid:     grid,
//		money:    100,
//		wave:     1,
//		running:  true,
//	}
//}
//
//func (g *Game) Run() {
//	defer g.Cleanup()
//
//	for g.running {
//		g.HandleEvents()
//		g.Update()
//		g.Render()
//		sdl.Delay(16) // ~60 FPS
//	}
//}
//
//func (g *Game) HandleEvents() {
//	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
//		switch e := event.(type) {
//		case *sdl.QuitEvent:
//			g.running = false
//		case *sdl.MouseButtonEvent:
//			if e.Button == sdl.BUTTON_LEFT && e.State == sdl.PRESSED {
//				gx, gy := int(e.X)/GridSize, int(e.Y)/GridSize
//				if gx >= 0 && gy >= 0 && gx < len(g.grid) && gy < len(g.grid[0]) && g.grid[gx][gy] == 0 && g.money >= 50 {
//					g.towers = append(g.towers, Tower{X: gx, Y: gy, Range: 100, Damage: 10})
//					g.grid[gx][gy] = 2
//					g.money -= 50
//				}
//			}
//		}
//	}
//}
//
//func (g *Game) Update() {
//	// Спавн врагов (упрощенно)
//	if len(g.enemies) < 5+g.wave {
//		g.SpawnEnemy()
//	}
//
//	// Движение врагов
//	for i := 0; i < len(g.enemies); i++ {
//		e := &g.enemies[i]
//		if e.PathID >= len(e.Path) {
//			g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
//			i--
//			continue
//		}
//
//		target := e.Path[e.PathID]
//		targetX, targetY := float64(target.X*GridSize+GridSize/2), float64(target.Y*GridSize+GridSize/2)
//		dx, dy := targetX-e.X, targetY-e.Y
//		dist := math.Sqrt(dx*dx + dy*dy)
//
//		if dist < 2 {
//			e.PathID++
//		} else {
//			e.X += dx / dist * e.Speed
//			e.Y += dy / dist * e.Speed
//		}
//	}
//
//	// Стрельба башен
//	for _, tower := range g.towers {
//		towerX, towerY := float64(tower.X*GridSize+GridSize/2), float64(tower.Y*GridSize+GridSize/2)
//		for i := 0; i < len(g.enemies); i++ {
//			e := &g.enemies[i]
//			dx, dy := towerX-e.X, towerY-e.Y
//			dist := math.Sqrt(dx*dx + dy*dy)
//			if dist <= tower.Range {
//				e.HP -= tower.Damage
//				if e.HP <= 0 {
//					g.money += 10
//					g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
//					i--
//				}
//				break
//			}
//		}
//	}
//}
//
//func (g *Game) Render() {
//	g.renderer.SetDrawColor(0, 0, 0, 255)
//	g.renderer.Clear()
//
//	// Рисуем сетку
//	for x := 0; x < WindowWidth; x += GridSize {
//		for y := 0; y < WindowHeight; y += GridSize {
//			cell := g.grid[x/GridSize][y/GridSize]
//			if cell == 1 {
//				g.renderer.SetDrawColor(100, 100, 100, 255) // Путь
//			} else {
//				g.renderer.SetDrawColor(50, 50, 50, 255) // Пусто
//			}
//			g.renderer.FillRect(&sdl.Rect{int32(x), int32(y), GridSize, GridSize})
//		}
//	}
//
//	// Рисуем башни (синие квадраты)
//	g.renderer.SetDrawColor(0, 0, 255, 255)
//	for _, t := range g.towers {
//		g.renderer.FillRect(&sdl.Rect{int32(t.X * GridSize), int32(t.Y * GridSize), GridSize, GridSize})
//	}
//
//	// Рисуем врагов (красные круги)
//	g.renderer.SetDrawColor(255, 0, 0, 255)
//	for _, e := range g.enemies {
//		g.renderer.FillRect(&sdl.Rect{int32(e.X - GridSize/3), int32(e.Y - GridSize/3), GridSize * 2 / 3, GridSize * 2 / 3})
//	}
//
//	// UI: деньги и волна
//	g.renderer.SetDrawColor(255, 255, 255, 255)
//	g.renderer.Present()
//}
//
//func (g *Game) SpawnEnemy() {
//	path := []sdl.Point{
//		{0, 1}, {1, 1}, {2, 1}, {2, 2}, {3, 2}, {3, 3}, {4, 3}, {4, 4}, {5, 4}, {5, 5},
//	}
//	g.enemies = append(g.enemies, Enemy{
//		X:     float64(path[0].X*GridSize + GridSize/2),
//		Y:     float64(path[0].Y*GridSize + GridSize/2),
//		HP:    20 + g.wave*5,
//		Speed: 1,
//		Path:  path,
//	})
//}
//
//func (g *Game) Cleanup() {
//	g.renderer.Destroy()
//	g.window.Destroy()
//	sdl.Quit()
//}
//
//func main() {
//	game := NewGame()
//	game.Run()
//	fmt.Println("Game closed!")
//}
