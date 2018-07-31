// Surfaceplot computes an SVG rendering of a 3D surface function
package main

import (
	"fmt"
	"math"
	"sort"
)

const (
	width, height = 1500, 800           // canvas size in pixels
	cells         = 200                 // number of grid cells
	xyrange       = 2.2                 // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * .06        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes
)

var (
	sinAng, cosAng = math.Sin(angle), math.Cos(angle)
	zmin, zmax     = zMinMax()
)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, bz, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, cz, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, dz, ok := corner(i+1, j+1)
			if !ok {
				continue
			}

			// Find the average height of the four corners and normalize it.
			h := (1 - ((az+bz+cz+dz)/4-zmin)/(zmax-zmin)) * 240 // hue from blue to red

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
				"style='stroke: gray; fill: hsl(%.2f, 50%%, 50%%); stroke-width: 0.7'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, h)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find the point (x,y) at the corner of cell (i,j).
	x := xyrange * (float64(i)/cells - .5)
	y := xyrange * (float64(j)/cells - .5)

	// Compute the surface height z at the point (x,y).
	z, ok := f(x, y)

	// Project the point (x,y,z) isometrically onto the 2D canvas (sx, sy)
	sx := width/2 + (x-y)*cosAng*xyscale
	sy := height/2 + (x+y)*sinAng*xyscale - z*zscale
	return sx, sy, z, ok
}

func zMinMax() (float64, float64) {
	var vals []float64

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// Find the point (x,y) at the corner of cell (i,j).
			x := xyrange * (float64(i)/cells - .5)
			y := xyrange * (float64(j)/cells - .5)

			// Compute the surface height z at the point (x,y).
			z, ok := f(x, y)
			if !ok {
				continue
			}
			vals = append(vals, z)
		}
	}
	sort.Float64s(vals)
	return vals[0], vals[len(vals)-1]
}

func f(x, y float64) (float64, bool) {
	z := 4 * math.Sin(2*math.Pi*x) * math.Cos(1.5*math.Pi*y) * (1 - x*x) * y * (1 - y)
	if math.IsInf(z, 0) {
		return 0, false
	}
	return z, true
}
