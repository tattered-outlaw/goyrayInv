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
	width := 1000
	height := width
	fmt.Println("starting rendering...")
	start := time.Now().UnixMilli()
	pixelSource := engine.Prepare(width)
	pngWriter := PngWriter{width: width, height: height, image: image.NewRGBA(image.Rect(0, 0, width, height))}
	wgCount := 128
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
	defer wg.Done()
	for x := grn; x < width; x += grOf {
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
