package modules

import (
	"errors"
	"fmt"
	"gitlab.com/dentych/env/configuration"
	"log"
	"os/exec"
)

var gitAliases = map[string]string{
	"alias.cp":    "cherry-pick",
	"alias.co":    "checkout",
	"alias.cl":    "clone",
	"alias.c":     "commit",
	"alias.st":    "status -sb",
	"alias.br":    "branch",
	"alias.d":     "diff",
	"alias.dc":    "diff --cached",
	"alias.p":     "pull -p",
	"alias.pu":    "push -u",
	"alias.f":     "fetch -p",
	"alias.b":     "branch",
	"alias.logn":  "log --all --graph --oneline --decorate",
	"alias.lognb": "log --graph --oneline --decorate",
	"alias.pushb": "!git push -u origin $(git rev-parse --abbrev-ref HEAD)",
}

type git struct {
}

func NewGitModule() Module {
	return git{}
}

func (git) Name() string {
	return "git"
}

func (g git) Install() {
	fmt.Println("Installing...")
	if !gitExists() {
		log.Fatalln("git command not found. Please install git to use this module.")
	}

	fmt.Println("Setting user.name")
	executeCommand("user.name", configuration.Config.Git.Name)
	fmt.Println("Setting user.email")
	executeCommand("user.email", configuration.Config.Git.Email)

	fmt.Println("Setting up git aliases")
	for k, v := range gitAliases {
		fmt.Println("Setting", k)
		executeCommand(k, v)
	}

	fmt.Println("Done!")
}

func (git) Uninstall() {
	fmt.Println("Uninstalling...")
	fmt.Println("Unsetting user.name")
	executeCommand("--unset", "user.name")
	fmt.Println("Unsetting user.email")
	executeCommand("--unset", "user.email")

	for k := range gitAliases {
		fmt.Println("Unsetting", k)
		executeCommand("--unset", k)
	}
	fmt.Println("Uninstallation done!")
}

func gitExists() bool {
	cmd := exec.Command("git", "version")
	_, err := cmd.Output()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return false
		} else {
			log.Fatalln("Error executing command", err)
		}
	}
	return true
}

func executeCommand(args ...string) {
	fullArgs := []string{"config", "--global"}
	fullArgs = append(fullArgs, args...)
	cmd := exec.Command("git", fullArgs...)
	err := cmd.Run()
	if err != nil {
		if cmd.ProcessState.ExitCode() != 5 {
			log.Fatalf("Failed to run command '%s': %s", cmd.String(), err)
		}
	}
}
