package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// -hオプション用文言
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [date]\n`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	t := time.Now()
	const format = "2006/01/02"
	fmt.Printf("%v\t10:00\t19:00\n", t.Format(format))
}
