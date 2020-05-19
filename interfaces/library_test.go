package interfaces

import (
	"bytes"
	"errors"
	"image"
	"projectServis/user_cases"
	"testing"

	"github.com/disintegration/imaging"
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
				return nil, errors.New("DecodeString error")
			}},
			wantErr: errors.New("DecodeString error"),
			wantImg: user_cases.Image{},
		},

		{
			testName:  "Should return error if imaging.Decode returns error",
			incomeImg: user_cases.Image{},
			libImage: &LibraryImages{Decode: func([]byte) (image.Image, error) {
				return nil, errors.New("Decode error")
			}},
			wantErr: errors.New("Decode error"),
			wantImg: user_cases.Image{},
		},

		{
			testName:  "Should return error if imaging.Resize returns error",
			incomeImg: user_cases.Image{Buffer: "name"},
			libImage: &LibraryImages{Resize100: func(image.Image) image.Image {
				return nil
			}},
			wantErr: errors.New("imaging.Resize error"),
			wantImg: user_cases.Image{},
		},

		{
			testName:  "Should return error if imaging.Encode returns error",
			incomeImg: user_cases.Image{Buffer: "name"},
			libImage: &LibraryImages{Encode: func(*bytes.Buffer, image.Image, imaging.Format) error {
				return errors.New("Encode error")
			}},
			wantErr: errors.New("Encode error"),
			wantImg: user_cases.Image{},
		},

		{
			testName:  "Should return error if base64.StdEncoding.EncodeToString returns error",
			incomeImg: user_cases.Image{},
			libImage: &LibraryImages{EncodeToString: func(src []byte) string {
				return "EncodeToString error"
			}},
			wantErr: errors.New("EncodeToString error"),
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
