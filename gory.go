package main

import (
	"io/ioutil"
	"os"

	"github.com/sparkymat/gory/game"
	"github.com/sparkymat/gory/repl"
)

func inputHandler(line string) {
	switch {
	case line == "help":
		displayFile("help.txt")
	case line == "motd":
		displayFile("motd.txt")
	case line == "quit":
		os.Exit(0)
	}
}

func displayFile(filename string) {
	var bytes []byte
	var err error

	if bytes, err = ioutil.ReadFile(filename); err == nil {
		println(string(bytes))
	} else {
		println("Error: Unable to read help.txt")
	}
}

func main() {
	var err error

	motd := "Welcome to Gory!"

	if _, err = os.Stat("motd.txt"); err == nil {
		var motdBytes []byte
		if motdBytes, err = ioutil.ReadFile("motd.txt"); err == nil {
			motd = string(motdBytes)
		}
	}

	inputChannel := make(chan string)

	go func() {
		for {
			inputHandler(<-inputChannel)
		}
	}()

	g := game.New()
	go g.Start()

	console := repl.App{
		Name:           "Gory",
		Channel:        inputChannel,
		WelcomeMessage: motd,
	}
	console.Run()
}
