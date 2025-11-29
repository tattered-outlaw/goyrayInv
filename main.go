package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"

	"goray/engine"
)

func main() {
	width := 1000
	height := width
	fmt.Println("starting rendering...")
	start := time.Now().UnixMilli()
	pixelSource := engine.Prepare(width)
	pngWriter := PngWriter{width: width, height: height, image: image.NewRGBA(image.Rect(0, 0, width, height))}
	render(width, height, pixelSource, pngWriter)
	fmt.Printf("rendering done in %d ms\n", time.Now().UnixMilli()-start)
	pngWriter.save("image")
}

type PixelSource interface {
	GetPixel(x, y int) color.Color
}

type PixelDestination interface {
	setPixel(x, y int, c color.Color)
}

func render(width, height int, source PixelSource, destination PixelDestination) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			destination.setPixel(x, y, source.GetPixel(x, y))
		}
	}
}

type PngWriter struct {
	width, height int
	image         *image.RGBA
}

func (p PngWriter) setPixel(x, y int, c color.Color) {
	if x >= 0 && x < p.width && y >= 0 && y < p.height {
		p.image.Set(x, y, c)
	}
}

func (p PngWriter) save(filename string) {
	f, _ := os.Create(filename + ".png")
	err := png.Encode(f, p.image)
	if err != nil {
		panic(err)
	}
}
