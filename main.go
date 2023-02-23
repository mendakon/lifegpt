package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kuchibashi/lifegpt/back"
	"github.com/kuchibashi/lifegpt/front"
)

var state = back.NewState(front.ScreenWidth, front.ScreenHeight)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	state = state.Update()
	front.Draw(screen, state)

	return nil
}

func main() {
	state.Randomize()
	ebiten.Run(update, front.ScreenWidth, front.ScreenHeight, 1.0, "Game of Life")
}
