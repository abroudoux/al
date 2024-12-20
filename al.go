package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const configFileName string = "al.json"

func main() {
	if len(os.Args) > 1 {
		flagMode()
		os.Exit(0)
	} 

	println("No flag provided. Use -h or --help for help")
	os.Exit(0)
}

func flagMode() {
	flag := os.Args[1]

	if flag == "--init" || flag == "init" || flag == "-i" {
		err := initConfigFile()
		if err != nil {
			println(err)
			os.Exit(1)
		}
	} else if flag == "-v" || flag == "--version" {
		println("0.0.2")
	} else if flag == "-h" || flag == "--help" {
		printHelpManual()
	} else {
		err := runAlias()
		if err != nil {
			println(err)
			os.Exit(1)
		}
	}
}

func printHelpManual() {
	fmt.Println("Usage: al [options]")
	fmt.Printf("  %-20s %s\n", "al [alias]", "Run alias")
	fmt.Printf("  %-20s %s\n", "al [--init | -i]", "Initialize config file")
}

func fileExists() bool {
	info, err := os.Stat(configFileName)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func initConfigFile() error {
	if fileExists() {
		return fmt.Errorf("config file already exists")
	}

	println("Initializing al config file...")
	file, err := os.Create(configFileName)
	if err != nil {
		return fmt.Errorf("error creating config file: %w", err)
	}

	defer file.Close()
	return nil
}

func runAlias() error {
	if !fileExists() {
		println("Config file not found. Use al --init to create one")
		askUserToCreateConfigFile()
		os.Exit(1)
	}

	alias := ""

	for _, arg := range os.Args[1:] {
		alias += arg + " "
	}

	alias = alias[:len(alias)-1]
	command, err := findAliasInConfigFile(alias)
	if err != nil {
		return fmt.Errorf("error finding alias: %w", err)
	}

	runCommand(command)
	return nil
}

func findAliasInConfigFile(alias string) (string, error) {
	data, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return "", fmt.Errorf("error reading %s: %w", configFileName, err)
	}

	var aliases map[string]string
	err = json.Unmarshal(data, &aliases)
	if err != nil {
		return "", fmt.Errorf("error parsing JSON in %s: %w", configFileName, err)
	}

	if value, exists := aliases[alias]; exists {
		return value, nil
	}

    return "", fmt.Errorf("alias '%s' not found in %s", alias, configFileName)
}

func runCommand(command string) {
	parts := strings.Fields(command)

	if len(parts) == 0 {
		fmt.Println("Invalid command")
		os.Exit(1)
	}

	cmdName := parts[0]
	cmdArgs := parts[1:]
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Command %s executed successfully\n", command)
}

func askUserToCreateConfigFile() {
	println("Would you like to create a config file? (y/n)")
	var response string
	fmt.Scanln(&response)

	if response == "y" || response == "Y" || response == "yes" {
		initConfigFile()
	} else {
		println("Exiting...")
		os.Exit(0)
	}
}