package main

import _ "image/png"

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	// "image/gif"
	"image/png"
	"log"
	"math/rand"
	"os"
	// "strconv"
)

var (
	imageHeight int
	imageWidth  int
)

func ImageRead(ImageFile string) (image image.Image) {
	// open "test.jpg"
	file, err := os.Open(ImageFile)
	if err != nil {
		log.Fatal(err)
	}

	// decode png into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	return img
}

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// singleImage := image.NewPaletted(image.Rect(0, 0, 100, 100),image.NewPaletted(r, p))
	// outputGif := gif.GIF{}
	black := color.RGBA{0, 0, 0, 255}

	draw.Draw(m, m.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
	// outpng, _ := os.Create("out.png")
	// toimg, _ := os.Create("test.gif")
	// defer toimg.Close()
	w := m.Bounds().Max.X
	h := m.Bounds().Max.Y
	fmt.Println("w", w, "h", h)
	// x := 0
	// y := 0
	// n := 500
	// prev := color.RGBA{0, 0, 0, 0}
	// for i := 0; i < n; i++ {
	// 	x = rand.Intn(w)
	// 	y = rand.Intn(h)
	// 	prev = m.At(x, y).(color.RGBA)
	// 	prev.R += 30
	// 	m.Set(x, y, prev)
	// }
	// stringval := ""
	numberOfFramesToGenerate := 1
	for imageGenerated := 0; imageGenerated < numberOfFramesToGenerate; imageGenerated++ {
		// fmt.Println(("out" + (imageGenerated) + ".png"))
		stringval := fmt.Sprint("out", imageGenerated, ".png")
		outpng, err := os.Create(stringval)
		if err != nil {
			panic(err.Error())
		}
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				pixel := m.At(x, y).(color.RGBA)
				// m.Set(x, y, singlePixel)
				test := rand.Intn(2)
				if test == 1 {
					pixel.R = 255
					pixel.G = 255
					pixel.B = 255
				} else {
					pixel.R = 0
					pixel.G = 0
					pixel.B = 0
				}
				m.Set(x, y, pixel)

			}

			// outputGif.Image = append(outputGif.Image, ImageRead("out.png").image.Paletted)
		}
		png.Encode(outpng, m)
		outpng.Close()
		//Add image to Gif we're going to output

	}
	// gif.EncodeAll(toimg, &outputGif)

}