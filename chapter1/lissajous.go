// Lissajous генерирует анимированный GIF из
// случайных фигур Лисажу.
package main

import(
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
    "time"
)

var palette = []color.Color{color.White, color.RGBA{0xFF, 0x00, 0x00, 0xFF}, color.RGBA{0x00, 0xFF, 0x00, 0xFF}, color.RGBA{0x00, 0x00, 0xFF, 0xFF}}

const (
    whiteIndex = 0 // Первый цвет палитры
    blackIndex = 1 // Следующий цвет палитры
    redIndex = 2
    greenIndex = 3
    blueIndex = 4
)

func main() {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles = 5
        res = 0.001
        size = 100
        nframes = 64
        delay = 8
    )
    
    rand.Seed(time.Now().UTC().UnixNano())
    freq := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles * 2 * math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t * freq + phase)
            if(os.Args[2] == "red") {
                img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), redIndex)
            }
            if(os.Args[2] == "green") {
                img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), greenIndex)
            }
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}