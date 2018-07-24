package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	const format = "2006/01/02"
	fmt.Printf("%v\t10:00\t19:00\n", t.Format(format))
}
