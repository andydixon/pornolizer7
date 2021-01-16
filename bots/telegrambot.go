package bots

import (
	
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"Pornolizer7/pornoliser"

)

var bot *tgbotapi.BotAPI

// Thread to execute for the Telegram pornolizer bot
func TelegramThread() {
	var err error
	bot, err = tgbotapi.NewBotAPI("1497588048:AAExZXON7qmLvCGAlhmqHNfLFDnYBFj_864")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("[Telegram] Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, pornoliser.Pornolise(update.Message.Text, 70, "en", 1337))
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

// Send a message through telegram to Andy for whatever reason. Something that'll piss him off after a few hours no doubt
func MessageAdminThroughTelegram(message string) {
	msg := tgbotapi.NewMessage(811300921, message)
	bot.Send(msg)
}