package main

import (
	"bufio"
	"fmt"
	"os"
)

//func main() {
//	var name string
//	fmt.Println("ur name")
//	fmt.Scanf("%s\n", &name)
//	fmt.Println("ur age")
//	var age int
//	fmt.Scanf("%d\n", &age)
//	fmt.Printf("Hello, ur name, %s, ur age %d\n", name, age)
//} 1 example

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan(){
		txt := sc.Text()
		fmt.Println("Эхо: %s\n", txt)
	}
}