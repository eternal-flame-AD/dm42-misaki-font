package main

import (
	"fmt"
	"image"
	"os"

	_ "image/png"
)

type Glyph [8]uint64

func MakeGlyphFromImg(onOffset func(x, y int) bool) (res Glyph) {
	for offsetX := 0; offsetX < 8; offsetX++ {
		for offsetY := 0; offsetY < 8; offsetY++ {
			if onOffset(offsetX, offsetY) {
				res[offsetX] |= 1 << offsetY
			}

		}
	}
	return
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	numGlyphX := img.Bounds().Dx() / 8
	numGlyphY := img.Bounds().Dy() / 8
	fmt.Println("{}")
	fmt.Println("LBL \"LFDAT\"")
	fmt.Printf("%d\n", numGlyphX*numGlyphY)
	fmt.Println("1")
	fmt.Println("NEWMAT")
	fmt.Println("STO \"fntDt\"")
	fmt.Println("INDEX \"fntDt\"")
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println("STOIJ")
	curGlyphIdx := 0
	for curglyphY := 0; curglyphY < numGlyphY; curglyphY++ {
		for curglyphX := 0; curglyphX < numGlyphX; curglyphX++ {
			if curGlyphIdx != 0 {
				fmt.Println("I+")
			}
			curGlyphIdx++
			glyph := MakeGlyphFromImg(func(x, y int) bool {
				r, g, b, _ := img.At(
					x+curglyphX*8+img.Bounds().Min.X,
					y+curglyphY*8+img.Bounds().Min.Y,
				).RGBA()
				return r+g+b == 0
			})

			el := uint64(0)
			mult := uint64(1)
			for i := 0; i < 8; i++ {
				el += glyph[i] * mult
				mult *= 256
			}
			fmt.Printf("%d\n", el)
			fmt.Println("STOEL")
		}
	}
	fmt.Println("END")
}
