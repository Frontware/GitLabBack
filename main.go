package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Frontware/GitLabBack/config"
)

func main() {

	token := flag.String("token", "", "GitLab token")
	baseURL := flag.String("url", "", "GitLab server base url")
	dir := flag.String("dir", "", "Backup directory")
	file := flag.String("config", config.File, "Config path")
	ssh := flag.Bool("ssh", false, "Enable ssh clone")

	flag.Parse()

	c, err := config.Read(*file)
	if err != nil {
		config.Setup(c)
		fmt.Printf("Write config to %s\n", *file)
		config.Write(*file, c)
		os.Exit(-1)
	}

	if *token != "" {
		c.Token = *token
	}

	if *baseURL != "" {
		c.BaseURL = *baseURL
	}

	if *dir != "" {
		c.BackupDir = *dir
	} else {
		c.BackupDir = "."
	}

	if *ssh {
		c.SSH = true
	}

	fmt.Printf("Creating backup in %s\n", c.BackupDir)

	backup(c)
}
