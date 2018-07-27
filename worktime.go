package main

import (
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	slackToken := os.Getenv("SLACK_TOKEN")

	if slackToken == "" {
		fmt.Fprintf(os.Stderr, "Slack token is empty\n")
		os.Exit(1)
	}

	api := slack.New(slackToken)
	params := slack.SearchParameters{}
	params.Count = 100

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

	response, _ := api.SearchMessages("from:me on:"+t.Format(format), params)
	messages := response.Matches

	lastMessage := messages[0]
	// lastMessage := messages[len(messages)-1]

	fmt.Println(lastMessage.Text)
	components := strings.Split(lastMessage.Timestamp, ".")
	intVal, _ := strconv.ParseInt(components[0], 10, 64)
	lastTime := time.Unix(intVal, 0)

	timeFormat := "15:04:05"
	fmt.Println(lastTime.Format(timeFormat))

	// for _, message := range messages.Matches {
	// 	fmt.Println(message.Text)
	// }

	// fmt.Printf("%v\t%v\t%v\n", t.Format(format), firstTime.Format(timeFormat), lastTime.format(timeFormat))
}
