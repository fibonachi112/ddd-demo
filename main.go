package main

import (
	"repo/config"
	DiFabric "repo/internal/DI/fabric"
	"repo/internal/http"
)

func main() {
	Di := DiFabric.Make()
	app := http.Make(config.Config, &Di)
	app.Serve()
}
