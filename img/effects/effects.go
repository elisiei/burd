package effects

import (
	"burd/img/adjust"
	"burd/img/clone"
	"burd/img/parallel"
	"image"
	"image/color"
)

// Invert returns a negated version of the image.
func Invert(src image.Image) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {
		return color.RGBA{255 - c.R, 255 - c.G, 255 - c.B, c.A}
	}

	img := adjust.Apply(src, fn)

	return img
}

// Grayscale returns a copy of the image in Grayscale using the weights
// 0.3R + 0.6G + 0.1B as a heuristic.
func Grayscale(img image.Image) *image.RGBA {
	return GrayscaleWithWeights(img, 0.3, 0.6, 0.1)
}

// GrayscaleWithWeights returns a copy of the image in Grayscale using the given weights.
// The weights should be in the range 0.0 to 1.0 inclusive.
func GrayscaleWithWeights(img image.Image, r, g, b float64) *image.RGBA {
	src := clone.AsRGBA(img)
	bounds := src.Bounds()
	srcW, srcH := bounds.Dx(), bounds.Dy()

	if bounds.Empty() {
		return &image.RGBA{}
	}

	dst := image.NewRGBA(bounds)

	parallel.Line(srcH, func(start, end int) {
		for y := start; y < end; y++ {
			for x := range srcW {
				pos := y*src.Stride + x*4

				c := r*float64(src.Pix[pos+0]) + g*float64(src.Pix[pos+1]) + b*float64(src.Pix[pos+2])
				k := uint8(c + 0.5)
				dst.Pix[pos] = k
				dst.Pix[pos+1] = k
				dst.Pix[pos+2] = k
				dst.Pix[pos+3] = src.Pix[pos+3]
			}
		}
	})

	return dst
}
