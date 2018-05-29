// Lissajous-mc creates a lissajous GIF in green on black colors
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // red
	color.RGBA{0xff, 0x7f, 0x00, 0xff}, // orange
	color.RGBA{0xff, 0xff, 0x00, 0xff}, // yellow
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // green
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // blue
	color.RGBA{0x4b, 0x00, 0x82, 0xff}, // indigo
	color.RGBA{0x94, 0x00, 0xd3, 0xff}, // violet
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 10     // number of complete x oscillations
		res     = 0.0001 // angular resolution
		size    = 250    // canvas size
		nframes = 64     // number of animation frames
		delay   = 8      // delay between frames in 10ms units
	)
	relFreq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference between x and y oscillators

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*relFreq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(math.Ceil(7.0/(cycles*2*math.Pi)*t)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lissajous-gb:\t%v\n", err)
	}
}
