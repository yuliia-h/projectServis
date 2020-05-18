package interfaces

import (
	"bytes"
	"encoding/base64"
	"github.com/disintegration/imaging"
	"projectServis/user_cases"
)

type LibraryImages struct {
}

func NewLibraryImages() *LibraryImages {
	return &LibraryImages{}
}

func (l LibraryImages) ResizeImageLibrary(image user_cases.Image) (user_cases.Image, error) {

	//Декодируем base64 в байты
	outPngData, err := base64.StdEncoding.DecodeString(image.Buffer)
	if err != nil {

		return user_cases.Image{}, err
	}
	//преобразовать байты в структуру image.Image

	im, err := imaging.Decode(bytes.NewReader(outPngData))
	if err != nil {
		return user_cases.Image{}, err
	}

	//var src = imaging.Resize(im, image.Width, image.Height, imaging.Lanczos)
	var src = imaging.Resize(im, 100, 0, imaging.Lanczos)

	var buff bytes.Buffer
	err = imaging.Encode(&buff, src, imaging.PNG)
	if err != nil {
		return user_cases.Image{}, err
	}

	image.Buffer = base64.StdEncoding.EncodeToString(buff.Bytes())
	image.Height = (100 * image.Height) / image.Width
	image.Width = 100

	//f, err := os.Create("a.png")
	//if err != nil {
	//	return user_cases.Image{}, err
	//}
	//
	//f.Write(buff.Bytes())
	//f.Close()

	return image, nil
}
