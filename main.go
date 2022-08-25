package main

import (
    "log"
    
    "github.com/hajimehoshi/ebiten/v2"
)

const (
    windowWidth int = 740
    windowHeight int = 580
    
    itemsLength int = 6
    itemSize int = 16
    
    typeStone int = 1
)

var (
    //stonePallete PalleteItem
    palleteItems []PalleteItem
)

type Game struct {}

type PalleteItem struct {
    x, y     int
    w, h     int
    itemType int
    sprite   *ebiten.Image
}

func init() {
    palleteItems = make([]PalleteItem, itemsLength)
    
    xinit := (windowWidth / 2) - ((itemsLength / 2) * itemSize)
    
    for i := 0; i < itemsLength; i++ {
        palleteItems[i] = PalleteItem { xinit + (itemSize * i), windowHeight - itemSize - 5, itemSize, itemSize, i, nil }
    }
}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Draw Pallete Menu
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
