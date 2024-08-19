package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/z7ealth/gron.git/src/consts"
)

type Game struct {
	Disc       Disc
	Motorcycle Motorcycle
	Running    bool
	Score      uint32
}

func NewGame() Game {
	disc := NewDisc()
	motorcycle := NewMotorcycle()

	return Game{
		Disc:       disc,
		Motorcycle: motorcycle,
		Running:    false,
		Score:      0,
	}
}

func (g *Game) Update() {
	if g.Running {
		g.Motorcycle.Update()
		g.checkCollissionWithDisc()
		g.checkCollissionWithEdges()
		g.checkCollissionWithTail()
	}
}

func (g *Game) UpdateMotorcycleDirection() {
	if rl.IsKeyDown(rl.KeyUp) && g.Motorcycle.Direction.Y != 1 || !(!rl.IsKeyDown(rl.KeyUp) || g.Running) {
		g.Motorcycle.Direction = rl.NewVector2(0, -1)
		g.Running = true
		return
	}
	if rl.IsKeyDown(rl.KeyDown) && g.Motorcycle.Direction.Y != -1 && g.Running {
		g.Motorcycle.Direction = rl.NewVector2(0, 1)
		g.Running = true
		return
	}
	if rl.IsKeyDown(rl.KeyLeft) && g.Motorcycle.Direction.X != 1 {
		g.Motorcycle.Direction = rl.NewVector2(-1, 0)
		g.Running = true
		return
	}
	if rl.IsKeyDown(rl.KeyRight) && g.Motorcycle.Direction.X != -1 {
		g.Motorcycle.Direction = rl.NewVector2(1, 0)
		g.Running = true
	}
}

func (g *Game) checkCollissionWithDisc() {
	if rl.Vector2Equals(g.Motorcycle.Body.At(0), g.Disc.Position) {
		g.Disc.Position = GetUniquePos(&g.Motorcycle.Body)
		g.Motorcycle.shouldIncrease = true
		g.Score += 10
	}
}

func (g *Game) checkCollissionWithEdges() {
	if g.Motorcycle.Body.At(0).X == consts.CELL_COUNT || g.Motorcycle.Body.At(0).X == -1 {
		g.gameOver()
	}
	if g.Motorcycle.Body.At(0).Y == consts.CELL_COUNT || g.Motorcycle.Body.At(0).Y == -1 {
		g.gameOver()
	}
}

func (g *Game) checkCollissionWithTail() {
	for i := 1; i < g.Motorcycle.Body.Len(); i++ {
		if rl.Vector2Equals(g.Motorcycle.Body.At(0), g.Motorcycle.Body.At(i)) {
			g.gameOver()
		}
	}
}

func (g *Game) Draw() {
	g.Disc.Draw()
	g.Motorcycle.Draw()
}

func (g *Game) Clean() {
	rl.UnloadTexture(g.Disc.texture)
}

func (g *Game) gameOver() {
	g.Motorcycle.Reset()
	g.Disc.Position = GetUniquePos(&g.Motorcycle.Body)
	g.Score = 0
	g.Running = false
}
