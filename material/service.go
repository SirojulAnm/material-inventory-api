package material

type Service interface {
	Add(input InputMaterial) (Material, error)
	Update(ID int, quantity int) (Material, error)
	GetAll() ([]Material, error)
	FindByID(ID int) (Material, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Add(input InputMaterial) (Material, error) {
	material := Material{}
	material.Name = input.Name
	material.Type = input.Type
	material.Quantity = input.Quantity

	newMaterial, err := s.repository.Save(material)
	if err != nil {
		return newMaterial, err
	}

	return newMaterial, nil
}

func (s *service) Update(ID int, quantity int) (Material, error) {
	newMaterial, err := s.repository.Update(ID, quantity)
	if err != nil {
		return newMaterial, err
	}

	return newMaterial, nil
}

func (s *service) GetAll() ([]Material, error) {
	var materials []Material
	materials, err := s.repository.Get()
	if err != nil {
		return materials, err
	}

	return materials, nil
}

func (s *service) FindByID(ID int) (Material, error) {
	var material Material
	material, err := s.repository.FindByID(ID)
	if err != nil {
		return material, err
	}

	return material, nil
}
