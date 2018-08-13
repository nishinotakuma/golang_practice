package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

var (
	width, height = 600, 320                     // canvas size in pixels
	cells         = 100                          // number of grid cells
	xyrange       = 30.0                         // axis ranges (-xyrange..+xyrange)
	xyscale       = float64(width) / 2 / xyrange // pixels per x or y unit
	zscale        = float64(height) * 0.4        // pixels per z unit
	angle         = math.Pi / 6                  // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func svg(w io.Writer) {
	zmin, zmax := minmax()
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s; fill: #222222' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(i, j, zmin, zmax), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")

		widthParam, ok := r.URL.Query()["width"]
		if ok {
			width, _ = strconv.Atoi(widthParam[0])
		}

		heightParam, ok := r.URL.Query()["width"]
		if ok {
			height, _ = strconv.Atoi(heightParam[0])
		}

		svg(w)
	})
	http.ListenAndServe(":8080", nil)
}

func minmax() (min float64, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xoff := 0; xoff <= 1; xoff++ {
				for yoff := 0; yoff <= 1; yoff++ {
					x := xyrange * (float64(i+xoff)/float64(cells) - 0.5)
					y := xyrange * (float64(j+yoff)/float64(cells) - 0.5)
					z := f(x, y)
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}
		}
	}
	return
}

func color(i, j int, zmin, zmax float64) string {
	min := math.NaN()
	max := math.NaN()
	for xoff := 0; xoff <= 1; xoff++ {
		for yoff := 0; yoff <= 1; yoff++ {
			x := xyrange * (float64(i+xoff)/float64(cells) - 0.5)
			y := xyrange * (float64(j+yoff)/float64(cells) - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	color := ""
	if math.Abs(max) > math.Abs(min) {
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}
		color = fmt.Sprintf("#%02x0000", int(red))
	} else {
		blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
		if blue > 255 {
			blue = 255
		}
		color = fmt.Sprintf("#0000%02x", int(blue))
	}
	return color
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	z := f(x, y)

	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
