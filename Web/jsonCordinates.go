package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type cordinate struct {
	d, m, s float64
	h       rune
}

// String форматирует координаты DMS
func (c cordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

// decimal конвертирует координаты d/m/s в десятичные градусы
func (c cordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}


func (c cordinate) MarshalJson() ([]byte, error) {
	return	json.Marshal(struct {
		DD	float64	`json:"decimal"`
		DMS	string	`json:"dms"`
		D	float64	`json:"degrees"`
		M	float64	`json:"minutes"`
		S	float64	`json:"seconds"`
		H	string	`json:"hemisphere"`
	}{
		DD:  c.decimal(),
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}

type location2 struct {
	Name	string	`json:"name"`
	Lat		cordinate	`json:"latitude"`
	Long	cordinate	`json:"longitude"`
}

func main() {
	elysium := location2{
		Name: "Elysium Planitia",
		Lat:  cordinate{4, 30, 0.0, 'N'},
		Long: cordinate{135, 54, 0.0, 'E'},
	}

	bytes, err := json.MarshalIndent(elysium, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))
}