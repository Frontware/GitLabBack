package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Frontware/GitLabBack/config"
)

// setup prompts for token and backup directory
func setup(conf *config.Config) {
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

func main() {

	token := flag.String("token", "", "GitLab token")
	baseURL := flag.String("url", "", "GitLab server base url")
	dir := flag.String("dir", "", "Backup directory")
	file := flag.String("config", config.File, "Config path")
	ssh := flag.Bool("ssh", false, "Enable ssh clone")

	flag.Parse()

	c, err := config.Read(*file)
	if err != nil {
		setup(c)
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
	}

	if *ssh {
		c.SSH = true
	}

	fmt.Printf("Creating backup in %s\n", c.BackupDir)

	backup(c)
}
