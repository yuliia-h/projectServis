package user_cases

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
	"github.com/noelyahan/impexp"
	"github.com/noelyahan/mergi"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

//-----------Library-----------------------------------------

type LibraryImager interface {
	ResizeImageLibrary(image Image) error
}

type LibraryImages struct {
}

func (library LibraryImages) ResizeImageLibrary(image Image) error {
	return nil
}

//---------library Mergi----------------
func re() {
	img, err := mergi.Import(impexp.NewFileImporter("testdata/coffee-171653_960_720.jpg"))
	if err != nil {
		fmt.Println(err)
		return
	}
	scale := 2

	// lets scale up
	newWidth := uint(img.Bounds().Max.X * scale)
	newHeight := uint(img.Bounds().Max.Y * scale)
	res, err := mergi.Resize(img, newWidth, newHeight)
	if err != nil {
		log.Fatal(err)
	}
	mergi.Export(impexp.NewFileExporter(res, getPath("scale_up.png")))

	// lets scale down
	newWidth = uint(img.Bounds().Max.X / scale)
	newHeight = uint(img.Bounds().Max.Y / scale)
	res, err = mergi.Resize(img, newWidth, newHeight)
	if err != nil {
		log.Fatal(err)
	}
	mergi.Export(impexp.NewFileExporter(res, getPath("scale_down.png")))

}

func getPath(name string) string {
	return fmt.Sprintf("examples/resize/res/%s", name)
}

//---------library Resize----------------
func resizeResize() {
	// open "test.jpg"
	file, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(1000, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}

//---------library Imaging----------------
func resizeImaging() {
	// Open a test image.
	src, err := imaging.Open("testdata/flowers.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Crop the original image to 300x300px size using the center anchor.
	src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	// Resize the cropped image to width = 200px preserving the aspect ratio.
	src = imaging.Resize(src, 200, 0, imaging.Lanczos)

	// Create a blurred version of the image.
	img1 := imaging.Blur(src, 5)

	// Create a grayscale version of the image with higher contrast and sharpness.
	img2 := imaging.Grayscale(src)
	img2 = imaging.AdjustContrast(img2, 20)
	img2 = imaging.Sharpen(img2, 2)

	// Create an inverted version of the image.
	img3 := imaging.Invert(src)

	// Create an embossed version of the image using a convolution filter.
	img4 := imaging.Convolve3x3(
		src,
		[9]float64{
			-1, -1, 0,
			-1, 1, 1,
			0, 1, 1,
		},
		nil,
	)

	// Create a new image and paste the four produced images into it.
	dst := imaging.New(400, 400, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img1, image.Pt(0, 0))
	dst = imaging.Paste(dst, img2, image.Pt(0, 200))
	dst = imaging.Paste(dst, img3, image.Pt(200, 0))
	dst = imaging.Paste(dst, img4, image.Pt(200, 200))

	// Save the resulting image as JPEG.
	err = imaging.Save(dst, "testdata/out_example.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
