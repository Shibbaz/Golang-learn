package main

type Rectangle struct {
	Width  float64
	Height float64
}

func Area(r Rectangle) float64 {
	return r.Width * r.Height
}