package shoe

type ShoeInterface interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

type Shoe struct {
	Logo string
	Size int
}

func (s *Shoe) setLogo(logo string) {
	s.Logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.Logo
}

func (s *Shoe) setSize(size int) {
	s.Size = size
}

func (s *Shoe) GetSize() int {
	return s.Size
}
