package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const fileName string = "al.json"

func main() {
	if len(os.Args) > 1 {
		flagMode()
	} 

	println("No flag provided. Use -h or --help for help")
}

func flagMode() {
	flag := os.Args[1]

	if flag == "init" || flag == "-i" {
		initConfigFile()
	} else if flag == "-v" || flag == "--version" {
		println("0.0.1")
	} else if flag == "-h" || flag == "--help" {
		printHelpManual()
	} else {
		runAlias()
	}

	os.Exit(0)
}

func printHelpManual() {
	println("Usage: branch [options]")
	println("Options:")
	println("al [--init | -i]  Initialize config file")

	os.Exit(0)
}

func fileExists() bool {
	info, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func initConfigFile() {
	if fileExists() {
		println("File already exists")
		os.Exit(0)
	}

	println("Initializing al config file...")
	file, err := os.Create(fileName)

	if err != nil {
		println("Error creating file")
		os.Exit(1)
	}

	defer file.Close()
}

func runAlias() {
	if !fileExists() {
		println("Config file not found. Use al --init to create one")
		os.Exit(1)
	}

	alias := ""

	for _, arg := range os.Args[1:] {
		alias += arg + " "
	}

	alias = alias[:len(alias)-1]
	value := findAliasInConfigFile(alias)
	println(value)
}

func findAliasInConfigFile(alias string) string {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading al.json:", err)
		os.Exit(1)
	}

	var aliases map[string]string
	err = json.Unmarshal(data, &aliases)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	if value, exists := aliases[alias]; exists {
		return value
	}

	return "Alias not found"
}