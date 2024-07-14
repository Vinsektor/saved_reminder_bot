package main

import (
	"flag"
	"log"

	tgClient "github.com/Vinsektor/saved_reminder_bot/clients/telegram"
	event_consumer "github.com/Vinsektor/saved_reminder_bot/consumer/event-consumer"
	"github.com/Vinsektor/saved_reminder_bot/events/telegram"
	"github.com/Vinsektor/saved_reminder_bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Printf("service started")

	consumer := event_consumer.New(eventProcessor, eventProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
