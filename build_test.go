package main

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestBuildCommands(t *testing.T) {
	buildGit(t)
	var c Config
	c.getYaml()
	for _, p := range c.Projects {
		if p.Name == "nanohard" {
			build(&p)
			break
		}
	}
	// destroyGit(t)
}

func TestBuildTags(t *testing.T) {
	buildGit(t)
	var c Config
	c.getYaml()
	for _, p := range c.Projects {
		if p.Name == "robonano" {
			build(&p)
			break
		}
	}
	// destroyGit(t)
}

// func (c *Config) getYaml() *Config {
// 	yamlFile, err := ioutil.ReadFile("config.yml")
// 	if err != nil {
// 		log.Printf("yamlFile.Get err #%v ", err)
// 	}
// 	err = yaml.Unmarshal(yamlFile, c)
// 	if err != nil {
// 		log.Fatalf("Unmarshal: %v", err)
// 	}
// 	return c
// }



func buildGit(t *testing.T)  {
	// make test dir for git
	if _, err := os.Stat("./test"); os.IsNotExist(err) {
		if err := os.Mkdir("./test", 0777); err != nil {
			t.Fatal("os.Mkdir", err)
		}
	}

	// git init
	cmd := exec.Command("git", "init")
	cmd.Dir = "./test"
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal("git init", string(output), err)
	}

	// git set origin
	// remote set-url origin git@github.com:nanohard/scribble
	cmd = exec.Command("git", "remote", "add", "origin", "git@github.com:nanohard/scribble")
	cmd.Dir = "./test"
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatal("git remote add origin", string(output), err)
	}

	// // touch file... hehe
	// cmd = exec.Command("touch", "v1.0.txt")
	// cmd.Dir = "./test"
	// output, err = cmd.CombinedOutput()
	// if err != nil {
	// 	t.Fatal("touch test.txt", string(output), err)
	// }
	//
	// // git add
	// cmd = exec.Command("git", "add", ".")
	// cmd.Dir = "./test"
	// output, err = cmd.CombinedOutput()
	// if err != nil {
	// 	t.Fatal("git add", string(output), err)
	// }
	//
	// // git commit
	// cmd = exec.Command("git", "commit", "-am", "'commit 1'")
	// cmd.Dir = "./test"
	// output, err = cmd.CombinedOutput()
	// if err != nil {
	// 	t.Fatal("git commit", string(output), err)
	// }
	//
	// // git create tag
	// cmd = exec.Command("git", "tag", "v1.0")
	// cmd.Dir = "./test"
	// output, err = cmd.CombinedOutput()
	// if err != nil {
	// 	t.Fatal("git tag", string(output), err)
	// }
	//
	// // touch file... hehe
	// cmd = exec.Command("touch", "v2.0.txt")
	// cmd.Dir = "./test"
	// output, err = cmd.CombinedOutput()
	// if err != nil {
	// 	t.Fatal("touch test.txt", string(output), err)
	// }
	//
	// // git add
	// cmd = exec.Command("git", "add", ".")
	// cmd.Dir = "./test"
	// output, err = cmd.CombinedOutput()
	// if err != nil {
	// 	t.Fatal("git add", string(output), err)
	// }
	//
	// // git commit
	// cmd = exec.Command("git", "commit", "-am", "'commit 2'")
	// cmd.Dir = "./test"
	// output, err = cmd.CombinedOutput()
	// if err != nil {
	// 	t.Fatal("git commit", string(output), err)
	// }
	//
	// // git create annotated tag
	// cmd = exec.Command("git", "tag", "-a", "v2.0", "-m", "'annotated 2.0'")
	// cmd.Dir = "./test"
	// output, err = cmd.CombinedOutput()
	// if err != nil {
	// 	t.Fatal("git tag", string(output), err)
	// }
	//
	// // git checkout first tag
	// cmd = exec.Command("git", "checkout", "tags/v1.0")
	// cmd.Dir = "./test"
	// output, err = cmd.CombinedOutput()
	// if err != nil {
	// 	t.Fatal("git tag", string(output), err)
	// }
}

func destroyGit(t *testing.T) {
	// remove test folder
	cmd := exec.Command("rm", "-rf", "test/")
	// cmd.Dir = "./test"
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal("rm -rf test", string(output), err)
	}
	log.Println(string(output))
}
