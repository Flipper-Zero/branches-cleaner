package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var cfg config

func main() {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Config", err)
	}

	if !isExistingDir(cfg.BranchesPath) {
		log.Fatalln(cfg.BranchesPath, "is not an existing directory")
	}

	c := cron.New()
	c.AddFunc(cfg.Cron, checkBranches)
	c.Start()
	log.Println("First run:", c.Entries()[0].Next)

	select {}
}

func checkBranches() {
	br, err := getBranchesAndTagsList(cfg.GithubToken, cfg.RepoOwner, cfg.RepoName)
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
