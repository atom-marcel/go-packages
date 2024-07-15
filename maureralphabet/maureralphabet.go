package maureralphabet

import (
	"errors"
	"image"
	"image/draw"
	"strings"
)

type Buchstabe struct {
	Str string // Nur Großbuchstaben
	Num int    // 1 == 'A' und 26 == 'Z'
}

type ImageOptions struct {
	Padding    int
	Stride     int
	Wordlength int
	Lines      int
	PixSize    int
	PixPlot    int
}

func ConvertToBuchstaben(s string) ([]Buchstabe, error) {
	buchstaben := make([]Buchstabe, len(s))
	str := strings.ToUpper(s)

	for index, c := range str {
		num := int(c)
		if (num >= 65 && num <= 90) || (num >= 97 && num <= 122) || num == 47 || num == 10 || num == 32 {
			buchstaben[index] = Buchstabe{
				Str: string(c),
				Num: int(c) - 64,
			}
			continue
		} else {
			return make([]Buchstabe, 0), errors.New("Buchstabe ist nicht gültig! ES dürfen nur A-Z und a-z verwendet werden")
		}
	}

	return buchstaben, nil
}

func ImageBuchstaben(buchstaben []Buchstabe, img_opt ImageOptions) image.Image { //pad int, stride int, b_width int, lines int, line_width int, plot_size int
	plot := img_opt.PixPlot
	width := plot*img_opt.Wordlength + 2*img_opt.Padding + (img_opt.Stride * (img_opt.Wordlength - 1))
	height := plot*img_opt.Lines + (img_opt.Lines-1)*img_opt.Stride + 2*img_opt.Padding
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	// White background
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, image.White)
		}
	}

	buchstaben_imgs := make([]image.Image, len(buchstaben))
	for index, b := range buchstaben {
		buchstaben_imgs[index] = ImageBuchstabe(plot, plot, b, img_opt.PixSize)
	}

	x := img_opt.Padding
	y := img_opt.Padding
	i := 1
	for index, b_img := range buchstaben_imgs {
		if buchstaben[index].Str == "\n" || i > img_opt.Wordlength {
			y += plot + img_opt.Stride
			i = 1
			x = img_opt.Padding
		}
		if ((buchstaben[index].Str == "\n") && i == 1) || ((buchstaben[index].Num == 32-64) && i == 1) {
			continue
		}
		i++
		draw.Draw(img, image.Rect(x, y, x+plot, y+plot), b_img, image.Point{X: 0, Y: 0}, draw.Src)
		x += plot + img_opt.Stride
	}

	return img
}

func ImageBuchstabe(width int, height int, b Buchstabe, lw int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	matrix := PixMatrix(b, width, height, lw)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if matrix[y][x] == 1 {
				img.Set(x, y, image.Black)
			} else {
				img.Set(x, y, image.White)
			}
		}
	}

	return img
}

func PixMatrix(b Buchstabe, width int, height int, psize int) [][]int {
	matrix := make([][]int, height)

	for index := range matrix {
		matrix[index] = make([]int, width)
	}

	// Initialize Matrix with zeros
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			matrix[y][x] = 0
		}
	}

	switch b.Str {
	case "A":
		RightMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
	case "B":
		LeftMatrix(&matrix, psize)
		RightMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
	case "C":
		LeftMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
	case "D":
		RightMatrix(&matrix, psize)
		TopMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
	case "E":
		LeftMatrix(&matrix, psize)
		RightMatrix(&matrix, psize)
		TopMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
	case "F":
		LeftMatrix(&matrix, psize)
		TopMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
	case "G":
		TopMatrix(&matrix, psize)
		RightMatrix(&matrix, psize)
	case "H":
		LeftMatrix(&matrix, psize)
		RightMatrix(&matrix, psize)
		TopMatrix(&matrix, psize)
	case "I":
		TopMatrix(&matrix, psize)
		LeftMatrix(&matrix, psize)
	case "J":
		RightMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "K":
		LeftMatrix(&matrix, psize)
		RightMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "L":
		LeftMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "M":
		RightMatrix(&matrix, psize)
		TopMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "N":
		LeftMatrix(&matrix, psize)
		RightMatrix(&matrix, psize)
		TopMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "O":
		LeftMatrix(&matrix, psize)
		TopMatrix(&matrix, psize)
		BottomMatrix(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "P":
		TopMatrix(&matrix, psize)
		RightMatrix(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "Q":
		LeftMatrix(&matrix, psize)
		RightMatrix(&matrix, psize)
		TopMatrix(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "R":
		TopMatrix(&matrix, psize)
		LeftMatrix(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "S":
		TopAngle(&matrix, psize)
	case "T":
		LeftAngle(&matrix, psize)
	case "U":
		RightAngle(&matrix, psize)
	case "V":
		BottomAngle(&matrix, psize)
	case "W":
		TopAngle(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "X":
		LeftAngle(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "Y":
		RightAngle(&matrix, psize)
		DotMatrix(&matrix, psize)
	case "Z":
		BottomAngle(&matrix, psize)
		DotMatrix(&matrix, psize)
	default:
		return matrix
	}

	return matrix
}

func TopMatrix(m *[][]int, size int) {
	for y := 0; y < size; y++ {
		for x := 0; x < len((*m)[y]); x++ {
			(*m)[y][x] = 1
		}
	}
}

func LeftMatrix(m *[][]int, size int) {
	for y := 0; y < len(*m); y++ {
		for x := 0; x < size; x++ {
			(*m)[y][x] = 1
		}
	}
}

func RightMatrix(m *[][]int, size int) {
	for y := 0; y < len(*m); y++ {
		for x := len((*m)[y]) - size; x < len((*m)[y]); x++ {
			(*m)[y][x] = 1
		}
	}
}

func BottomMatrix(m *[][]int, size int) {
	for y := len(*m) - size; y < len(*m); y++ {
		for x := 0; x < len((*m)[y]); x++ {
			(*m)[y][x] = 1
		}
	}
}

func DotMatrix(m *[][]int, size int) {
	h_height := len(*m) / 2
	h_width := len((*m)[0]) / 2
	lb := h_width - size/2
	rb := h_width + size/2
	bb := h_height + size/2
	tb := h_height - size/2

	for y := tb; y < bb; y++ {
		for x := lb; x < rb; x++ {
			(*m)[y][x] = 1
		}
	}
}

func TopAngle(m *[][]int, size int) {
	for y := 0; y < len(*m); y++ {
		for x := 0; x < len((*m)[y]); x++ {
			y0 := 2 * x
			y1 := -2*x + len((*m)[y])*2
			if y0-size < y && y0+size > y {
				(*m)[y][x] = 1
			}
			if y1-size < y && y1+size > y {
				(*m)[y][x] = 1
			}
		}
	}
}

func BottomAngle(m *[][]int, size int) {
	for y := 0; y < len(*m); y++ {
		for x := 0; x < len((*m)[y]); x++ {
			y0 := 2*x - len((*m)[y])
			y1 := -2*x + len((*m)[y])
			if y0-size < y && y0+size > y {
				(*m)[y][x] = 1
			}
			if y1-size < y && y1+size > y {
				(*m)[y][x] = 1
			}
		}
	}
}

func LeftAngle(m *[][]int, size int) {
	for y := 0; y < len(*m); y++ {
		for x := 0; x < len((*m)[y]); x++ {
			var f float64 = .5
			y0 := int(f * float64(x))
			y1 := int(-f*float64(x)) + len((*m))
			if y0-size/2 < y && y0+size/2 > y {
				(*m)[y][x] = 1
			}
			if y1-size/2 < y && y1+size/2 > y {
				(*m)[y][x] = 1
			}
		}
	}
}

func RightAngle(m *[][]int, size int) {
	for y := 0; y < len(*m); y++ {
		for x := 0; x < len((*m)[y]); x++ {
			var f float64 = .5
			y0 := int(f*float64(x)) + len((*m))/2
			y1 := int(-f*float64(x)) + len((*m))/2
			if y0-size/2 < y && y0+size/2 > y {
				(*m)[y][x] = 1
			}
			if y1-size/2 < y && y1+size/2 > y {
				(*m)[y][x] = 1
			}
		}
	}
}
