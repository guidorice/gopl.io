/* Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 *
 * Copyright (c) 2013 Lucas Beyer
 * License: MIT
 * https://github.com/lucasb-eyer/go-colorful/tree/master/doc/gradientgen
 */

// ex 3.5 Implement a full-color Mandelbrot set using the function image.NewRGBA
// and the type color.RGBA or color.YCbCr.

package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/lucasb-eyer/go-colorful"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2048, 2048
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			c := mandelbrot(z)
			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// sqrt(c2+c2)
func Magnitude(n complex128) float64 {
	r := real(n)
	i := imag(n)
	vector := math.Sqrt(math.Pow(r, 2) + math.Pow(i, 2))
	return math.Abs(vector)
}

const maxIterations = 18
const escapeRadius = 3.0

var log2 = math.Log(2)

// mandelbrot function with smoothed escape
// http://linas.org/art-gallery/escape/escape.html
func mandelbrot(c complex128) color.Color {
	var z complex128
	var i uint8
	var modulus float64

	for i = uint8(0); i < maxIterations; i++ {
		z = z*z + c // the mandelbrot equation
		modulus = Magnitude(z)
		if modulus > escapeRadius {
			break
		}
	}
	z = z*z + c // the mandelbrot equation (couple of extra iterations helps)
	z = z*z + c
	i += 2
	modulus = Magnitude(z)
	mu := float64(i) + 1 - math.Log(math.Log(modulus))/log2
	if math.IsNaN(mu) || math.IsInf(mu, 1) || math.IsInf(mu, -1) {
		return color.Black
	}
	normalizedMu := math.Min(1, mu/float64(maxIterations))
	return gradient.GetInterpolatedColorFor(normalizedMu)
}

// The "keypoints" of the gradient.
var gradient = GradientTable{
	{MustParseHex("#9e0142"), 0.0},
	{MustParseHex("#d53e4f"), 0.1},
	{MustParseHex("#f46d43"), 0.2},
	{MustParseHex("#fdae61"), 0.3},
	{MustParseHex("#fee090"), 0.4},
	{MustParseHex("#ffffbf"), 0.5},
	{MustParseHex("#e6f598"), 0.6},
	{MustParseHex("#abdda4"), 0.7},
	{MustParseHex("#66c2a5"), 0.8},
	{MustParseHex("#3288bd"), 0.9},
	{MustParseHex("#5e4fa2"), 1.0},
}

// This table contains the "keypoints" of the colorgradient you want to generate.
// The position of each keypoint has to live in the range [0,1]
type GradientTable []struct {
	Col colorful.Color
	Pos float64
}

// This is the meat of the gradient computation. It returns a HCL-blend between
// the two colors around `t`.
// Note: It relies heavily on the fact that the gradient keypoints are sorted.
func (self GradientTable) GetInterpolatedColorFor(t float64) colorful.Color {
	for i := 0; i < len(self)-1; i++ {
		c1 := self[i]
		c2 := self[i+1]
		if c1.Pos <= t && t <= c2.Pos {
			// We are in between c1 and c2. Go blend them!
			t := (t - c1.Pos) / (c2.Pos - c1.Pos)
			return c1.Col.BlendHcl(c2.Col, t).Clamped()
		}
	}
	// Nothing found? Means we're at (or past) the last gradient keypoint.
	// return self[len(self)-1].Col
	c, _ := colorful.Hex("#000000")
	return c
}

// This is a very nice thing Golang forces you to do!
// It is necessary so that we can write out the literal of the colortable below.
func MustParseHex(s string) colorful.Color {
	c, err := colorful.Hex(s)
	if err != nil {
		panic("MustParseHex: " + err.Error())
	}
	return c
}
