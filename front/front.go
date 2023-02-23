package front

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/kuchibashi/lifegpt/back"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 800
	CellSize     = 1
)

var (
	aliveColor = color.NRGBA{0xff, 0xff, 0xff, 0xff}
	deadColor  = color.NRGBA{0x00, 0x00, 0x00, 0xff}
)

func Draw(screen *ebiten.Image, state *back.State) {
	screen.Fill(deadColor)

	for x := 0; x < state.Width; x++ {
		for y := 0; y < state.Height; y++ {
			if state.Get(x, y) {
				screen.Set(x*CellSize, y*CellSize, aliveColor)
			}
		}
	}
}
