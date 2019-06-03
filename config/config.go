package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// File default config file name
const File = ".GitLabBack.json"

// Config basic struct for backup configuration
type Config struct {
	BaseURL   string `json:"baseURL"`
	Token     string `json:"token"`
	BackupDir string `json:"backupdir"`
	SSH       bool   `json:"ssh"`
}

// Read reads config file
func Read(file string) (*Config, error) {

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return &Config{}, err
	}

	conf := &Config{}

	if err = json.Unmarshal(b, conf); err != nil {
		return &Config{}, err
	}

	if conf.BackupDir != "" && !strings.HasSuffix(conf.BackupDir, "/") {
		conf.BackupDir += "/"
	}

	return conf, nil
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
