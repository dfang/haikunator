package main

import (
	"fmt"

	haikunator "github.com/dfang/haikunator/v2"
)

func main() {
	haikunator := haikunator.New()
	fmt.Println(haikunator.Haikunate())
}
