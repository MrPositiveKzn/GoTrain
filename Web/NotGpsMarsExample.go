package main

import (
	"fmt"
	"math"
)

type gps struct {
	current			location
	world			world
	destination		location
}

type location struct {
	name			string
	lat, long		float64
}

type world struct {
	radius	 		float64
}

func (l location) description() string {
	return fmt.Sprintf("%v (%.1f°, %.1f°)", l.name, l.lat, l.long)
}

func (g gps) distance() float64  {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string  {
	return fmt.Sprintf("%.1f km to %v", g.distance(), g.destination.description())
}

func (w world) distance(p1, p2 location) float64  {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

type rover struct {
	gps
}

func main() {
	mars := world{radius: 3389.5}
	bradbury := location{name: "Samara", lat: 4.2536, long: 12.3525}
	elisium := location{name: "Kazan", lat: 2.21412, long: 8.12412}
	gps := gps{
		current:     bradbury,
		world:       mars,
		destination: elisium,
	}
	curiosity := rover{gps}
	fmt.Println(curiosity.message())
}