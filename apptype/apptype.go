package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrustType = int

type PxCanvasonfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}

// genaral application state
type State struct {
	BrushColor     color.Color
	BrustType      int
	SwatchSelected int
	FilePath       string
}

// reciver fuction for set the file path
func (state *State) SetFilePath(path string) {
	state.FilePath = path
}
