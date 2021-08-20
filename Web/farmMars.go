package main

import (
	"fmt"
	"math/rand"
	"time"
)

type honeyBee struct {
	name string
}

func (hb honeyBee) String() string {
	return hb.name
}

func (hb honeyBee) move() string {
	switch rand.Intn(2) {
	case 0:
		return "Bee jjujit"
	default:
		return "bee flying"
	}
}

func (hb honeyBee) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "Bee eat pilca"
	default:
		return "bee eat nectar"
	}
}

type wolf struct {
	name string
}

func (w wolf) String() string {
	return w.name
}

func (w wolf) move() string {
	switch rand.Intn(2) {
	case 0:
		return "amazin can be first"
	default:
		return "auf"
	}
}

func (w wolf) eat() string {
	switch rand.Intn(4) {
	case 0:
		return "wolf eat fox"
	case 1:
		return "wolf eat mouse"
	case 2:
		return "wolf eat wood"
	default:
		return "wolf eat human"
	}
}

type animal interface {
	move() 	string
	eat()	string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("animal eating", a, a.eat())
	default:
		fmt.Printf("animal moving", a, a.move())
	}
}

const sunset, sunrise = 18, 8

func main() {
	rand.Seed(time.Now().UnixNano())
	animals := []animal{
		honeyBee{"Jora"},
		wolf{name: "Shahzod"},
	}
	var sol, hour int

	for {
		fmt.Printf("%2d:00 ", hour)
		if hour < sunrise || hour >= sunset {
			fmt.Println("animals sleep.")
		} else {
			i := rand.Intn(len(animals))
			step(animals[i])
		}

		time.Sleep(500 * time.Millisecond)

		hour++
		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3 {
				break
			}
		}
	}
}