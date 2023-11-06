/*
Package temperature provides utilities for converting temperatures between
Celsius, Fahrenheit, and Kelvin scales.

Usage:

To convert temperatures, simply use the provided functions like ToFahrenheit,
ToCelsius, and ToKelvin, passing in the temperature and the unit from which
you're converting.

Example:

c := tempconv.FToC(tempconv.Fahrenheit(212))
fmt.Println(c) // Output: 100

f := tempconv.CToF(tempconv.Celsius(100))
fmt.Println(f) // Output: 212

The package also provides types Celsius, Fahrenheit, and Kelvin to represent
temperatures in each scale and avoid confusion between numeric values.

Note:

This package does not handle temperatures below absolute zero.
*/
package tempconv
