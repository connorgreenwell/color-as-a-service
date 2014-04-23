package main

import (
	"github.com/go-martini/martini"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"strconv"
)

func main() {
	m := martini.Classic()
	m.Get("/color/:hex.png", func(params martini.Params, w http.ResponseWriter, r *http.Request) {
		rect := image.Rect(0, 0, 50, 50)
		col := color.RGBA{255, 255, 255, 255}

		if len(params["hex"]) != 6 {
			return
		}

		red, err := strconv.ParseInt(params["hex"][0:2], 16, 16)
		green, err := strconv.ParseInt(params["hex"][2:4], 16, 16)
		blue, err := strconv.ParseInt(params["hex"][4:6], 16, 16)

		if err != nil {
			return
		}
		col = color.RGBA{uint8(red), uint8(green), uint8(blue), 255}

		s := r.URL.Query()
		if s["x"] != nil {
			x, err := strconv.Atoi(s["x"][0])
			if err != nil {
				x = 50
			}
			rect.Max.X = x
		}
		if s["y"] != nil {
			y, err := strconv.Atoi(s["y"][0])
			if err != nil {
				y = 50
			}
			rect.Max.Y = y
		}
		if rect.Max.X > 1000 || rect.Max.Y > 1000 || rect.Max.X <= 0 || rect.Max.Y <= 0 {
			return
		}

		img := image.NewRGBA(rect)
		for i := 0; i < rect.Max.X; i++ {
			for j := 0; j < rect.Max.Y; j++ {
				img.SetRGBA(i, j, col)
			}
		}
		png.Encode(w, img)
	})
	m.Run()
}
