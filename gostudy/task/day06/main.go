package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2001/02/03-15:03:05"))
	fmt.Println(os.Getgid())
}
