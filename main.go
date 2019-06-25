package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Projects []Project
}

// type project map[string][]Project
type Project struct {
	Name    string  `yaml:"name"`
	RepoURL string  `yaml:"repo_url"`
	Build   Build   `yaml:"build"`
}

// type options map[string]Options
type Build struct {
	// Command string `yaml:"command"`
	Commands []Command `yaml:"commands"`
	Tags    bool       `yaml:"tags"`
}

type Command struct {
	Command string `yaml:"command"`
}

func main() {
	var c Config
	c.getYaml()
	fmt.Println(c.Projects[0].Build.Commands[0].Command)
	// Init gorilla mux.
	r := mux.NewRouter()

	// Server options
	server := &http.Server{
		Addr: os.Getenv("HOST") + ":" + os.Getenv("PORT"),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	r.HandleFunc("/api/{project}", getProject)

	// Shutdown logic --------------------------------------------------------
	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		<-gracefulStop
		log.Println("Preparing to shut down...")

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		_ = server.Shutdown(ctx)
		log.Println("Exiting")
		os.Exit(0)
	}()
	// End Shutdown logic ---------------------------------------------------------

	log.Println("Web server starting")
	if os.Getenv("SSL") == "true" {
		// Run SSL server.
		if err := server.ListenAndServeTLS(
			os.Getenv("CERTFILE"), os.Getenv("KEYFILE")); err != nil {
				log.Println("server.ListenAndServeTLS():", err)
		}
	} else if os.Getenv("SSL") == "false" {
		if err := server.ListenAndServe(); err != nil {
			log.Println("server.ListenAndServe():", err)
		}
	}
}

func build(p *Project)  {
	// if p.Build.Tags {
	//
	// } else {
	// 	exec.Command()
	// }
}

func getProject(_ http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	project := vars["project"]

	var c Config
	c.getYaml()
	for _, p := range c.Projects {
		if p.Name == project {
			build(&p)
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
