package main

import (
	"log/slog"

	"github.com/artemKapitonov/url-shortner/internal/app"
)

func main() {
	a := app.New()

	err := a.Run()
	if err != nil {
		slog.Error("Fail with starting application Error:", err)
	}
}
