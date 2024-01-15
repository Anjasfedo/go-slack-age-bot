package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/shomali11/slacker"
)

// Main function
func main() {
	// Load environment variables from .env file
	if err := loadEnv(); err != nil {
		log.Fatal("Error on Load .env:", err)
	}

	// Create a new Slacker bot instance
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Goroutine to print command events asynchronously
	go printCommandEvents(bot.CommandEvents())

	// Define a Slack bot command
	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"My yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			// Extract the "year" parameter from the command request
			year := request.Param("year")

			// Convert the year parameter to an integer
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.Reply("Invalid year format. Please provide a valid year.")
				return
			}

			// Calculate the age based on the current year
			currentYear := time.Now().Year()
			age := currentYear - yob

			// Prepare the response message
			r := fmt.Sprintf("Your current age is %d years.", age)

			// Reply to the Slack channel with the calculated age
			response.Reply(r)
		},
	})

	// Set up context for the bot
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the Slack bot and listen for events
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// Function to load environment variables from .env file
func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		// Split each line into key-value pairs and set environment variables
		pair := strings.SplitN(line, "=", 2)
		if len(pair) == 2 {
			key, value := pair[0], pair[1]
			os.Setenv(key, value)
		} else {
			log.Printf("Invalid line in .env file: %s", line)
		}
	}

	return nil
}

// Function to print command events
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		// Print information about each command event
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
