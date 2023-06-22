package faceCrop

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"testing"
)

func TestCropLocalImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	c := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	for i := 0; i < img.Rect.Dx(); i++ {
		for j := 0; j < img.Rect.Dy(); j++ {
			img.Set(i, j, c)
		}
	}
	f, _ := os.Create("test_image.jpg")
	jpeg.Encode(f, img, nil)
	f.Close()

	box := BoundingBox{Height: 0.5, Left: 0.5, Top: 0.5, Width: 0.5}
	input := CropLocalImageInput{ImagePath: "test_image.jpg", Box: box}
	croppedImage, err := CropLocalImage(input)
	if err != nil {
		t.Errorf("Failed to crop image: %v", err)
	}
	if croppedImage.Bounds().Dx() != 50 || croppedImage.Bounds().Dy() != 50 {
		t.Errorf("The size of the cropped image is incorrect")
	}
}

func TestSaveImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	c := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	for i := 0; i < img.Rect.Dx(); i++ {
		for j := 0; j < img.Rect.Dy(); j++ {
			img.Set(i, j, c)
		}
	}
	err := SaveImage(img, "saved_image.jpg")
	if err != nil {
		t.Errorf("Failed to save image: %v", err)
	}
	if _, err := os.Stat("saved_image.jpg"); os.IsNotExist(err) {
		t.Errorf("The saved image file does not exist")
	}
	os.Remove("test_image.jpg")
	os.Remove("saved_image.jpg")
}
