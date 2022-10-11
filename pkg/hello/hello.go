package hello

import "fmt"

const Name = "FudaTools"
const Version = "v0.0.1"
const Desc = "for fuda"

func Hello() {
	fmt.Printf("%s %s %s\n", Name, Version, Desc)
}
