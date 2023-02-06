package main

import (
	"flag"
	"log"
	tgClient "scanner_bot/clients/telegram"
	event_consumer "scanner_bot/consumer/event-consumer"
	"scanner_bot/events/telegram"
	"scanner_bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, token()),
		files.New(storagePath),
	)

	log.Print("service started...")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal()
	}

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
