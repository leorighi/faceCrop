package faceCrop

import (
	"image"

	"github.com/disintegration/imaging"
)

func SaveImage(img image.Image, outputPath string) error {
	return imaging.Save(img, outputPath)
}
