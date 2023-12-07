package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	envFilePath := os.Getenv("ENV_FILE_PATH")

	app, err := NewApp(ctx, envFilePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	err = app.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	app.Close()
}
