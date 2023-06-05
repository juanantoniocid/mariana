package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"juanantoniocid/mariana/internal/game"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth*game.SizeMultiplier, game.ScreenHeight*game.SizeMultiplier)
	ebiten.SetWindowTitle("Mariana (Zippo Studios)")

	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
