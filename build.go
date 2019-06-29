package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func buildTags(p *Project) {
	// exec pre commands
	for _, c := range p.Build.PreCommands {
		s := strings.Split(c.Command, " ")
		cmd := exec.Command(s[0], s[1:]...)
		cmd.Dir = p.Build.WorkingDir
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("buildTags() PreCommands", string(output), err)
		}
	}

	// clear any local changes
	cmd := exec.Command("git", "reset", "--hard")
	cmd.Dir = p.Build.WorkingDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("buildTags() git reset", string(output), err)
	}

	// git fetch
	cmd = exec.Command("git", "fetch")
	cmd.Dir = p.Build.WorkingDir
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Println("buildTags() git fetch", string(output), err)
	}

	// get current tag
	cmd = exec.Command("git", "describe", "--tags")
	cmd.Dir = p.Build.WorkingDir
	currentTagB, err := cmd.CombinedOutput()
	currentTag := string(currentTagB)
	currentTag = strings.TrimSuffix(currentTag, "\n")
	if err != nil {
		log.Println("buildTags() currentTag", currentTag, err)
	}

	// get latest tag
	cmd = exec.Command("bash", "-c", "git describe --tags `git rev-list --tags --max-count=1`")
	cmd.Dir = p.Build.WorkingDir
	latestTagB, err := cmd.CombinedOutput()
	latestTag := string(latestTagB)
	latestTag = strings.TrimSuffix(latestTag, "\n")
	if err != nil {
		log.Println("buildTags() latestTag", latestTag, err)
	}

	// checkout latest tag
	cmd = exec.Command("git", "checkout", "tags/"+fmt.Sprintf("%s", latestTag))
	cmd.Dir = p.Build.WorkingDir
	checkoutTag, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("buildTags() checkoutTag", string(checkoutTag), err)
	}

	// exec post commands
	for _, c := range p.Build.PostCommands {
		s := strings.Split(c.Command, " ")
		cmd := exec.Command(s[0], s[1:]...)
		cmd.Dir = p.Build.WorkingDir
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("buildTags() PostCommands", string(output), err)
		}
	}
}

func buildCommands(p *Project) {
	for _, c := range p.Build.Commands {
		s := strings.Split(c.Command, " ")
		cmd := exec.Command(s[0], s[1:]...)
		cmd.Dir = p.Build.WorkingDir
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("buildTags() checkoutTag", string(output), err)
		}
	}
}

func build(p *Project) {
	if p.Build.Type == "tags" {
		buildTags(p)
	} else if p.Build.Type == "commands" {
		buildCommands(p)
	}
}
