package internal

import (
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
