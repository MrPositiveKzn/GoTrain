package main

import "fmt"

type report struct {
	sol			int
	temperature	temperature
	location	location
}

type temperature struct {
	high, low	celsius
}

//type location struct {  dublicate for notgpsexample

	lat, long	float64
}

type celsius float64

func (t temperature) average() celsius{
	return (t.high + t.low) / 2
}

func main() {
	bradbury := location{lat: 124.252, long: 124.47423234}
	t := temperature{low: 32.1, high: 48.2}
	report := report{sol: 2, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v° C\n", report.temperature.high)
	fmt.Printf("average temp"  , report.temperature.average())

}
