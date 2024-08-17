package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
  Disc Disc
}

func NewGame() Game {

  disc := NewDisc()

  return Game{
    Disc: disc,
  }
}

func (g *Game) Draw() {
  g.Disc.Draw()
}

func (g *Game) Clean() {
  rl.UnloadTexture(g.Disc.texture)
}
