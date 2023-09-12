package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main()  {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("not able to load env")
		return
	}

	token := os.Getenv("SLACK_OAUTH_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")

	client := slack.New(token, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Pretext: "Validação",
		Text: "Oi text aushauhs",
		Fields: []slack.AttachmentField{
			{
				Title: "Mensagem",
				Value: "oi",
			},
		},
		
	}

	_, timestamp, err := client.PostMessage(channelId, slack.MsgOptionAttachments(attachment))
	
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(timestamp)
}