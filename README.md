# GitLabBack

[![Go Report Card](https://goreportcard.com/badge/github.com/Frontware/GitLabBack)](https://goreportcard.com/report/github.com/Frontware/GitLabBack)

**GitLabBack** is a tool to backup your git repositories from GitLab to your disk.

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
        Config path (default "~/.gitlabback")
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

Default configuration file is ".gitlabback".

- baseURL Link to GitLab server (default: https://gitlab.com/)
- token Private token
- backupdir Your backup directory
- ssh Clone with ssh (token is not needed)

```json
{"baseURL":"","token":"YOUR TOKEN","backupdir":"","ssh":false}
```

## Build

If you have Golang 1.11.x+ installed, you can clone the repository and:

```bash
$ go build
```

To ease the compilation process, the developer can use the provided Makefile.

```bash
$ make                    # builds for linux and windows
$ make linux              # builds for linux
$ make windows            # builds for windows
$ GOARCH=arm64 make linux # builds for arm64 linux
```

Start off from fresh build.

```bash
$ make clean # removes go binaries
```
