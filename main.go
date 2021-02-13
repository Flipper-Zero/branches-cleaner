package main

import (
	"github.com/caarlos0/env/v6"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var cfg config

func main() {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Config", err)
	}

	if !isExistingDir(cfg.BranchesPath) {
		log.Fatalln(cfg.BranchesPath, "is not an existing directory")
	}

	for {
		checkBranches()
		time.Sleep(cfg.Delay)
	}
}

func checkBranches() {
	br, err := getBranchesList(cfg.GithubToken, cfg.RepoOwner, cfg.RepoName)
	if err != nil {
		log.Println("Can't get branches list from GitHub:", err)
		return
	}

	content, err := ioutil.ReadDir(cfg.BranchesPath)
	if err != nil {
		log.Println("Can't get content of", cfg.BranchesPath, err)
		return
	}

	for _, c := range content {
		if !c.IsDir() {
			continue
		}
		if arrayContains(br, c.Name()) {
			continue
		}
		log.Println("Deleting", c.Name())
		err = os.RemoveAll(filepath.Join(cfg.BranchesPath, c.Name()))
		if err != nil {
			log.Println("Can't delete", c.Name(), err)
		}
	}
}
