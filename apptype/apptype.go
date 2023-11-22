package apptype

import (
	"fyne.io/fyne/v2"
)

type BrustType = int

type PxCanvasonfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}
