package main

import (
	"fmt"
	"log"

	"github.com/asic777/tgtweeter-bot/pkg/config"
	"github.com/asic777/tgtweeter-bot/pkg/storage"
	"github.com/asic777/tgtweeter-bot/pkg/storage/boltdb"
	"github.com/asic777/tgtweeter-bot/pkg/telegram"
	"github.com/asic777/tgtweeter-bot/pkg/twitter"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	//twitter.Fmain()
	fmt.Println(1)
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(2)
	fmt.Println(cfg.TelegramToken)
	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}
	botApi.Debug = true
	fmt.Println(3)
	// pocketClient, err := pocket.NewClient(cfg.PocketConsumerKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db, err := initBolt()
	if err != nil {
		log.Fatal(err)
	}
	storage := boltdb.NewTokenStorage(db)
	fmt.Println(4)
	bot := telegram.NewBot(botApi, storage, cfg.Messages)

	// redirectServer := server.NewAuthServer(cfg.BotURL, storage, pocketClient)

	// go func() {
	// 	if err := redirectServer.Start(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	fmt.Println(5)
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initBolt() (*bolt.DB, error) {
	twitter.Fmain()
	db, err := bolt.Open("mask.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Batch(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(storage.AccessTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(storage.RequestTokens))
		return err
	}); err != nil {
		return nil, err
	}

	return db, nil
}
