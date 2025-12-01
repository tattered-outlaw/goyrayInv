package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
	"time"

	"goray/engine"
)

func main() {
	f := engine.BallScene
	scale := 4
	width := 960 * scale
	height := 540 * scale
	fmt.Println("starting rendering...")
	start := time.Now().UnixMilli()
	pixelSource := f(width, height)
	pngWriter := PngWriter{width: width, height: height, image: image.NewRGBA(image.Rect(0, 0, width, height))}
	wgCount := 1
	var wg sync.WaitGroup
	wg.Add(wgCount)
	for i := 0; i < wgCount; i++ {
		go render(width, height, pixelSource, pngWriter, i, wgCount, &wg)
	}
	wg.Wait()
	fmt.Printf("rendering done in %d ms\n", time.Now().UnixMilli()-start)
	pngWriter.save("image")
}

type PixelSource interface {
	GetPixel(x, y int) color.Color
}

type PixelDestination interface {
	setPixel(x, y int, c color.Color)
}

func render(width, height int, source PixelSource, destination PixelDestination, grn, grOf int, wg *sync.WaitGroup) {
	debug := false
	debugX := 0 // width / 2
	debugY := 0 // height / 2
	defer wg.Done()
	for x := grn; x < width; x += grOf {
		for y := 0; y < height; y++ {
			if x == debugX && y == debugY {
				destination.setPixel(x, y, source.GetPixel(x, y))
			} else {
				if !debug {
					destination.setPixel(x, y, source.GetPixel(x, y))
				}
			}
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
