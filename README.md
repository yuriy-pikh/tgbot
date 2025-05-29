# My Telegram Bot in Go

This is a simple Telegram bot written in Go, using the `telebot` and `cobra` libraries. It is built using a Makefile for task automation and Docker for containerization.

## Features

This bot has the following commands:

* `/start` - Greets you and starts the conversation.
* `/help` - Shows a list of all available commands and their descriptions.
* `/echo <text>` - The bot will repeat any text you send after this command.
* `/wordcount <text>` - The bot will count the number of words in the provided text.
* `version` (CLI command) - Shows the program version.

## Installation and Startup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/yuriy-pikh/tgbot.git # Or the URL of your fork
    cd tgbot
    ```
2.  **Get a bot token:**
    Create your bot via BotFather in Telegram and get its token.
3.  **Set the `TELE_TOKEN` environment variable:**
    It's best to set this variable for the current terminal session before running the bot. To avoid saving the token in your command history, use:
    ```bash
    read -s -p "Enter your TELE_TOKEN: " TELE_TOKEN
    export TELE_TOKEN
    ```
    *Note: This method sets the variable only for the current session. For permanent storage, you can add `export TELE_TOKEN="YOUR_BOT_TOKEN"` to your shell's configuration file (e.g., `.bashrc`, `.zshrc`), but this is less secure.*

4.  **Build and Run:**

    *   **Local launch (using Go):**
        This method is suitable for development and quick tests.
        ```bash
        go run main.go tgbot
        # or `go run main.go start`
        ```

    *   **Local launch (using Makefile for building):**
        The Makefile automates the process of building the binary file.
        ```bash
        make build
        ./tgbot start
        ```

    *   **Running with Docker:**
        This method uses Docker to create and run a containerized version of the bot.

        1.  **Build the Docker image:**
            The Makefile will determine the version based on Git tags and commit hash.
            ```bash
            make image
            ```
            This will create an image with a tag like `urapikh/tgbot:0.1.0-a1b2c3d-amd64` or `urapikh/tgbot:a1b2c3d-amd64` (if there are no Git tags). You will see the exact tag in the output of the `make image` command.

        2.  **Run the container:**
            Use the tag obtained in the previous step.
            ```bash
            # Replace <image-tag> with the actual tag of your image (e.g., urapikh/tgbot:0.1.0-a1b2c3d-amd64)
            docker run -d -e TELE_TOKEN=$TELE_TOKEN --name my-telegram-bot <image-tag>
            ```
            For example:
            ```bash
            docker run -d -e TELE_TOKEN=$TELE_TOKEN --name my-telegram-bot urapikh/tgbot:a1b2c3d-amd64
            ```

## Testing the Bot

Find your bot on Telegram (e.g., t.me/YOUR_BOT_NAME_bot) and try the following commands:

* `/start`
* `/help`
* `/echo Hello, bot!`
* `/wordcount How many words are here?`
* Write any arbitrary text to see how the bot reacts to unknown commands or regular messages.

## Makefile Commands

* `make build`: Builds the application binary.
* `make image`: Builds the Docker image.
* `make push`: (If Docker Hub is configured) Pushes the image to the registry.
* `make clean`: Removes the built binary.
* `make format`: Formats the code.
* `make lint`: Runs the linter.
* `make test`: Runs tests.