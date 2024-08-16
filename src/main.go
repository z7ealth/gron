package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/z7ealth/gron.git/src/consts"
	"github.com/z7ealth/gron.git/src/internal"
)

func main() {
  rl.InitWindow(consts.CELL_SIZE * consts.CELL_COUNT, consts.CELL_SIZE * consts.CELL_COUNT, "GRON")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(internal.GetColor(consts.BACKGROUND_COLOR))

    fps := fmt.Sprintf("FPS: %v", rl.GetFPS())
		rl.DrawText(fps, (consts.CELL_SIZE * consts.CELL_COUNT) - 60, 12, 12, rl.White)

		rl.EndDrawing()
	}
}
