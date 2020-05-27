package interfaces

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"projectServis/user_cases"

	"github.com/disintegration/imaging"
)

type LibraryImages struct {
	DecodeString   func(string) ([]byte, error)
	Decode         func([]byte) (image.Image, error)
	Resize         func(image.Image) *image.NRGBA
	Encode         func(*bytes.Buffer, image.Image, imaging.Format) error
	EncodeToString func(src []byte) string
	SomeImage      func(i interface{}) image.Image
}

func NewLibraryImages() *LibraryImages {
	return &LibraryImages{
		DecodeString: func(s string) ([]byte, error) {
			return base64.StdEncoding.DecodeString(s)
		},
		Decode: func(s []byte) (image.Image, error) {
			return imaging.Decode(bytes.NewReader(s))
		},
		Resize: func(img image.Image) *image.NRGBA {
			return imaging.Resize(img, 100, 0, imaging.Lanczos)
		},
		Encode: func(s *bytes.Buffer, img image.Image, format imaging.Format) error {
			return imaging.Encode(s, img, imaging.PNG)
		},
		EncodeToString: func(src []byte) string {
			return base64.StdEncoding.EncodeToString(src)
		},
	}
}

func (l LibraryImages) ResizeImageLibrary(image user_cases.Image) (user_cases.Image, error) {

	//Декодируем base64 в байты
	outPngData, err := l.DecodeString(image.Buffer)
	if err != nil {
		return user_cases.Image{}, err
	}

	//преобразовать байты в структуру image.Image
	im, err := l.Decode(outPngData)
	if err != nil {
		return user_cases.Image{}, err
	}

	var src = l.Resize(im)

	if src == nil {
		return user_cases.Image{}, errors.New("image is empty")
	}

	var buff bytes.Buffer
	err = l.Encode(&buff, src, imaging.PNG)
	if err != nil {
		return user_cases.Image{}, err
	}

	image.Buffer = l.EncodeToString(buff.Bytes())
	if image.Buffer == "" {
		return user_cases.Image{}, errors.New("encodeToString error")
	}

	return image, nil
}
