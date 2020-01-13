package main

import (
	"fmt"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Printf("hello word\n")
	fmt.Printf(os.Args[0])
}
