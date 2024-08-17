package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/z7ealth/gron.git/src/consts"
)

type Disc struct {
	Position rl.Vector2
	texture  rl.Texture2D
}

func NewDisc() Disc {
	return Disc{
		Position: GetRandomPos(),
		texture:  LoadTexture("food.png"),
	}
}

func (d *Disc) Draw() {
	rl.DrawTexture(d.texture, int32(d.Position.X)*consts.CELL_SIZE, int32(d.Position.Y)*consts.CELL_SIZE, rl.White)
}
