// converts command line it's numeric args to °C, °F and K scale
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mr-susbedo/golang/ch2/tempconv"
)

func main() {
	for i, arg := range os.Args[1:] {
		temp, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		}

		f := tempconv.Fahrenheit(temp)
		c := tempconv.Celcius(temp)
		k := tempconv.Kelvin(temp)

		fmt.Printf("[Arg:%d]\t%f\n", i+1, temp)
		fmt.Printf("[Celcius]\n\t%s == %s\n\t%s == %s\n", c, tempconv.CToF(c), c, tempconv.CToK(c))
		fmt.Printf("[Fahrenheit]\n\t%s == %s\n\t%s == %s\n", f, tempconv.FToC(f), f, tempconv.FToK(f))
		fmt.Printf("[Kelvin]\n\t%s == %s\n\t%s == %s\n", k, tempconv.KToF(k), k, tempconv.KToC(k))
		fmt.Println(strings.Repeat("-", 36))
	}
}
