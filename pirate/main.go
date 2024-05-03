package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/henomis/lingoose/llm/openai"
	"github.com/henomis/lingoose/thread"
)

func main() {
	if err := getOrSetKey(); err != nil {
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Text :")

	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	myThread := thread.New().AddMessages(
		thread.NewSystemMessage().AddContent(
			thread.NewTextContent("All replies must be given in a pirate style of speech"),
		),
		thread.NewUserMessage().AddContent(
			thread.NewTextContent(text),
		),
	)

	if err = openai.New().Generate(context.Background(), myThread); err != nil {
		fmt.Println(err)
		fmt.Println("Error: something went wrong,Please ckeck your API Key & account")
		os.Exit(1)
	}

	fmt.Println("Pirate:" + myThread.LastMessage().Contents[0].AsString())
	os.Exit(0)
}
