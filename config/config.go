package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/morfien101/burgerJoint/menu"
)

type config struct {
	Verbose     bool `json:"verbose_logger"`
	BurgerStore struct {
		Staff struct {
			TillWorkers int `json:"till_workers"`
			Milkshakes  int `json:"milkshake_maker"`
			Fryer       int `json:"fryer"`
			Grill       int `json:"grill"`
			Soda        int `json:"soda"`
		} `json:"staff"`
		Equipment struct {
			MilkshakeMixer int `json:"milkshake_mixer"`
			Fryers         int `json:"fryers"`
			Grill          int `json:"grills"`
			Soda           int `json:"soda"`
		} `json:"equipment"`
		MenuItems []menu.MenuItem
	} `json:"burger_store"`
}

// DefaultConfig Prints a default configuration
func DefaultConfig() ([]byte, error) {
	b, err := json.MarshalIndent(new(config), "", "  ")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// PrintDefaultConfig will print a copy of the default configuration out to the console.
func PrintDefaultConfig() {
	config, err := DefaultConfig()
	if err != nil {
		log.Fatal("Failed to determine the default config.")
	}

	fmt.Println(string(config))
}
