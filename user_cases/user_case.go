package user_case

type Services interface {
	Resize(image Image) error
	AddImage(image Image) error
	GetImages() ([]Image, error)
}

type S struct {
}

func (s S) Resize(image Image) error {
	return nil
}

func (s S) AddImage(image Image) error {
	return nil
}

func (s S) GetImages() ([]Image, error) {
	return nil, nil
}
