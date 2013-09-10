package main

import (
	"fmt"
	"github.com/zx9597446/marchingsquare"
	"github.com/zx9597446/rdp"
	"image"
	"image/color"
	"image/png"
	"os"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func debugDraw(result []image.Point, oldfile, newfile string, what color.Color) {
	old, err := os.Open(oldfile)
	defer old.Close()
	panicIfErr(err)
	oldimg, _, err := image.Decode(old)
	panicIfErr(err)
	b := oldimg.Bounds()
	newimg := image.NewRGBA(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			newimg.Set(x, y, oldimg.At(x, y))
		}
	}
	for _, pt := range result {
		newimg.Set(pt.X, pt.Y, what)
	}
	file, err := os.Create(newfile)
	defer file.Close()
	panicIfErr(err)
	png.Encode(file, newimg)
}

func main() {
	ret := marchingsquare.ProcessWithFile("terrain.png", marchingsquare.TransparentTest)
	result := rdp.Process(ret, 0.50)
	fmt.Println(len(ret), len(result))
	debugDraw(result, "terrain.png", "new.png", color.RGBA{255, 0, 0, 255})
}
