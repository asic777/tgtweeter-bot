package main

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}
	botApi.Debug = true

	scraper := twitterscraper.New()

	// err := scraper.SetProxy("http://89.232.123.2:3128")
	// if err != nil {
	// 	panic(err)
	// }

	// for tweet := range scraper.SearchTweets(context.Background(),
	// 	"doge elon mask -filter:retweets", 50) {
	// 	if tweet.Error != nil {
	// 		panic(tweet.Error)
	// 	}
	// 	fmt.Println(tweet.Text)
	// }

	tweet, err := scraper.GetTweet("1589283421704290306")
	if err != nil {
		panic(err)
	}
	fmt.Println(tweet.Text)

	fmt.Println()

	profile, err := scraper.GetProfile("cz_binance")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", profile)

	fmt.Println()

	for tweet := range scraper.GetTweets(context.Background(), "cz_binance", 5) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		fmt.Println(tweet.Text)
		fmt.Println(tweet.TimeParsed)
		fmt.Println(tweet.Timestamp)
		fmt.Println(tweet.ID)
		fmt.Println(tweet.Hashtags)
		fmt.Println()
		fmt.Println(tweet)
	}
}
