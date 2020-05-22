package interfaces

import (
	"bytes"
	"errors"
	"github.com/disintegration/imaging"
	"image"
	"projectServis/user_cases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	// arrange
	cases := []struct {
		testName  string
		incomeImg user_cases.Image
		libImage  *LibraryImages
		wantErr   error
		wantImg   user_cases.Image
	}{
		{
			testName:  "Should return error if base64.StdEncoding.DecodeString returns error",
			incomeImg: user_cases.Image{Buffer: "name"},
			libImage: &LibraryImages{DecodeString: func(s string) ([]byte, error) {
				return nil, errors.New("decodeString error")
			}},
			wantErr: errors.New("decodeString error"),
			wantImg: user_cases.Image{},
		},

		{
			testName:  "Should return error if imaging.Decode returns error",
			incomeImg: user_cases.Image{},
			libImage: &LibraryImages{
				DecodeString: func(s string) ([]byte, error) {
					return nil, nil
				},
				Decode: func([]byte) (image.Image, error) {
					return nil, errors.New("decode error")
				}},
			wantErr: errors.New("decode error"),
			wantImg: user_cases.Image{},
		},

		{
			testName:  "Should return error if imaging.Resize returns error",
			incomeImg: user_cases.Image{},
			libImage: &LibraryImages{
				DecodeString: func(s string) ([]byte, error) {
					return []byte(s), nil
				},
				Decode: func(s []byte) (image.Image, error) {
					return nil, nil
				},
				Resize: func(img image.Image) *image.NRGBA {
					return nil
				}},
			wantErr: errors.New("image is empty"),
			wantImg: user_cases.Image{},
		},

		{
			testName:  "Should return error if imaging.Encode returns error",
			incomeImg: user_cases.Image{},
			libImage: &LibraryImages{
				DecodeString: func(s string) ([]byte, error) {
					return []byte(s), nil
				},
				Decode: func(s []byte) (image.Image, error) {
					return nil, nil
				},
				Resize: func(image.Image) *image.NRGBA {
					var temp image.NRGBA
					return &temp
				},
				Encode: func(*bytes.Buffer, image.Image, imaging.Format) error {
					return errors.New("encode error")
				}},
			wantErr: errors.New("encode error"),
			wantImg: user_cases.Image{},
		},

		{
			testName:  "Should return error if base64.StdEncoding.EncodeToString returns error",
			incomeImg: user_cases.Image{},
			libImage: &LibraryImages{
				DecodeString: func(s string) ([]byte, error) {
					return nil, nil
				},
				Decode: func(s []byte) (image.Image, error) {
					return nil, nil
				},
				Resize: func(img image.Image) *image.NRGBA {
					var temp image.NRGBA
					return &temp
				},
				Encode: func(*bytes.Buffer, image.Image, imaging.Format) error {
					return nil
				},
				EncodeToString: func(src []byte) string {
					return string(src)
				}},
			wantErr: errors.New("encodeToString error"),
			wantImg: user_cases.Image{},
		},
	}
	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			gotImg, gotError := c.libImage.ResizeImageLibrary(c.incomeImg)
			assert.Equal(t, c.wantErr, gotError)
			assert.Equal(t, c.wantImg, gotImg)
		})
	}
}
