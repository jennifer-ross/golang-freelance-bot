# Golang Freelance Bot

[![Go Badge](https://img.shields.io/badge/Go-1.23-blue)](#)
[![Postgres](https://img.shields.io/badge/Postgres-%23316192.svg?logo=postgresql&logoColor=white)](#)
[![Redis](https://img.shields.io/badge/Redis-%23DD0031.svg?logo=redis&logoColor=white)](#)
[![Docker Badge](https://img.shields.io/badge/Docker-ready-green)](#)
[![License](https://img.shields.io/badge/license-MIT-green)](#)

[Русская версия](/docs/README_RU.md)

Golang Freelance Bot is a bot developed in Go, designed for quickly receiving orders from freelance platforms in real time and promptly notifying the user via Telegram. Thanks to its Docker integration, the bot can be easily deployed and configured to work in any environment.

# Attention

🚧 **This project is still under development and not ready for deployment.** 🚧

## Features

- **Instant Order Retrieval**: The bot is set up to quickly extract and notify users about new orders from selected freelance platforms.
- **Optimized for Performance**: Written in Go to ensure high performance and low resource consumption.
- **Docker Compatibility**: Easily built and launched as a Docker image for deployment in various environments.
- **Easy Setup**: Using Makefile, you can quickly build the Docker image and launch the bot.

## Supported Freelance Platforms
- [Kwork (ru)](https://kwork.ru/)

## Installation

To install and run the project, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/jennifer-ross/golang-freelance-bot.git
    cd golang-freelance-bot
    ```

2. [Set up the `.env` file and fill it with the required data](#example-env-file-configuration):

    ```bash
    mv .env.example .env
    ```

3. Build the Docker image:

    ```bash
    make build
    ```
   or
   ```bash
   docker-compose build
   ```

4. Start the container:

    ```bash
    make up-d
    ```
   or
    ```bash
    docker-compose up -d
    ```

## How to Use
After launching, the bot will automatically start tracking orders from the selected freelance platform. You can configure the bot's settings through the `.env` file.

### Example `.env` file configuration

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