package internal

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/z7ealth/gron.git/src/consts"
)

type Game struct {
	Disc         Disc
	Motorcycle   Motorcycle
	Running      bool
	Score        uint32
	DiscSound    rl.Sound
	CrashSound   rl.Sound
	IntroSound   rl.Sound
	RunningSound rl.Sound
	IntroPlayed  bool
}

func NewGame() Game {
	return Game{
		Disc:         NewDisc(),
		Motorcycle:   NewMotorcycle(),
		Running:      false,
		Score:        0,
		DiscSound:    LoadSound("motorcycle_gets_disc.wav"),
		CrashSound:   LoadSound("motorcycle_crash.wav"),
		IntroSound:   LoadSound("motorcycle_intro.wav"),
		RunningSound: LoadSound("motorcycle_running.wav"),
		IntroPlayed:  false,
	}
}

func (g *Game) ShouldUpdate() bool {
	currentTime := time.Now()
	elapsed := currentTime.Sub(g.Motorcycle.lastUpdate).Milliseconds()
	interval := consts.MOTORCYCLE_MOVEMENT_INTERVAL * time.Millisecond

	if elapsed >= interval.Milliseconds() {
		g.Motorcycle.lastUpdate = currentTime
		return true
	}

	return false
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
		rl.PlaySound(g.DiscSound)
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
	rl.ClearBackground(GetColor(consts.BACKGROUND_COLOR))

	score := fmt.Sprintf("Score: %v", g.Score)
	rl.DrawText(score, 12, 12, 12, rl.White)

	fps := fmt.Sprintf("FPS: %v", rl.GetFPS())
	rl.DrawText(fps, (consts.CELL_SIZE*consts.CELL_COUNT)-60, 12, 12, rl.White)

	g.Disc.Draw()
	g.Motorcycle.Draw()
}

func (g *Game) AdjustSoundVolumes() {
	rl.SetSoundVolume(g.DiscSound, 0.1)
	rl.SetSoundVolume(g.CrashSound, 0.1)
	rl.SetSoundVolume(g.IntroSound, 0.1)
	rl.SetSoundVolume(g.RunningSound, 0.1)
}

func (g *Game) Clean() {
	rl.UnloadTexture(g.Disc.texture)

	rl.UnloadSound(g.DiscSound)
	rl.UnloadSound(g.CrashSound)
	rl.UnloadSound(g.IntroSound)
	rl.UnloadSound(g.RunningSound)
	rl.CloseAudioDevice()

	rl.CloseWindow()
}

func (g *Game) gameOver() {
	g.Motorcycle.Reset()
	g.Disc.Position = GetUniquePos(&g.Motorcycle.Body)
	g.Score = 0
	g.Running = false

	StopSoundIfPlaying(g.IntroSound)
	StopSoundIfPlaying(g.RunningSound)

	g.IntroPlayed = false
	rl.PlaySound(g.CrashSound)
}
