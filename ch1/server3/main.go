// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler
func handler(w http.ResponseWriter, r *http.Request) {
	cycles := 5
	if v := r.FormValue("cycles"); v != "" {
		c, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
		} else {
			cycles = c
		}
	}
	lissajous(w, cycles)
}

func lissajous(out io.Writer, c int) {
	cycles := float64(c) // number of complete x oscillations
	const (
		res     = 0.0001 // angular resolution
		size    = 500    // canvas size
		nframes = 128    // number of animation frames
		delay   = 4      // delay between frames in 10ms units
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
