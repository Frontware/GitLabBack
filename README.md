# GitLabBack

[![Go Report Card](https://goreportcard.com/badge/github.com/Frontware/GitLabBack)](https://goreportcard.com/report/github.com/Frontware/GitLabBack)

**GitLabBack** is tool to backup your git repositories from GitLab.

**GitLabBack** creates a backup by cloning repositories from GitLab. If the backup already exists, **GitLabBack** will pull changes instead of cloning the repositories. However, if git fails to pull new changes from GitLab, **GitLabBack** will delete the backup and clone it again.

## Usage

```bash
$ ./GitLabBack
```

## Help

```bash
$ ./GitLabBack -h
Usage of GitLabBack:
  -config string
        Config path (default ".GitLabBack.json")
  -dir string
        Backup directory
  -ssh
        Enable ssh clone
  -token string
        GitLab token
  -url string
        GitLab server base url
```

## Configuration

Default configuration file is ".GitLabBack.json".

```json
{"baseURL":"","token":"YOUR TOKEN","backupdir":"","ssh":false}
```

## Build

If you have Golang 1.11.x+ installed, you can clone the repository and:

```bash
$ go build
```
