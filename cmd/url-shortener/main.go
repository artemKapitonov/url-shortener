package main

import "github.com/artemKapitonov/url-shortener/internal/app"

func main() {
	a := app.New()

	if err := a.Run(); err != nil {
		panic(err)
	}

}
