// +build linux,!windows

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"
	"strings"
)

// init grabs home directory
func init() {
	usr, err := user.Current()
	if err != nil {
		return
	}
	File = usr.HomeDir + "/" + File
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

// Setup prompts for token and backup directory
func Setup(conf *Config) {
	var token string
	var dir string

	fmt.Print("Token: ")
	fmt.Scanf("%s", &token)
	fmt.Print("Backup Directory: ")
	fmt.Scanf("%s", &dir)

	if strings.TrimSpace(dir) != "" && !strings.HasSuffix(dir, "/") {
		dir += "/"
	}

	conf.Token = token
	conf.BackupDir = dir
}
