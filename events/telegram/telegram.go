package telegram

import "scanner_bot/clients/telegram"

type EventProcessor struct {
	tg    *telegram.Client
	limit int
	//storage
}

func New(client *telegram.Client, storage) {

}
