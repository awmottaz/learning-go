// Surfaceplot computes an SVG rendering of a 3D surface function
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 1500, 800           // canvas size in pixels
	cells         = 200                 // number of grid cells
	xyrange       = 2.5                 // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * .06        // pixels per z unit
	angle         = math.Pi / 8         // angle of x, y axes
)

var sinAng, cosAng = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: #302e44; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find the point (x,y) at the corner of cell (i,j).
	x := xyrange * (float64(i)/cells - .5)
	y := xyrange * (float64(j)/cells - .5)

	// Compute the surface height z at the point (x,y).
	z := f(x, y)

	// Project the point (x,y,z) isometrically onto the 2D canvas (sx, sy)
	sx := width/2 + (x-y)*cosAng*xyscale
	sy := height/2 + (x+y)*sinAng*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	return 4 * math.Sin(2*math.Pi*x) * math.Cos(1.5*math.Pi*y) * (1 - x*x) * y * (1 - y)
}
