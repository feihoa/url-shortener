package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// config cleanenv
	// logger slog
	// storage sqllite
	// router chi, chi render
	// run server
}
