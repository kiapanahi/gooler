/*
Copyright © 2024 Kia Raad <kia.panahirad@gmail.com>
*/
package main

import (
	"context"
	"log"

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
}
