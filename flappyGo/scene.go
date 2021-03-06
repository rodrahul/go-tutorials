package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	bg *sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene, error) {

	t, err := img.LoadTexture(r, "/Users/rrode/go/src/gopdf/flappyGo/res/imgs/background.png")
	if err != nil {
		return nil, fmt.Errorf("could not open background: %v", err)
	}

	return &scene{t}, nil
}

func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()
	if err := r.Copy(s.bg, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}
	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
}
