package telegram

import "github.com/Vinsektor/saved_reminder_bot/clients/telegram"

type Processor struct {
	tg     *telegram.TgClient
	offset int
	// storage
}

func New(client *telegram.TgClient) {

}
