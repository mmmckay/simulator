package main

import (
	"math/rand"

	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

var hexSizeX = 30
var hexSizeY = 20
var scale = float64(3)

// intial vertices for hex and generic square object contained within
var hexVertices = [][2]float64{
	{4, 0},
	{10, 0},
	{14, 6},
	{10, 12},
	{4, 12},
	{0, 6},
}
var objectVertices = [][2]float64{
	{4, 3},
	{10, 3},
	{10, 9},
	{4, 9},
}

// hex -> x,y coords
// hex 0,0 -> 4,0
// hex 1,0 -> 14,6
// hex 2,0 -> 24,0
// hex 3,0 -> 34,6
// hex 0,1 -> 4,12
// hex 0,2 -> 4,24
// hex 0,3 -> 4,36
// hex 1,1 -> 14,18
// hex 1,2 -> 14,30
// hex 1,3 -> 14,42
// hex 2,1 -> 24,12
// hex 2,2 -> 24,24
// hex 2,3 -> 24,36
// hex 3,3 -> 34,42

func main() {
	// scale points
	for index := range hexVertices {
		hexVertices[index][0] = hexVertices[index][0] * scale
		hexVertices[index][1] = hexVertices[index][1] * scale
	}
	for index := range objectVertices {
		objectVertices[index][0] = objectVertices[index][0] * scale
		objectVertices[index][1] = objectVertices[index][1] * scale
	}

	// initial canvas, sized by number of hexes
	// TODO: find way to extract scaling/tiling to one place
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     int(hexVertices[0][0] + float64(hexSizeX)*10*scale),
		Height:    int(hexVertices[0][1] + float64(hexSizeY*12+6)*scale),
		FrameRate: 1,
		Title:     "Hex",
	})

	// This gets called at the FrameRate rate above
	c.Draw(func(ctx *canvas.Context) {
		// iterate hexes
		for y := 0; y < hexSizeY; y++ {
			for x := 0; x < hexSizeX; x++ {
				// for each hex, move to the initial point
				// then draw lines between to each point defined in hexVertices
				ctx.MoveTo(hexVertices[0][0]+float64(x*10)*scale, hexVertices[0][1]+float64(y*12+(x%2)*6)*scale)
				for i := 0; i < len(hexVertices); i++ {
					if i < len(hexVertices)-1 {
						ctx.LineTo(hexVertices[i+1][0]+float64(x*10)*scale, hexVertices[i+1][1]+float64(y*12+(x%2)*6)*scale)
					} else {
						ctx.LineTo(hexVertices[0][0]+float64(x*10)*scale, hexVertices[0][1]+float64(y*12+(x%2)*6)*scale)
					}
				}
				// Set color based on random values I'm choosing for now
				// if the coor chosen is green, also optionally draw an object
				drawObject := false
				if x > 5 && x < 10 && y > 8 && y < 12 {
					ctx.SetColor(colornames.Blue)
				} else {
					ctx.SetColor(colornames.Green)
					if rand.Intn(50) < 8 {
						drawObject = true
					}
				}

				// draw the actual hex
				// FillPreserve fills in the shape drawn (hex) and preserves the path that has been drawn
				// then change the color to the desired border color and Stroke changes the lines themselves to that color
				ctx.FillPreserve()
				ctx.SetColor(colornames.Black)
				ctx.Stroke()

				// if an object should be drawn, perform the same line drawing and filling process for the smaller square defined inside the hex
				if drawObject {
					ctx.MoveTo(objectVertices[0][0]+float64(x*10)*scale, objectVertices[0][1]+float64(y*12+(x%2)*6)*scale)
					for i := 0; i < len(objectVertices); i++ {
						if i < len(objectVertices)-1 {
							ctx.LineTo(objectVertices[i+1][0]+float64(x*10)*scale, objectVertices[i+1][1]+float64(y*12+(x%2)*6)*scale)
						} else {
							ctx.LineTo(objectVertices[0][0]+float64(x*10)*scale, objectVertices[0][1]+float64(y*12+(x%2)*6)*scale)
						}
					}
					ctx.SetColor(colornames.Gray)
					ctx.Fill()
				}
			}
		}
	})
}
