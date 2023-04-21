package main

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	isBotInitialized bool
	isNotAdminTold   bool
)

func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	switch {
	case strings.HasPrefix(update.Message.Text, "/start"):

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! I am your friendly neighborhood bot. How can I assist you?")
		bot.Send(msg)
		break

	case strings.HasPrefix(update.Message.Text, "/help"):

		help(bot, update)
		break

	case strings.HasPrefix(update.Message.Text, "/echo "):

		echo(bot, update)
		break

	case strings.HasPrefix(update.Message.Text, "/roll "):

		DiceRoll(bot, update)
		break

	case strings.HasPrefix(update.Message.Text, "/kick"):

		if update.Message.ReplyToMessage == nil {
			break
		}

		userID := update.Message.ReplyToMessage.From.ID
		kick(bot, update.Message.Chat.ID, userID)

		break

	default:
		break
	}
}