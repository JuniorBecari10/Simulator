package main

import (
    "log"
    "image/color"
    _ "image/png"
    "image"
    
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
    spritesheet *ebiten.Image
)

type Game struct {}

type PaletteItem struct {
    x, y     int
    w, h     int
    itemType int
    sprite   *ebiten.Image
}

func (p *PalleteItem) Render(screen *ebiten.Image) {
    
}

func init() {
    paletteItems = make([]PaletteItem, itemsLength)
    
    var err error
    spritesheet, _, err = ebitenutil.NewImageFromFile("paletteitems.png")
    
    if err != nil {
        log.Fatal(err)
    }
    
    xinit := (windowWidth / 2) - ((itemsLength / 2) * itemSize)
    
    for i := 0; i < itemsLength; i++ {
        paletteItems[i] = PaletteItem { xinit + (itemSize * i), windowHeight - itemSize - 2, itemSize, itemSize, i, spritesheet.SubImage(image.Rect(i * 16, 0, 16, 16)).(*ebiten.Image) }
    }
}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Draw Palette Menu
    image := ebiten.NewImage(windowWidth, paletteHeight)
    image.Fill(color.RGBA{0, 48, 73, 255})
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
