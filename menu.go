package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

// Menu related constants and data

type GameMode int

const (
	ModeMenu GameMode = iota
	ModePlaying
)

var menuItems = []string{
	"Start New Game",
	"Play vs Computer",
	"Play vs Human",
	"Quit",
}

type Menu struct {
	Selection int
}

func (m *Menu) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x20, 0x20, 0x20, 0xff}) // Dark background

	x, y := 100, 100
	for i, item := range menuItems {
		col := color.RGBA{0xff, 0xa5, 0x00, 0xff}
		if i == m.Selection {
			col = color.RGBA{0xff, 0x69, 0xb4, 0xff}
		}
		text.Draw(screen, item, basicfont.Face7x13, x, y+(i*24), col)
	}
}

const (
	menuStartX     = 100
	menuStartY     = 100
	menuItemHeight = 24 // adjust if you use a different font/spacing
)

func (m *Menu) Update(g *Game) error {
	// Handle keyboard as before
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		m.Selection = (m.Selection + 1) % len(menuItems)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		m.Selection--
		if m.Selection < 0 {
			m.Selection = len(menuItems) - 1
		}
	}

	// Mouse navigation
	mouseX, mouseY := ebiten.CursorPosition()
	for i := range menuItems {
		itemY := menuStartY + i*menuItemHeight
		if mouseX >= menuStartX && mouseX <= menuStartX+200 && // 200: menu width estimate
			mouseY >= itemY-12 && mouseY <= itemY+12 { // center text vertically
			m.Selection = i
			// Optionally, break here if you want only one item to be selectable at a time.
		}
	}
	// Activate menu item on Enter or mouse click
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		switch m.Selection {
		case 0:
			log.Println("Start New Game selected")
			g.StartNewGame()
			g.Mode = ModePlaying
		case 1:
			log.Println("Play vs Computer selected")
			g.StartComputerGame()
			g.Mode = ModePlaying
		case 2:
			log.Println("Play vs Human selected")
			g.StartHumanGame()
			g.Mode = ModePlaying
		case 3:
			log.Println("Quit selected")
			return ebiten.Termination
		}
	}
	return nil
}

// func (m *Menu) Update(g *Game) error {
// 	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
// 		m.Selection = (m.Selection + 1) % len(menuItems)
// 	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
// 		m.Selection--
// 		if m.Selection < 0 {
// 			m.Selection = len(menuItems) - 1
// 		}
// 	} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
// 		switch m.Selection {
// 		case 0:
// 			log.Println("Start New Game selected")
// 			g.StartNewGame()
// 			g.Mode = ModePlaying
// 		case 1:
// 			log.Println("Play vs Computer selected")
// 			g.StartComputerGame()
// 			g.Mode = ModePlaying
// 		case 2:
// 			log.Println("Play vs Human selected")
// 			g.StartHumanGame()
// 			g.Mode = ModePlaying
// 		case 3:
// 			log.Println("Quit selected")
// 			return ebiten.Termination
// 		}
// 	}
// 	return nil
// }
