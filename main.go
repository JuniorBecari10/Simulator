package main

import (
    "log"
    "image/color"
    
    "github.com/hajimehoshi/ebiten/v2"
)

const (
    windowWidth int = 740
    windowHeight int = 580
    
    itemsLength int = 6
    itemSize int = 48 // image 16
    paletteHeight int = itemSize + 4
    
    typeStone int = 0
    typeWood int = 1
    typeSand int = 2
)

var (
    paletteItems []PaletteItem
)

type Game struct {}

type PaletteItem struct {
    x, y     int
    w, h     int
    itemType int
    sprite   *ebiten.Image
}

func init() {
    paletteItems = make([]PaletteItem, itemsLength)
    
    xinit := (windowWidth / 2) - ((itemsLength / 2) * itemSize)
    
    for i := 0; i < itemsLength; i++ {
        paletteItems[i] = PaletteItem { xinit + (itemSize * i), windowHeight - itemSize - 2, itemSize, itemSize, i, nil }
    }
}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Draw Palette Menu
    image := ebiten.NewImage(windowWidth, paletteHeight)
    image.Fill(color.White)
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(0, float64(windowHeight - paletteHeight))
    
    screen.DrawImage(image, op)
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
