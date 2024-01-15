# Slack Age Calculator 🎂

This Go program calculates a user's age based on the year of birth (yob) provided in a Slack command. The program utilizes the slacker library for Slack bot development and reads environment variables from a .env file.

## Usage 🚀

1. Create a `.env` file in the project root with the following format:

   ```env
    SLACK_BOT_TOKEN=your_slack_bot_token
    SLACK_APP_TOKEN=your_slack_app_token
   ```

2. run the program:

   ```
    go run main.go
   ```

The bot will listen for commands in Slack and respond with the calculated age.

## Code Explanation 📜

### main Function 🚀

#### Load Environment Variables:

- Calls the `loadEnv` function to read environment variables from the `.env` file.

#### Create Slack API Client:

- Uses the Slack API token from the environment variables to create a new Slack API client.

#### Define Slack Bot Command:

- Defines a Slack bot command that calculates age based on the provided year of birth.

#### Listen for Slack Events:

- Sets up context for the bot and starts listening for Slack events.

### loadEnv Function 🌐

#### Open and Read .env File:

- Opens the `.env` file and reads it line by line.

#### Set Environment Variables:

- Parses key-value pairs from each line and sets them as environment variables.

### Handler Function 📊

#### Extract Year of Birth (yob):

- Extracts the "year" parameter from the command request.

#### Calculate Age:

- Converts the year parameter to an integer and calculates the age based on the current year.

#### Prepare Response:

- Prepares a response message with the calculated age.

#### Reply to Slack Channel:

- Replies to the Slack channel with the calculated age.

## Closing Notes 📝

Feel free to adjust the configuration, and if you encounter any issues or have suggestions for improvement, please open an issue or submit a pull request.

Happy coding! 🚀👨‍💻
