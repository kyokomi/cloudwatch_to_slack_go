package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/apex/go-apex"
	"github.com/tkuchiki/go-timezone"
)

var (
	slackWebHookURL string
	location        *time.Location
)

type PostMessage struct {
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Title      string   `json:"title"`
	Pretext    string   `json:"pretext"`
	Color      string   `json:"color"`
	Fields     []Field  `json:"fields"`
	Text       string   `json:"text"`
	MarkdownIn []string `json:"mrkdwn_in"`
	Fallback   *string  `json:"fallback"`
	AuthorName *string  `json:"author_name"`
	AuthorLink *string  `json:"author_link"`
	AuthorIcon *string  `json:"author_icon"`
	TitleLink  *string  `json:"title_link"`
	ImageURL   *string  `json:"image_url"`
	ThumbURL   *string  `json:"thumb_url"`
	Footer     *string  `json:"footer"`
	FooterIcon *string  `json:"footer_icon"`
	Ts         *int     `json:"ts"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func inLocationTimeString(timeStr string) string {
	changeTime, err := time.Parse("2006-01-02T15:04:05.999999999-0700", timeStr)
	if err != nil {
		log.Println(err)
		return timeStr
	}
	return changeTime.In(location).String()
}

func postMessageToSlack(message CloudWatchAlarmMessage) error {
	pm := PostMessage{
		Attachments: []Attachment{
			{
				Pretext:    "<!channel>",
				Color:      stateColor(message.NewStateValue),
				Title:      message.AlarmDescription,
				Text:       message.NewStateReason,
				MarkdownIn: []string{"fields"},
				Fields: []Field{
					{
						Title: "State",
						Value: fmt.Sprintf("`%s`", message.NewStateValue),
						Short: true,
					},
					{
						Title: "Name",
						Value: message.AlarmName,
						Short: true,
					},
					{
						Title: "Region",
						Value: message.Region,
						Short: true,
					},
					{
						Title: "Namespace",
						Value: fmt.Sprintf("`%s`", message.Trigger.Namespace),
						Short: true,
					},
					{
						Title: "Time",
						Value: inLocationTimeString(message.StateChangeTime),
						Short: false,
					},
				},
			},
		},
	}
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(pm); err != nil {
		log.Println("json.NewDecoder error: ", err.Error())
		return err
	}

	resp, err := http.DefaultClient.Post(slackWebHookURL, "application/json", &buf)
	if err != nil {
		log.Println("http.DefaultClient.Post error: ", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ioutil.ReadAll error: ", err.Error())
			return err
		}
		return errors.New(bytes.NewBuffer(data).String())
	}

	return nil
}

func stateColor(stateValue string) string {
	switch stateValue {
	case "OK":
		return "#7CD197"
	case "INSUFFICIENT_DATA":
		return "#FFC107"
	case "ALARM":
		return "#F35A00"
	default:
		return "#9E9E9E"
	}
}

func execute(event json.RawMessage) (interface{}, error) {
	slackWebHookURL = os.Getenv("SLACK_WEBHOOK_URL")
	timezoneName := os.Getenv("TIMEZONE_NAME")
	if offset, err := timezone.GetOffset(timezoneName); err != nil {
		log.Println(err)
		location = time.Local
	} else {
		location = time.FixedZone(timezoneName, offset)
	}

	records := SNSRecords{}
	if err := json.Unmarshal(event, &records); err != nil {
		return nil, err
	}
	log.Println(string(event))

	for _, r := range records.Records {
		message := CloudWatchAlarmMessage{}
		if err := json.Unmarshal([]byte(r.SNS.Message), &message); err != nil {
			log.Println("json.Unmarshal error: ", err)
			continue
		}

		if err := postMessageToSlack(message); err != nil {
			log.Println("postMessageToSlack error: ", err)
			continue
		}
	}
	return records, nil
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		return execute(event)
	})
}
