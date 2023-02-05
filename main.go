package main

import (
	"flag"
	"fmt"
	"log"
	"scanner_bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient := telegram.NewClient(tgBotHost, token())

	fmt.Println(tgClient.Updates(10, 10))

	//fetcher = fetcher.New()

	//processor = processor.New()

	//consumer.Start(fetcher, processor)

}

func token() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram",
	)

	//помещает значение в Токен
	flag.Parse()

	if *token == "" {
		log.Fatal("token is missing")
	}

	return *token
}
