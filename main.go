package main

import (
	"log"

	"echo-bot/bot"
	"echo-bot/config"
	"echo-bot/db"
)

func main() {
	// Load configuration
	botToken, connStr := config.LoadEnv()

	// Connect to the database
	database, err := db.ConnectDB(connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer database.Close()

	// Start the bot
	bot.StartBot(botToken, database)
}
