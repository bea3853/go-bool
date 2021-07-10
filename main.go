package main

import "fmt"

func main() {

	estasol := true
	disposta := true
	saindo := false

	vouAPraia := estasol && disposta && saindo

	fmt.Println(vouAPraia)
	fmt.Printf("%T", vouAPraia)
}
