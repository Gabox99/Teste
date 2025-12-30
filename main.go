package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("O seu OS é:", runtime.GOOS)
	fmt.Println("O seu ARCH é:", runtime.GOARCH)
}
