package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("advanced/temp/txt.txt")
	if err != nil{
		panic(err)
	}
	defer f.Close()
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan(){
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())
	//big files

	fContent, err := ioutil.ReadFile("advanced/temp/txt.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println(string(fContent))
}//for small files
