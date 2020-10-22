package main

import (
	"fmt"

	"github.com/dfang/haikunator"
)

func main() {
	haikunator := haikunator.New()
	fmt.Println(haikunator.Haikunate())
}
