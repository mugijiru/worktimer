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
const debug = false

func main() {
	flag.Usage = usage
	flag.Parse()

	api := initSlack()
	targetStartDate, _ := getTargetDates()

	response := searchMessagesOnDate(api, targetStartDate, 1)
	messages := response.Matches

	// asc とか関係なくページ内では降順で並んで返って来る
	firstMessage := messages[len(messages)-1]
	firstTime := getTimeFromMessage(firstMessage)

	lastPage := response.Paging.Pages

	if lastPage > 1 {
		response = searchMessagesOnDate(api, targetStartDate, lastPage)
		messages = response.Matches
	}

	lastMessage := messages[0]
	lastTime := getTimeFromMessage(lastMessage)

	if debug {
		fmt.Println("first: " + firstMessage.Text)
		fmt.Println("last: " + lastMessage.Text)
	}

	fmt.Printf("%v\t%v\t%v\n", targetStartDate.Format(dateFormat), firstTime.Format(timeFormat), lastTime.Format(timeFormat))
}

func getTimeFromMessage(message slack.SearchMessage) time.Time {
	components := strings.Split(message.Timestamp, ".")
	Unixtime, _ := strconv.ParseInt(components[0], 10, 64)
	return time.Unix(Unixtime, 0)
}

func searchMessagesOnDate(api *slack.Client, date time.Time, page int) *slack.SearchMessages {
	params := slack.NewSearchParameters()
	params.Count = 100
	params.SortDirection = "asc"
	params.Sort = "timestamp"
	params.Page = page

	response, _ := api.SearchMessages("from:me on:"+date.Format(dateFormat), params)
	return response
}

func initSlack() *slack.Client {
	slackToken := os.Getenv("SLACK_TOKEN")

	if slackToken == "" {
		fmt.Fprintf(os.Stderr, "Slack token is empty\n")
		os.Exit(1)
	}

	return slack.New(slackToken)
}

func getTargetDates() (time.Time, time.Time) {
	targetStartDate := time.Now()
	targetEndDate := targetStartDate

	if flag.Arg(0) != "" {
		targetStartDate, _ = time.Parse(dateFormat, flag.Arg(0))
	}

	if flag.Arg(1) != "" {
		targetEndDate, _ = time.Parse(dateFormat, flag.Arg(1))
	}

	return targetStartDate, targetEndDate
}

// -hオプション用文言
func usage() {
	fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [YYYY/MM/DD]
`, os.Args[0], os.Args[0])
	flag.PrintDefaults()
}
