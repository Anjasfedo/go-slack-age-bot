package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"

	"github.com/shomali11/slacker"
)

func main() {

}

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
