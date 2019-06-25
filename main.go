package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
	// "github.com/ghodss/yaml"
)

type Config struct {
	Port     string
	Projects []Project `yaml:"projects"`
}

// type project map[string][]Project
type Project struct {
	Name    string
	RepoURL string    `yaml:"repo_url"`
	Build   Options `yaml:"build"`
}

// type options map[string]Options
type Options struct {
	Command string `yaml:"command"`
	Tags    bool   `yaml:"tags"`
}

// projects:
//   nanohard:
//     repo_url: git@github.com:nanohard/site-nanohard-main
//     build:
//       command: cd /var/www/nanohard.net && git reset --hard HEAD && git fetch
//   robonano:
//     repo_url: git@github.com:nanohard/robonano
//     build:
//       tags: true

// type Message struct {
// 	Environments map[string]models `yaml:"Environments"`
// }

// type models map[string][]Model

// type Model struct {
// 	AppType     string `yaml:"app-type"`
// 	ServiceType string `yaml:"service-type"`
// }

// Environments:
//  sys1:
//     models:
//     - app-type: app1
//       service-type: fds
//     - app-type: app2
//       service-type: era
//  sys2:
//     models:
//     - app-type: app1
//       service-type: fds
//     - app-type: app2
//       service-type: era

func (c *Config) getYaml() *Config {

	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c Config
	c.getYaml()
	fmt.Println(c.Projects[0].Build.Tags)
}
