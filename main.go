package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kiapanahi/gooler/cmd"
	"github.com/kiapanahi/gooler/pkg/observability"
)

func main() {
	shutdown, err := observability.SetupOTelSDK(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Default().Println("OpenTelemetry SDK initialized")
	defer shutdown(context.Background())

	cmd.Execute()

	// Wait for SIGTERM notification
	terminationChannel := make(chan os.Signal, 1)
	signal.Notify(terminationChannel, syscall.SIGTERM)
	sig := <-terminationChannel

	log.Default().Printf("Received %s, shutting down gracefully", sig)
}
