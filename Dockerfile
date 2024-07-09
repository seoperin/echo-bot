# Используем официальный образ Golang
FROM golang:1.22-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем все файлы в контейнер
COPY . .

# Устанавливаем зависимости и компилируем приложение
RUN go mod tidy
RUN go build -o echo-bot .

# Устанавливаем переменную окружения для токена
ENV TELEGRAM_BOT_TOKEN="your-telegram-bot-token"

# Запускаем приложение
CMD ["./echo-bot"]
