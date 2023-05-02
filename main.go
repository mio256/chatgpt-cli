package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ayush6624/go-chatgpt"
)

func main() {
	env()
	key := os.Getenv("OPENAI_API_KEY")

	var messages []chatgpt.ChatMessage

	ctx := context.Background()

	client, err := chatgpt.NewClient(key)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Print("> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		str := scanner.Text()

		if str == "exit" {
			break
		}

		messages = append(messages, chatgpt.ChatMessage{
			Role:    chatgpt.ChatGPTModelRoleUser,
			Content: str,
		})

		req := chatgpt.ChatCompletionRequest{
			Model:    chatgpt.GPT35Turbo,
			Messages: messages,
		}

		res, err := client.Send(ctx, &req)
		if err != nil {
			fmt.Println(err)
		}

		message := res.Choices[0].Message
		fmt.Println(message.Content)

		messages = append(messages, message)
	}
}
