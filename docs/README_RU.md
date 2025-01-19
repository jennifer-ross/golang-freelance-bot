# Golang Freelance Bot

[![Go Badge](https://img.shields.io/badge/Go-1.23-blue)](#)
[![Postgres](https://img.shields.io/badge/Postgres-%23316192.svg?logo=postgresql&logoColor=white)](#)
[![Redis](https://img.shields.io/badge/Redis-%23DD0031.svg?logo=redis&logoColor=white)](#)
[![Docker Badge](https://img.shields.io/badge/Docker-ready-green)](#)
[![License](https://img.shields.io/badge/license-MIT-green)](#)

[English version ](./README.md)

Golang Freelance Bot — это бот, разработанный на языке Go, предназначенный для быстрого получения заказов с фриланс-бирж в режиме реального времени и оперативного уведомления пользователя в telegram. Благодаря своей интеграции с Docker, бот может быть легко развёрнут и настроен для работы в любом окружении.

# Внимание

🚧 **Этот проект все еще находится в стадии разработки и не готов для запуска.** 🚧

## Особенности

- **Мгновенное получение заказов**: Бот настроен на оперативное извлечение и уведомление о новых заказах с выбранных фриланс-бирж.
- **Оптимизирован для производительности**: Написан на языке Go для высокой производительности и низкого потребления ресурсов.
- **Docker-совместимость**: Легко создаётся и запускается как Docker-образ для развёртывания в различных средах.
- **Лёгкая настройка**: С использованием Makefile вы можете быстро собрать Docker-образ и запустить бота.

## Список бирж
- [Kwork (ru)](https://kwork.ru/)

## Установка

Для того чтобы установить и запустить проект, выполните следующие шаги:

1. Клонируйте репозиторий:

    ```bash
    git clone https://github.com/jennifer-ross/golang-freelance-bot.git
    cd golang-freelance-bot
    ```

2. [Настройте файл `.env` и заполните его данными](#пример-настройки-env-файла):

    ```bash
    mv .env.example .env
    ```

3. Соберите Docker-образ:

    ```bash
    make build
    ```
   или
   ```bash
   docker-compose build
   ```

4. Запустите контейнер:

    ```bash
    make up-d
    ```
   или
    ```bash
    docker-compose up -d
    ```

## Как использовать

После запуска бот автоматически начинает отслеживать заказы с выбранной фриланс-биржи. Вы можете настроить параметры бота через файл `.env`

### Пример настройки `.env` файла

```dotenv
# App environment variables
GO111MODULE=on
CGO_ENABLED=0
TELEGRAM_BOT_TOKEN=YOUR_TELEGRAM_TOKEN_FROM_BOTFATHER

# PostgreSQL environment variables
POSTGRES_HOST=postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=golang_db

# Ports
POSTGRES_PORT=5432
REDIS_PORT=6379
