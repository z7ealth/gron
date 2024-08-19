package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/z7ealth/gron.git/src/consts"
	"github.com/z7ealth/gron.git/src/internal"
)

func main() {
	rl.InitWindow(consts.CELL_SIZE*consts.CELL_COUNT, consts.CELL_SIZE*consts.CELL_COUNT, "GRON")
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()

	game := internal.NewGame()
	game.AdjustSoundVolumes()
	defer game.Clean()

	for !rl.WindowShouldClose() {

		if game.Running && !game.IntroPlayed {
			rl.PlaySound(game.IntroSound)
			game.IntroPlayed = true
		}

		if game.Running && !rl.IsSoundPlaying(game.IntroSound) && !rl.IsSoundPlaying(game.RunningSound) {
			rl.PlaySound(game.RunningSound)
		}

		if game.ShouldUpdate() {
			game.Update()
		}

		game.UpdateMotorcycleDirection()

		rl.BeginDrawing()

		game.Draw()

		rl.EndDrawing()
	}
}
