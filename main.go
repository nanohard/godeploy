package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
	// "time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Projects []Project `yaml:"projects"`
}

// type project map[string][]Project
type Project struct {
	Name    string
	RepoURL string  `yaml:"repo_url"`
	Build   Options `yaml:"build"`
}

// type options map[string]Options
type Options struct {
	Command string `yaml:"command"`
	Tags    bool   `yaml:"tags"`
}

func main() {
	// Init gorilla mux.
	// r := mux.NewRouter()
	//
	// // Server options
	// server := &http.Server{
	// 	Addr: os.Getenv("HOST") + ":" + os.Getenv("PORT"),
	// 	// Good practice to set timeouts to avoid Slowloris attacks.
	// 	WriteTimeout: time.Second * 15,
	// 	ReadTimeout:  time.Second * 15,
	// 	IdleTimeout:  time.Second * 60,
	// 	Handler:      r, // Pass our instance of gorilla/mux in.
	// }
	//
	// r.HandleFunc("/api/{project}", build)
	//
	// log.Println("Web server starting")
	// if os.Getenv("SSL") == "true" {
	// 	// Run SSL server.
	// 	if err := server.ListenAndServeTLS(
	// 		os.Getenv("CERTFILE"), os.Getenv("KEYFILE")); err != nil {
	// 			log.Println("server.ListenAndServeTLS():", err)
	// 	}
	// } else if os.Getenv("SSL") == "false" {
	// 	if err := server.ListenAndServe(); err != nil {
	// 		log.Println("server.ListenAndServe():", err)
	// 	}
	// }

	var c Config
	c.getYaml()
	fmt.Println(c)
}

func build(_ http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	project := vars["project"]

	var c Config
	c.getYaml()
	for _, p := range c.Projects {
		if p.Name == project {

			break
		}
	}
	fmt.Println(c.Projects[0].Build.Tags)
}

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
