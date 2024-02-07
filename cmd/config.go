package main

import (
	"encoding/json"
	"io"
	"os"
)

type config struct {
	DBstring string
	JWTkey   string
}

var CFG *config

func parseConfig() error {
	file, err := os.Open("cfg.json")
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	var cfg config
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return err
	}
	CFG = &cfg
	return nil
}
