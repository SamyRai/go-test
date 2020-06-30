package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

var red color.Color = color.RGBA{R: 200, G: 30, B: 30, A: 255}
var green color.Color = color.RGBA{G: 255, A: 255}
var black color.Color = color.RGBA{A: 255}

func drawDiagonalLines(image image.RGBA, color color.Color, incline int, offset int) {
	for x := 0; x < 200; x++ {
		y := x/incline + offset
		image.Set(x, y, color)
	}
}

func drawHorizontalLines(image image.RGBA, color color.Color, offset int) {
	for x := 0; x < 200; x++ {
		y := offset
		image.Set(x, y, color)
	}
}

func main() {
	ourImage := image.NewRGBA(image.Rect(0, 0, 200, 200))
	draw.Draw(ourImage, ourImage.Bounds(), &image.Uniform{C: green}, image.Point{}, draw.Src)

	for i := -200; i < 200; i += 5 {
		drawDiagonalLines(*ourImage, red, 1, i)
	}

	for i := 0; i < 200; i += 5 {
		drawHorizontalLines(*ourImage, black, i)
	}

	file, err := os.Create("lines.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, ourImage)
}
