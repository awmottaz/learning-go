// Mandelbrot emits a PNG image of the Mandelbrot set.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2048, 2048
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var i int
	for py := 0; py < height; py++ {
		yhi := (float64(py)+0.25)/height*(ymax-ymin) + ymin
		ylo := (float64(py)-0.25)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			xhi := (float64(px)+0.25)/width*(xmax-xmin) + xmin
			xlo := (float64(px)-0.25)/width*(xmax-xmin) + xmin

			z1 := complex(xhi, yhi)
			z2 := complex(xhi, ylo)
			z3 := complex(xlo, yhi)
			z4 := complex(xlo, ylo)

			c1 := mandelbrot(z1)
			c2 := mandelbrot(z2)
			c3 := mandelbrot(z3)
			c4 := mandelbrot(z4)
			avgC := avgRGBA(c1, c2, c3, c4)
			if i%1000 == 0 {
				log.Println(c1, c2, c3, c4, avgC)
			}
			i++
			// Image point (px, py) represents complex value z.
			img.Set(px, py, avgC)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: Ignoring errors.
}

func mandelbrot(z complex128) color.RGBA {
	const (
		iterations = 200
		contrast   = 35
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{contrast * n, 0, 40, contrast * n}
		}
	}
	return color.RGBA{0, 0, 0, 1}
}

func avgRGBA(colors ...color.RGBA) color.RGBA {
	var r, g, b, a, i int

	for _, c := range colors {
		r += int(c.R) * int(c.R)
		g += int(c.G) * int(c.G)
		b += int(c.B) * int(c.B)
		a += int(c.A) * int(c.A)
		i++
	}

	avgR := uint8(math.Sqrt(float64(r) / float64(i)))
	avgG := uint8(math.Sqrt(float64(g) / float64(i)))
	avgB := uint8(math.Sqrt(float64(b) / float64(i)))
	avgA := uint8(math.Sqrt(float64(a) / float64(i)))

	return color.RGBA{avgR, avgG, avgB, avgA}
}
