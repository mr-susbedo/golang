package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

// yellow, black, red, green, blue
var palette = []color.Color{color.RGBA{255, 255, 0, 255}, color.Black, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255}}

// func main() {
// 	_, filename, _, _ := runtime.Caller(0)
// 	dir := filepath.Dir(filename)
// 	f, _ := os.Create(filepath.Join(dir, "lissa-test.gif"))
// 	defer f.Close()
// 	lissajous(f)
// }

type LissajousParams struct {
	Cycles  int
	Res     float64
	Size    int
	Nframes int
	Delay   int
}

func Lissajous(out io.Writer, params LissajousParams) {

	freq := rand.Float64() * 3.0 // relative frequency of y osciallator
	anim := gif.GIF{LoopCount: params.Nframes}
	phase := 0.0 // phase diff

	for i := 0; i < params.Nframes; i++ {
		twiceTheSizePlus1 := 2*params.Size + 1
		rect := image.Rect(0, 0, twiceTheSizePlus1, twiceTheSizePlus1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(params.Cycles)*2*math.Pi; t += params.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(t/params.Res) % 5
			img.SetColorIndex(params.Size+int(x*float64(params.Size)+0.5), params.Size+int(y*float64(params.Size)+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, params.Delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // ignoring encoding errors
}
