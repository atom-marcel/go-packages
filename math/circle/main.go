package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"time"
)

func main() {
	start := time.Now()
	img := image.NewNRGBA(image.Rect(0, 0, 200, 200))

	for y := 0; y < 200; y++ {
		for x := 0; x < 200; x++ {
			img.Set(x, y, image.White)
		}
	}

	matrix := make([][]int, 200)
	for i, _ := range matrix {
		matrix[i] = make([]int, 200)
	}

	mid_y := 200 / 2
	mid_x := 200 / 2
	r := 50

	for phi0 := -math.Pi + math.Pi/360; phi0 < math.Pi; phi0 += math.Pi / 360 {
		x, y := Polar(float64(r), phi0)
		x += mid_x
		y += mid_y
		img.Set(x, y, image.Black)
	}

	f, _ := os.Create("test.png")
	png.Encode(f, img)
	f.Close()

	elapsed := time.Since(start)

	fmt.Printf("Time elapsed: %s", elapsed)
}

func Polar(r float64, phi float64) (int, int) {
	x := r * math.Cos(phi)
	y := r * math.Sin(phi)

	return int(math.Round(x)), int(math.Round(y))
}
