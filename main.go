package main

import (
    "log"
    
    "github.com/hajimehoshi/ebiten/v2"
)

const (
    windowWidth int = 640
    windowHeight int = 480
)

type Game struct {}

func init() {
    
}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    
}

func (g *Game) Layout(ow, oh int) (w, h int) {
    return windowWidth, windowHeight
}

func main() {
    ebiten.SetWindowSize(windowWidth, windowHeight)
    ebiten.SetWindowTitle("Simulator")
    
    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}
