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

func debugDraw(result []marchingsquare.Point, oldfile, newfile string, what color.Color) {
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
	points := make([]rdp.Point, 0)
	for _, v := range ret {
		points = append(points, rdp.Point{v.X, v.Y})
	}
	result := rdp.Process(points, 0.50)
	fmt.Println(len(ret), len(result))
	ret2 := make([]marchingsquare.Point, 0)
	for _, v := range result {
		ret2 = append(ret2, marchingsquare.Point{v.X, v.Y})
	}
	debugDraw(ret2, "terrain.png", "new.png", color.RGBA{255, 0, 0, 255})
}
