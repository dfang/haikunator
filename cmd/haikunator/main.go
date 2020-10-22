package main

import (
	"fmt"

	haikunator "github.com/dfang/haikunatorgo/v2"
)

func main() {
	haikunator := haikunator.New()
	fmt.Println(haikunator.Haikunate())
}
