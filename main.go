package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6468537462641-6452997310741-2ic5BGAeuqzFTQwLMRxIfVN8")
	os.Setenv("CHANANNEL_ID", "C06DSFTEJFK")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{
		os.Getenv("CHANANNEL_ID"),
	}

	fileArr := []string{"guru.txt"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Printf("Name : %s , Url : %s\n", file.Name, file.URL)
	}

}
