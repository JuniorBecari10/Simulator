package main

import (
    "log"
    "image/color"
    _ "image/png"
    "image"
    "math"
    //"fmt"
    
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
    windowWidth int = 960
    windowHeight int = 580
    
    itemsLength int = 3 // 6
    itemSize int = 48 // image 16
    paletteHeight int = itemSize + 16 + 4
    itemMargin int = 5
    
    blockSize int = 32 // scale 2
    blockTickSpeed int = 5
    
    typeStone int = 0
    typeWood int = 1
    typeSand int = 2
)

var (
    paletteItems [itemsLength]PaletteItem
    spritesheet *ebiten.Image
    blocks []Block
    
    clear Button
    eraser Button
    
    buttons []Button
    
    blockTickCount int = 0
    
    selected int
)

type Game struct {}

type Rectangle struct {
    x, y int
    w, h int
}

type PaletteItem struct {
    x, y     int
    itemType int
    sprite   *ebiten.Image
}

type Block struct {
    x, y int
    blockType int
    sprite *ebiten.Image
}

type Button struct {
    x, y int
    action func()
    sprite *ebiten.Image
}

func (p *PaletteItem) Tick() {
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && p.Hovered() {
        selected = p.itemType
    }
}

func (b *Block) Tick() {
    switch b.blockType {
        case typeSand:
            if !ThereIsBlock(b.x, b.y + blockSize) && b.y < windowHeight - paletteHeight - blockSize {
                b.y += blockSize
            } else if !ThereIsBlock(b.x + blockSize, b.y + blockSize) && b.y < windowHeight - paletteHeight - blockSize && b.x + blockSize < windowWidth {
                b.x += blockSize
                b.y += blockSize
            } else if !ThereIsBlock(b.x - blockSize, b.y + blockSize) && b.y < windowHeight - paletteHeight - blockSize && b.x > 0 {
                b.x -= blockSize
                b.y += blockSize
            }
            
            break
    }
}

func (p *PaletteItem) Render(screen *ebiten.Image) {
    op := &ebiten.DrawImageOptions{}
    
    if selected == p.itemType {
        bg := ebiten.NewImage(itemSize, itemSize)
        bg.Fill(color.RGBA{255, 255, 255, 255})
        op.GeoM.Translate(float64(p.x), float64(p.y))
        screen.DrawImage(bg, op)
    } else if p.Hovered() {
        bg := ebiten.NewImage(itemSize, itemSize)
        bg.Fill(color.RGBA{0, 61, 94, 255})
        op.GeoM.Translate(float64(p.x), float64(p.y))
        screen.DrawImage(bg, op)
    }
    
    op = &ebiten.DrawImageOptions{}
    
    op.GeoM.Scale(3, 3)
    op.GeoM.Translate(float64(p.x), float64(p.y))
    screen.DrawImage(p.sprite, op)
}

func (b *Block) Render(screen *ebiten.Image) {
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Scale(2, 2)
    op.GeoM.Translate(float64(b.x), float64(b.y))
    screen.DrawImage(b.sprite, op)
}

func (p *PaletteItem) GetRectangle() Rectangle {
    return Rectangle { p.x, p.y, itemSize, itemSize }
}

func (p *PaletteItem) Hovered() bool {
    mx, my := ebiten.CursorPosition()
    
    return Collide(p.GetRectangle(), Rectangle { mx, my, 1, 1 })
}

func Collide(r1 Rectangle, r2 Rectangle) bool {
    return r1.x < r2.x + r2.w &&
           r1.x + r1.w > r2.x &&
           r1.y < r2.y + r2.h &&
           r1.y + r1.h > r2.y
}

func ThereIsBlock(x, y int) bool {
    for _, v := range blocks {
        if v.x == x && v.y == y {
            return true
        }
    }
    
    return false
}

func init() {
    var err error
    spritesheet, _, err = ebitenutil.NewImageFromFile("sprites.png")
    
    blocks = make([]Block, 0)
    
    if err != nil {
        log.Fatal(err)
    }
    
    xinit := (windowWidth / 2) - ((itemsLength / 2) * (itemSize + itemMargin))
    
    for i := 0; i < itemsLength; i++ {
        paletteItems[i] = PaletteItem { xinit + ((itemSize + itemMargin) * i),
                                        (windowHeight - paletteHeight) + (itemSize / 4),
                                        i,
                                        spritesheet.SubImage(image.Rect(i * 16, 0, (i * 16) + 16, 16)).(*ebiten.Image) }
    }
    
    clear = Button { 10,
                     (windowHeight - paletteHeight) + (itemSize / 4),
                     func() {
                         blocks = make([]Block, 0)
                     },
                     spritesheet.SubImage(image.Rect(0, 32, 16, 48))
                   }
}

func (g *Game) Update() error {
    for _, v := range paletteItems {
        v.Tick()
    }
    
    blockTickCount++
    
    if blockTickCount >= blockTickSpeed {
        blockTickCount = 0
        
        for i := range blocks {
            blocks[i].Tick()
        }
    }
    
    mx, my := ebiten.CursorPosition()
    
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && my < windowHeight - paletteHeight {
        bx := int(math.Round(float64(mx / blockSize))) * blockSize
        by := int(math.Round(float64(my / blockSize))) * blockSize
        newblock := Block { bx, by, selected, spritesheet.SubImage(image.Rect(selected * 16, 16, (selected * 16) + 16, 32)).(*ebiten.Image)}
        
        if ThereIsBlock(bx, by) {
            return nil
        }
        
        blocks = append(blocks, newblock)
    }
    
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Draw Palette Menu
    image := ebiten.NewImage(windowWidth, paletteHeight)
    image.Fill(color.RGBA{0, 48, 73, 255})
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(0, float64(windowHeight - paletteHeight))
    
    screen.DrawImage(image, op)
    
    for _, v := range paletteItems {
        v.Render(screen)
    }
    
    for _, b := range blocks {
        b.Render(screen)
    }
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
