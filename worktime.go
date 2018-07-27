package main

import (
	"flag"
	"fmt"
	// "github.com/nlopes/slack"
	"os"
	"time"
)

func main() {
	slackToken := os.Getenv("SLACK_TOKEN")

	if slackToken == "" {
		fmt.Fprintf(os.Stderr, "Slack token is empty\n")
		os.Exit(1)
	}

	// -hオプション用文言
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [date]
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	const format = "2006/01/02"
	t := time.Now()

	if flag.Arg(0) != "" {
		t, _ = time.Parse(format, flag.Arg(0))
	}

	fmt.Printf("%v\t10:00\t19:00\n", t.Format(format))
}
