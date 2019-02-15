// Package ascii implements an ascii art encoder from images.
package ascii

import (
	"bytes"
	"image"
	"image/color"

	"github.com/nfnt/resize"
)

var table = []byte("MND8OZ$7I?+=~:,..")

func Thumbnail(src image.Image, maxWidth, maxHeight uint) string {
	src = resize.Thumbnail(maxWidth, maxHeight, src, resize.Lanczos3)

	var buf bytes.Buffer
	bounds := src.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.GrayModel.Convert(src.At(x, y))
			y := c.(color.Gray).Y
			pos := int(int(y) * 16 / 255)
			buf.WriteByte(table[pos])
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
