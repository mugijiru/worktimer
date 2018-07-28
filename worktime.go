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

const dateFormat = "2006/01/02"
const timeFormat = "15:04:05"

func main() {
	slackToken := os.Getenv("SLACK_TOKEN")

	if slackToken == "" {
		fmt.Fprintf(os.Stderr, "Slack token is empty\n")
		os.Exit(1)
	}

	api := slack.New(slackToken)

	// -hオプション用文言
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [date]
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	targetDate := time.Now()

	if flag.Arg(0) != "" {
		targetDate, _ = time.Parse(dateFormat, flag.Arg(0))
	}

	response := searchMessagesOnDate(api, targetDate)

	messages := response.Matches

	lastMessage := messages[0]
	// lastMessage := messages[len(messages)-1]

	fmt.Println(lastMessage.Text)

	lastTime := getTimeFromMessage(lastMessage)

	fmt.Println(lastTime.Format(timeFormat))

	// for _, message := range messages.Matches {
	// 	fmt.Println(message.Text)
	// }

	// fmt.Printf("%v\t%v\t%v\n", t.Format(format), firstTime.Format(timeFormat), lastTime.format(timeFormat))
}

func getTimeFromMessage(message slack.SearchMessage) time.Time {
	components := strings.Split(message.Timestamp, ".")
	Unixtime, _ := strconv.ParseInt(components[0], 10, 64)
	return time.Unix(Unixtime, 0)
}

func searchMessagesOnDate(api *slack.Client, date time.Time) *slack.SearchMessages {
	params := slack.SearchParameters{}
	params.Count = 100
	params.SortDirection = "timestamp"

	response, _ := api.SearchMessages("from:me on:"+date.Format(dateFormat), params)
	return response
}
