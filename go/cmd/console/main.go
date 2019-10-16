package main // import "github.com/simon-engledew/linuxkit-hashicorp-vault/go/cmd/console"

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/xdsopl/framebuffer/src/framebuffer"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"golang.org/x/sys/unix"
)

func main() {
	flag.Parse()
	text := flag.Arg(0)
	if text == "" {
		fmt.Fprintf(os.Stderr, "usage: console TEXT\n")
		os.Exit(1)
	}

	fd, err := os.Open("/dev/tty2")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file descriptor for /dev/tty2: %s\n", err)
		os.Exit(1)
	}

	if err := unix.IoctlSetInt(int(fd.Fd()), 0x4B3A, 0x01); err != nil {
		fmt.Fprintf(os.Stderr, "failed to set graphics mode: %s\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile("/sys/class/graphics/fbcon/cursor_blink", []byte("0"), 0600); err != nil {
		fmt.Fprintf(os.Stderr, "failed to stop cursor blink: %s\n", err)
		os.Exit(1)
	}

	fb, err := framebuffer.Open("/dev/fb0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open /dev/fb0: %s\n", err)
		os.Exit(1)
	}

	draw.Draw(fb, fb.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.ZP, draw.Src)

	mid := fb.Bounds().Size().Div(2).Sub(image.Point{X: len(text) * 7, Y: 13}.Div(2)).Mul(64)

	black := color.RGBA{0, 0, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(mid.X), fixed.Int26_6(mid.Y)}

	d := &font.Drawer{
		Dst:  fb,
		Src:  image.NewUniform(black),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)
}
