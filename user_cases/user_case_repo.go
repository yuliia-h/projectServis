package user_cases

//------------Repository----------------------------------

//интерфейсный тип
type Repositories struct {
	repositorfield RepositoryImager
}

//конструктор
func NewRepositories(repo RepositoryImager) *Repositories {
	return &Repositories{
		repositorfield: repo,
	}
}

type RepositoryImager interface {
	FindImageId(image Image) error
	GetAllImage(image Image) error
	ChangeImageId(image Image) error
	SaveImage(image Image) error
}

type RepositoryImages struct {
}

func (repo RepositoryImages) FindImageId(image Image) error {
	return nil
}

func (repo RepositoryImages) GetAllImage(image Image) error {
	return nil
}

func (repo RepositoryImages) ChangeImageId(image Image) error {
	return nil
}

func (repo RepositoryImages) SaveImage(image Image) error {
	return nil
}
