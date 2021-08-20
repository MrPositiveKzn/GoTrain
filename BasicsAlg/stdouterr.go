package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout,"for output" )
	io.WriteString(os.Stderr, "for error")
	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 200; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}
	fmt.Fprintln(os.Stdout ,"\n")
}
