package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	err = ttf.Init()
	if err != nil {
		return fmt.Errorf("could not initialize TTF: %v", err)
	}
	defer ttf.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(1600, 1200, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	defer w.Destroy()

	if err := drawTitle(r); err != nil {
		return fmt.Errorf("could not draw title: %v", err)
	}

	time.Sleep(1 * time.Second)

	s, err := newScene(r)
	if err != nil {
		return fmt.Errorf("could not draw background: %v", err)
	}
	if err = s.paint(r); err != nil {
		return fmt.Errorf("could not draw background: %v", err)
	}
	defer s.destroy()

	time.Sleep(5 * time.Second)

	return nil
}

func drawTitle(r *sdl.Renderer) error {
	r.Clear()
	f, err := ttf.OpenFont("/Users/rrode/go/src/gopdf/flappyGo/res/fonts/Flappy.ttf", 10)
	if err != nil {
		return fmt.Errorf("could not open font: %v", err)
	}
	defer f.Close()

	c := sdl.Color{R: 255, G: 100, B: 0, A: 255}
	s, err := f.RenderUTF8Solid("Flappy Gopher", c)
	if err != nil {
		return fmt.Errorf("could not render text: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("could not render texture: %v", err)
	}
	defer t.Destroy()
	if err = r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}
	r.Present()

	return nil
}
