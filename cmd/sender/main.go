package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	ConsumerKey       string `split_words:"true" required:"true"`
	ConsumerSecretKey string `split_words:"true" required:"true"`
	AccessToken       string `split_words:"true" required:"true"`
	AccessSecret      string `split_words:"true" required:"true"`
}

var (
	config EnvConfig
	client *twitter.Client
)

func receive(event cloudevents.Event) {
	if event.DataContentType() != "image/jpeg" {
		log.Print("Wrong content type for event:", event.DataContentType())
		return
	}

	result, resp, err := client.Media.Upload(event.Data(), event.DataContentType())
	if err != nil || resp.StatusCode >= 400 {
		log.Printf("Unable to upload media (%d): %s", resp.StatusCode, err)
		return
	}

	client.Statuses.Update("Hey, I added a caption", &twitter.StatusUpdateParams{MediaIds: []int64{result.MediaID}})

}

func main() {
	if err := envconfig.Process("twitter", &config); err != nil {
		log.Fatal("Unable to process environment:", err)
	}
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal("Unable to initialize CloudEvents client:", err)
	}

	authConfig := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecretKey)
	token := oauth1.NewToken(config.AccessToken, config.AccessSecret)
	httpClient := authConfig.Client(oauth1.NoContext, token)

	client = twitter.NewClient(httpClient)

	log.Fatal(c.StartReceiver(context.Background(), receive))
}
