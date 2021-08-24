package backend

import (
	"fmt"
	"github.com/EtoDemerzel0427/ANSI-art/art"
	"github.com/disintegration/imaging"
	"github.com/gdamore/tcell/v2"
	"os"
)

type Terminal struct {
	screen tcell.Screen
	as *art.Solver
}



var (
	termStyle  = tcell.StyleDefault.Foreground(tcell.ColorFloralWhite).Background(tcell.NewRGBColor(0, 23, 31))
	//gridStyle = tcell.StyleDefault.Foreground(tcell.ColorYellow).Background(tcell.NewRGBColor(0, 23, 31)).Dim(true)
	gridStyle  = tcell.StyleDefault.Foreground(tcell.ColorDimGray).Background(tcell.NewRGBColor(0, 23, 31)).Dim(true)
)

func NewTerminal() *Terminal {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if e := screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	screen.SetStyle(termStyle)
	screen.Clear()

	as := art.NewSolver(200, 60, 20., 20., "MESSI", art.AsciiText)

	return &Terminal{
		screen: screen,
		as: as,
	}
}

func (t *Terminal) DrawGrid(content string) {
	pixel := 0
	for j := 0; j < 60; j++ {
		for i := 0; i < 200; i++ {
			if content[pixel] == '\n' {
				pixel++
			}
			t.screen.SetContent(i, j, rune(content[pixel]), nil, gridStyle)
			pixel++
		}
	}
	t.screen.Show()
}

func (t *Terminal) Debug() {
	img, _ := imaging.Open("./example.jpeg")
	img = t.as.TuneImage(img)

	t.DrawGrid(t.as.Convert(img))

}
