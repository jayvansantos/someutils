package main

import (
	"fmt"
	"github.com/laher/someutils"
	"os"
)

func main() {
	err := someutils.Zip(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
