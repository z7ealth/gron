package internal

import (
	"time"

	"github.com/gammazero/deque"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/z7ealth/gron.git/src/consts"
)

type Motorcycle struct {
	Body           deque.Deque[rl.Vector2]
	Direction      rl.Vector2
	lastUpdate     time.Time
	shouldIncrease bool
}

func NewMotorcycle() Motorcycle {
	return Motorcycle{
		Body:           getInitialPos(),
		Direction:      getInitialDirection(),
		lastUpdate:     time.Now(),
		shouldIncrease: false,
	}
}

func getInitialDirection() rl.Vector2 {
	return rl.NewVector2(0, 1)
}

func getInitialPos() deque.Deque[rl.Vector2] {
	initialPos := deque.New[rl.Vector2]()
	initialPos.PushBack(rl.NewVector2(consts.CELL_COUNT/2, consts.CELL_COUNT-4))
	initialPos.PushBack(rl.NewVector2(consts.CELL_COUNT/2, consts.CELL_COUNT-3))
	initialPos.PushBack(rl.NewVector2(consts.CELL_COUNT/2, consts.CELL_COUNT-2))
	return *initialPos
}

func (m *Motorcycle) Update() {
	m.Body.PushFront(rl.Vector2Add(m.Body.At(0), m.Direction))
	if m.shouldIncrease {
		m.shouldIncrease = false
	} else {
		m.Body.PopBack()
	}
}

func (m *Motorcycle) Draw() {
	var rectangle rl.Rectangle
	for i := 0; i < m.Body.Len(); i++ {
		if i == 0 {
			rectangle = rl.NewRectangle(m.Body.At(0).X*consts.CELL_SIZE, m.Body.At(0).Y*consts.CELL_SIZE, consts.CELL_SIZE, consts.CELL_SIZE)
			rl.DrawRectangleRounded(rectangle, 0.6, 6, rl.SkyBlue)
			continue
		}
		rectangle = rl.NewRectangle(m.Body.At(i).X*consts.CELL_SIZE, m.Body.At(i).Y*consts.CELL_SIZE, consts.CELL_SIZE, consts.CELL_SIZE)
		rl.DrawRectangleRounded(rectangle, 0.6, 6, rl.DarkGray)
	}
}

func (m *Motorcycle) Reset() {
	m.Direction = getInitialDirection()
	m.Body = getInitialPos()
}
