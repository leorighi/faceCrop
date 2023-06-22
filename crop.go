package faceCrop

import (
	"fmt"
	"image"
	"io"

	"github.com/disintegration/imaging"
)

type BoundingBox struct {
	Height float64
	Left   float64
	Top    float64
	Width  float64
}

type CropLocalImageInput struct {
	ImagePath string
	Box       BoundingBox
}

type CropS3ImageInput struct {
	Body io.ReadCloser
	Box  BoundingBox
}

func CropLocalImage(input CropLocalImageInput) (image.Image, error) {
	img, err := imaging.Open(input.ImagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %w", err)
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	rect := image.Rect(
		int(float64(width)*input.Box.Left),
		int(float64(height)*input.Box.Top),
		int(float64(width)*(input.Box.Left+input.Box.Width)),
		int(float64(height)*(input.Box.Top+input.Box.Height)),
	)
	return imaging.Crop(img, rect), nil
}

func CropS3Image(input CropS3ImageInput) (image.Image, error) {
	img, err := imaging.Decode(input.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	rect := image.Rect(
		int(float64(width)*input.Box.Left),
		int(float64(height)*input.Box.Top),
		int(float64(width)*(input.Box.Left+input.Box.Width)),
		int(float64(height)*(input.Box.Top+input.Box.Height)),
	)
	return imaging.Crop(img, rect), nil
}
