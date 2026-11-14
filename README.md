# Slack Age Bot

A Slack bot that calculates a user's age from their year of birth.

## Requirements

- Go 1.26 or newer
- A Slack app with bot and app-level tokens

## Configuration

Create a `.env` file in this directory:

```dotenv
SLACK_BOT_TOKEN=xoxb-your-bot-token
SLACK_APP_TOKEN=xapp-your-app-token
```

Keep tokens private. The `.env` file is ignored by Git.

## Run

```bash
go mod download
go run .
```

## Command

Send the following command in a channel where the bot is present:

```text
my yob is 1990
```

The bot replies with the calculated age.
