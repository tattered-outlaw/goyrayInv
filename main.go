package main

import (
	"fmt"
	"goray/internal"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
	"time"
)

func main() {
	scale := 4
	width := 960 * scale
	height := 540 * scale
	start := time.Now().UnixMilli()
	scene := internal.GroupScene1(width, height)
	engine := internal.NewEngine(scene)
	fmt.Printf("starting rendering at %d ms\n", time.Now().UnixMilli()-start)
	pngWriter := PngWriter{width: width, height: height, image: image.NewRGBA(image.Rect(0, 0, width, height))}
	wgCount := 16
	var wg sync.WaitGroup
	wg.Add(wgCount)
	for i := 0; i < wgCount; i++ {
		go render(width, height, engine, pngWriter, i, wgCount, &wg)
	}
	wg.Wait()
	fmt.Printf("rendering done in %d ms\n", time.Now().UnixMilli()-start)
	pngWriter.save("image")
}

type PixelDestination interface {
	setPixel(x, y int, c color.Color)
}

func render(width, height int, engine *internal.Engine, destination PixelDestination, grNumber, grCount int, wg *sync.WaitGroup) {
	debug := false
	debugX := 0 // width / 2
	debugY := 0 // height / 2
	defer wg.Done()
	for x := grNumber; x < width; x += grCount {
		for y := 0; y < height; y++ {
			if x == debugX && y == debugY {
				destination.setPixel(x, y, internal.GetPixel(engine, x, y))
			} else {
				if !debug {
					destination.setPixel(x, y, internal.GetPixel(engine, x, y))
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
	f, _ := os.Create("out/" + filename + ".png")
	err := png.Encode(f, p.image)
	if err != nil {
		panic(err)
	}
}
