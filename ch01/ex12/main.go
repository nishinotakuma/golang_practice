// ./ex12
// http://localhost:8000/?cycles=20
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
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 255}}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	var cycles = 5
	handler := func(w http.ResponseWriter, r *http.Request) {
		cyclesStr := r.FormValue("cycles")
		if cyclesStr != "" {
			var err error
			cycles, err = strconv.Atoi(cyclesStr)
			if err != nil {
				fmt.Fprintf(w, "bad cycles param: %s", cyclesStr, err)
				return
			}
		}
		lissajous(cycles, w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func lissajous(cycles int, out io.Writer) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
