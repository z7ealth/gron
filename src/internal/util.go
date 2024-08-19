package internal

import (
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gammazero/deque"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/z7ealth/gron.git/src/consts"
)

// Colors

func handleColorParseErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetColor(color string) rl.Color {
	splitted := strings.Split(color, ", ")

	r, err := strconv.Atoi(splitted[0])
	handleColorParseErr(err)
	g, err := strconv.Atoi(splitted[1])
	handleColorParseErr(err)
	b, err := strconv.Atoi(splitted[2])
	handleColorParseErr(err)
	a, err := strconv.Atoi(splitted[3])
	handleColorParseErr(err)

	return rl.NewColor(uint8(r), uint8(g), uint8(b), uint8(a))
}

// Positions

func GetRandomPos() rl.Vector2 {
	x := float32(rl.GetRandomValue(0, consts.CELL_COUNT-1))
	y := float32(rl.GetRandomValue(0, consts.CELL_COUNT-1))
	return rl.NewVector2(x, y)
}

func GetUniquePos(invalid *deque.Deque[rl.Vector2]) rl.Vector2 {
	position := GetRandomPos()
	var unique bool
	for !unique {
		unique = true
		for i := 0; i < invalid.Len(); i++ {
			if rl.Vector2Equals(invalid.At(i), position) {
				position = GetRandomPos()
				unique = false
			}
		}
	}
	return position
}

// Textures

func LoadTexture(imageName string) rl.Texture2D {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Unable to get texture path")
	}

	fileName := path.Join(cwd, "assets/graphics/objects", imageName)

	image := rl.LoadImage(fileName)
	defer rl.UnloadImage(image)

	return rl.LoadTextureFromImage(image)
}
