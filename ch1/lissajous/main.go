package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
)

// yellow, black, red, green, blue
var palette = []color.Color{color.RGBA{255, 255, 0, 255}, color.Black, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255}}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	f, _ := os.Create(filepath.Join(dir, "lissa-test.gif"))
	defer f.Close()
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // no. of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y osciallator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase diff

	for i := 0; i < nframes; i++ {
		twiceTheSizePlus1 := 2*size + 1
		rect := image.Rect(0, 0, twiceTheSizePlus1, twiceTheSizePlus1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(t/res) % 5
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // ignoring encoding errors
}
