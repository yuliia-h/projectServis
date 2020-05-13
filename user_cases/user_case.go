package user_cases

type Image struct {
	Id     string `json:"id"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Buffer []byte `json:"buffer"`
}

type IService interface {
	Resize(image Image) error
	AddImage(image Image) error
	GetImages() ([]Image, error)
}

type Service struct {
}

func (service Service) Resize(image Image) error {
	return nil
}

func (service Service) AddImage(image Image) error {
	return nil
}

func (service Service) GetImages() ([]Image, error) {
	return nil, nil
}
