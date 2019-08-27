package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// File default config file name
var File = ".gitlabback"

// Config basic struct for backup configuration
type Config struct {
	BaseURL   string `json:"baseURL"`
	Token     string `json:"token"`
	BackupDir string `json:"backupdir"`
	SSH       bool   `json:"ssh"`
}

// Write creates a config file if it doesn't exist
// else overwrites it.
func Write(name string, c *Config) {
	var f *os.File
	if _, err := os.Stat(name); err != nil {
		f, err = os.Create(name)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		f, err = os.OpenFile(name, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
		if err != nil {
			log.Fatal(err.Error())
		}
		f.Truncate(0)
	}

	defer f.Close()

	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.write(f, b)
}

func (c *Config) write(w io.Writer, b []byte) {
	if _, err := w.Write(b); err != nil {
		log.Fatal(err.Error())
	}
}
