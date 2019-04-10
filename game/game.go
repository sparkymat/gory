package game

import (
	"time"
)

// Game represents a game instance
type Game struct {
}

// New creates a new game instance
func New() *Game {
	var g Game

	return &g
}

// Start the game server
func (g *Game) Start() {
	ticker := time.NewTicker(time.Millisecond * 1000)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		}
	}
}
