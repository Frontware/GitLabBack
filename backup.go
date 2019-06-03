package main

import (
	"log"
	"os"
	"strings"

	"github.com/Frontware/GitLabBack/config"
	"github.com/Frontware/GitLabBack/git"
	"github.com/Frontware/GitLabBack/gitlab"
)

// oauthHTTPURL returns url for open authorization
func oauthHTTPURL(httpURL, token string) string {
	if strings.HasPrefix(httpURL, "https") {
		return httpURL[0:8] + "oauth2:" + token + "@" + httpURL[8:]
	}
	return httpURL[0:7] + "oauth2:" + token + "@" + httpURL[7:]
}

func backup(conf *config.Config) {
	// init GitLab client
	c := gitlab.New(conf)

	// Request groups
	groups, err := c.ListGroups()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, group := range groups {
		// Request projects for a group
		projects, err := c.ListProjects(group.ID)
		if err != nil {
			log.Print(err.Error())
		}

		for _, project := range projects {
			var repoURL string

			if conf.SSH {
				repoURL = project.SSHURLToRepo
			} else {
				if len(conf.Token) > 0 {
					repoURL = oauthHTTPURL(project.HTTPURLToRepo, conf.Token)
				} else {
					repoURL = project.HTTPURLToRepo
				}
			}

			myRepo := git.Repo{
				Repo:     repoURL,
				Dir:      conf.BackupDir + project.PathWithNamespace,
				Name:     project.Name,
				FullName: project.NameWithNamespace,
			}

			if _, err := os.Stat(myRepo.Dir); err == nil {
				log.Printf("Pulling %s", myRepo.FullName)
				if _, err := myRepo.Pull(); err == nil {
					// Pull successfully
					continue
				}
				log.Printf("Error while pulling %s", myRepo.FullName)
				log.Printf("Removing %s", myRepo.FullName)
				if err := os.RemoveAll(myRepo.Dir); err != nil {
					log.Print(err.Error())
				}
			}

			log.Printf("Cloning %s", myRepo.FullName)
			if _, err := myRepo.Clone(); err != nil {
				log.Printf("Error while cloning %s", myRepo.FullName)
			}
		}
	}
}
