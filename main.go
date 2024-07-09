package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	messageCounts map[int64]int
	mutex         sync.Mutex
)

func init() {
	messageCounts = make(map[int64]int)
}

func incrementMessageCount(userID int64) int {
	mutex.Lock()
	defer mutex.Unlock()
	messageCounts[userID]++
	return messageCounts[userID]
}

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			userID := update.Message.From.ID
			userMessage := update.Message.Text

			// Increment the message count for the user
			count := incrementMessageCount(userID)

			// Prepare the response message
			responseText := fmt.Sprintf("Ты написал: %s\nВсего сообщений: %d", userMessage, count)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
			bot.Send(msg)
		}
	}
}
