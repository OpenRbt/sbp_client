package main

import (
	"context"
	"fmt"
	"os"
	app "sbp/internal/app"
)

func main() {
	ctx := context.Background()

	envFilePath := os.Getenv("ENV_FILE_PATH")

	// app init
	app, err := app.NewApp(ctx, envFilePath)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	// app run
	err = app.Run()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	// close all dependencies after the app terminates
	app.Close()
}
