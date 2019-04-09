package repl

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path"
	"strings"

	"bitbucket.org/pkg/inflect"
	"github.com/chzyer/readline"
)

// App is an instance of the REPL
type App struct {
	Name           string
	Handler        func(string)
	WelcomeMessage string
}

// Run starts the REPL
func (c *App) Run() {
	var err error
	var currentUser *user.User
	if currentUser, err = user.Current(); err != nil {
		panic(err.Error())
	}
	historyFilePath := path.Join(currentUser.HomeDir, fmt.Sprintf(".%v_history", inflect.Underscore(c.Name)))

	lineReader, err := readline.NewEx(&readline.Config{
		Prompt:            "> ",
		HistoryFile:       historyFilePath,
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
	})

	if err != nil {
		panic(err.Error())
	}

	log.SetOutput(lineReader.Stderr())

	println(c.WelcomeMessage)

	for {
		line, err := lineReader.Readline()
		if err != nil {
			if err == readline.ErrInterrupt {
			} else if err == io.EOF {
				os.Exit(0)
				println(err.Error())
			} else {
				println(err.Error())
			}
		} else {
			cleanedLine := strings.TrimSpace(line)
			c.Handler(cleanedLine)
		}
	}
}
