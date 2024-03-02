package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	//TODO logger : slog
	//TODO storage : postgresql
	//TODO client : echo
	//TODO run server
}
