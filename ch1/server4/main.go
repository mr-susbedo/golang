package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/mr-susbedo/golang/ch1/lissajous"
)

const (
	DefaultCycles  = 5     // number of complete x oscillator revolutions
	DefaultRes     = 0.001 // angular resolution
	DefaultSize    = 100   // image canvas covers [-size..+size]
	DefaultNFrames = 64    // no. of animation frames
	DefaultDelay   = 8     // delay between frames in 10ms units
)

func setDefaultParams(params lissajous.LissajousParams) lissajous.LissajousParams {
	if params.Cycles == 0 {
		params.Cycles = DefaultCycles
	}
	if params.Res == 0 {
		params.Res = DefaultRes
	}
	if params.Size == 0 {
		params.Size = DefaultSize
	}
	if params.Nframes == 0 {
		params.Nframes = DefaultNFrames
	}
	if params.Delay == 0 {
		params.Delay = DefaultDelay
	}
	return params
}

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}

		var params lissajous.LissajousParams

		for k, v := range r.Form {
			log.Printf("Form[%q] = %q\n", k, v)
			key := strings.ToLower(k)

			if len(v) > 0 {
				switch key {
				case "res":
					value, err := strconv.ParseFloat(v[0], 64) // ParseFloat instead of Atoi for "res"
					if err != nil {
						log.Printf("Error converting: %q with %q\n", key, v)
					} else {
						params.Res = value
					}
				default:
					value, err := strconv.Atoi(v[0])
					if err != nil {
						log.Printf("Error converting: %q with %q\n", key, v)
					}

					switch key {
					case "cycles":
						params.Cycles = value
					case "size":
						params.Size = value
					case "nframes":
						params.Nframes = value // Note: Corrected Nframes to NFrames
					case "delay":
						params.Delay = value
					default:
						log.Printf("Unknown param: %q\n", key)
					}
				}
			} else {
				log.Printf("No value provided for: %q\n", key)
			}

		}

		params = setDefaultParams(params)
		log.Printf("Params: %+v\n", params)
		lissajous.Lissajous(w, params)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
