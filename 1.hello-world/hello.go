package main

import (
	"fmt"
	"os"
)


var nama = os.Getenv("NAMA")

func main() {
	fmt.Println("Hello World Go: ", nama)
}
