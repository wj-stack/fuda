package hello

import "fmt"

const Name = "FudaTools"
const Version = "V1.0"
const Desc = "for fuda"

func Hello() {
	fmt.Printf("%s %s %s\n", Name, Version, Desc)
}
