package bot

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	messageCounts   map[int64]int
	lastMessageTime map[int64]time.Time
	mutex           sync.Mutex
)

func init() {
	messageCounts = make(map[int64]int)
	lastMessageTime = make(map[int64]time.Time)
}

func incrementMessageCount(userID int64) int {
	mutex.Lock()
	defer mutex.Unlock()
	messageCounts[userID]++
	return messageCounts[userID]
}

func updateLastMessageTime(userID int64) {
	mutex.Lock()
	defer mutex.Unlock()
	lastMessageTime[userID] = time.Now()
}

func getLastMessageTime(userID int64) time.Time {
	mutex.Lock()
	defer mutex.Unlock()
	return lastMessageTime[userID]
}

func StartBot(botToken string, db *sql.DB) {
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

			// Update the last message time for the user
			lastMessageTime := getLastMessageTime(userID)
			updateLastMessageTime(userID)

			// Calculate the time difference
			var timeDiff string
			if !lastMessageTime.IsZero() {
				timeDiff = time.Since(lastMessageTime).String()
			} else {
				timeDiff = "Это ваше первое сообщение."
			}

			// Convert userID to string
			userIDStr := strconv.FormatInt(userID, 10)

			// Prepare the response message
			responseText := fmt.Sprintf("Ты (%s) написал: %s\nВсего сообщений: %d\nПрошло времени с последнего сообщения: %s", userIDStr, userMessage, count, timeDiff)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
			bot.Send(msg)
		}
	}
}
