# FaceCrop

[![Go Report Card](https://goreportcard.com/badge/github.com/leorighi/faceCrop)](https://goreportcard.com/report/github.com/leorighi/faceCrop)

`faceCrop` is a GoLang package providing robust utilities for cropping faces (or any region) in images. It supports both local and Amazon S3 sources.

## Installation

To install this package, run:

```bash
go get github.com/leorighi/faceCrop
```

## Usage

Here's an example of how to use `faceCrop`:

```go
package main

import (
	"fmt"
	"github.com/leorighi/faceCrop"
)

func main() {
	box := faceCrop.BoundingBox{Height: 0.5, Left: 0.25, Top: 0.25, Width: 0.5}

	// Crop local image
	localInput := faceCrop.CropLocalImageInput{ImagePath: "path/to/your/image.jpg", Box: box}
	croppedImage, err := faceCrop.CropLocalImage(localInput)
	if err != nil {
		fmt.Println("Error cropping local image:", err)
	}

	// Save the cropped image
	err = faceCrop.SaveImage(croppedImage, "path/to/save/cropped/image.jpg")
	if err != nil {
		fmt.Println("Error saving cropped image:", err)
	}
}
```
## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)

