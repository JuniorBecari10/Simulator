package main

import (
    "log"
    
    "github.com/hajimehoshi/ebiten/v2"
)

const (
    windowWidth int = 640
    windowHeight int = 480
    
    itemsLength int = 6
    itemSize float64 = 
    
    typeStone int = 1
)

var (
    stonePallete PalleteItem
)

type Game struct {}

type PalleteItem struct {
    x, y     float64
    w, h     float64
    itemType int
    sprite   *ebiten.Image
}

func init() {
    stonePallete = PalleteItem {  }
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
