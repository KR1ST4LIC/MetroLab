package user

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartMsg(userID int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(userID, "Хорошей игры!")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🖱 кликать 🖱"),
			tgbotapi.NewKeyboardButton("💶бустер автокликов💶"),
		),
		tgbotapi.NewKeyboardButtonRow(
			//tgbotapi.NewKeyboardButton("🏪 магазин 🏪"),
			tgbotapi.NewKeyboardButton("💶баланс💶"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("💰 реклама 💰"),
			//tgbotapi.NewKeyboardButton("🏠 кланы 🏠"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🔝Топ игроков🔝"),
			//tgbotapi.NewKeyboardButton("🏠 кланы 🏠"),
		),
	)
	return msg
}
