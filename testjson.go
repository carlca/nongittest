package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"
)

// Config struct is read from ~/.golog config file
type Config struct {
	UseShortMessages bool
}

func main() {
	usr, err := user.Current()
	if err != nil {
		reportError(err)
	}
	configFile := path.Join(usr.HomeDir, ".golog")
	file, err := os.Open(configFile)
	if err != nil {
		reportError(err)
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		reportError(err)
	}
	if config.UseShortMessages {
		fmt.Println("Use short messages")
	} else {
		fmt.Println("Use long messages")
	}
}

func reportError(err error) {
	fmt.Println("error:", err)
}
