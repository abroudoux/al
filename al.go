package main

import (
	"os"
)

const fileName string = "al.json"

func main() {
	if len(os.Args) > 1 {
		flagMode()
	} 
}

func flagMode() {
	arg := os.Args[1]

	if arg == "init" || arg == "-i" {
		initConfigFile()
	} else if arg == "-v" || arg == "--version" {
		println("2.0.2")
	} else if arg == "-h" || arg == "--help" {
		printHelpManual()
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