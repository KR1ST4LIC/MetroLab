package manager

import (
	"context"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"MetroLab/internal/storage"
)

type Manager struct {
	storage storage.Storage
}

func NewManager(s storage.Storage) *Manager {
	return &Manager{
		storage: s,
	}
}

func (m *Manager) Run(ctx context.Context, token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		localUpdate := update
		m.handleBotUpdates(ctx, localUpdate, bot)
	}
}

func (m *Manager) handleBotUpdates(ctx context.Context, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil {
		m.ParsingUserMessage(ctx, update, bot)
	}

	if update.CallbackQuery != nil {
		m.ParsingUserCbd(ctx, update, bot)
		return
	}
}

func (m *Manager) ParsingUserMessage(ctx context.Context, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	text := update.Message.Text
	//userID := update.Message.From.ID

	switch text {
	case "/start":
	}
}

func (m *Manager) ParsingUserCbd(ctx context.Context, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	cbd := strings.Split(update.CallbackData(), "?")

	switch cbd[0] {
	case "getmoney":
	}
}
