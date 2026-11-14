package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	// Load environment variables from .env file
	godotenv.Load()

	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	if botToken == "" || botToken == "xoxb-YOUR_BOT_TOKEN_HERE" {
		log.Fatal("SLACK_BOT_TOKEN is not set. Please update .env file with your bot token")
	}
	if appToken == "" || appToken == "xapp-YOUR_APP_TOKEN_HERE" {
		log.Fatal("SLACK_APP_TOKEN is not set. Please update .env file with your app token (should start with xapp-)")
	}

	bot := slacker.NewClient(botToken, appToken)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "Yob calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("Error converting year to integer:", err)
			}
			age := 2024 - yob
			r := fmt.Sprintf("You are %d years old", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
