# Переменные
REPO_URL=https://github.com/seoperin/echo-bot.git
SERVICE_NAME=echo-bot

# Обновить репозиторий
update:
	git pull origin main

# Пересобрать Docker образ
build:
	docker compose build

# Перезапустить контейнеры
up:
	docker compose up -d

# Остановить контейнеры
down:
	docker compose down

# Запуск
run: build up

# Полный деплой
deploy: update run