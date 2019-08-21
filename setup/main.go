package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/replicatedhq/krew-plugin-template/pkg/logger"
	"github.com/pkg/errors"
	"github.com/manifoldco/promptui"
)

type TemplateContext struct {
	Owner      string
	Repo       string
	PluginName string
}

func main() {
	owner, err := promptForOwner()
	if err != nil {
		fmt.Printf("%v\n", errors.Cause(err))
		os.Exit(1)
	}

	repo, err := promptForRepo()
	if err != nil {
		fmt.Printf("%v\n", errors.Cause(err))
		os.Exit(1)
	}

	pluginName, err := promptForPluginName()
	if err != nil {
		fmt.Printf("%v\n", errors.Cause(err))
		os.Exit(1)
	}

	templateContext := TemplateContext{
		Owner:      owner,
		Repo:       repo,
		PluginName: pluginName,
	}

	log := logger.NewLogger()
	log.info("Updating README")
	if err := renderReadme(templateContext); err != nil {
		fmt.Printf("%v\n", errors.Cause(err))
		os.Exit(1)
	}

	log.Info("Updating sample code with names")
	if err := renderTemplates(templateContext); err != nil {
		fmt.Printf("%v\n", errors.Cause(err))
		os.Exit(1)
	}

	log.Info("Updating go.mod")
	if err := renderGoMod(templateContext); err != nil {
		fmt.Printf("%v\n", errors.Cause(err))
		os.Exit(1)
	}

	fmt.Printf("%s/%s\n", owner, repo)
	fmt.Printf("%s\n", pluginName)
}

func promptForOwner() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "failed to get working dir")
	}

	pathParts := strings.Split(filepath.Dir(cwd), string(os.PathSeparator))

	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("Invalid")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "GitHub Organization or Username",
		Validate: validate,
		Default:  pathParts[len(pathParts)-2],
	}

	result, err := prompt.Run()
	if err != nil {
		return "", errors.Wrap(err, "failed to prompt for github owner")
	}

	return result, nil
}

func promptForRepo() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "failed to get working dir")
	}

	pathParts := strings.Split(filepath.Dir(cwd), string(os.PathSeparator))

	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("Invalid")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "GitHub Repo Name",
		Validate: validate,
		Default:  pathParts[len(pathParts)-1],
	}

	result, err := prompt.Run()
	if err != nil {
		return "", errors.Wrap(err, "failed to prompt for github repo")
	}

	return result, nil
}

func promptForPluginName() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "failed to get working dir")
	}

	pathParts := strings.Split(cwd, string(os.PathSeparator))

	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("Invalid")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Plugin Name",
		Validate: validate,
		Default:  pathParts[len(pathParts)-1],
	}

	result, err := prompt.Run()
	if err != nil {
		return "", errors.Wrap(err, "failed to prompt for plugin name")
	}

	return result, nil
}
