package main

import (
	"fmt"
	daHelper "test_modules/helper"

	hellomod "github.com/donvito/hellomod"
	hellomodV2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	hellomodV2.SayHello("Melvin")
	fmt.Println(daHelper.Addition(1, 2))
}
