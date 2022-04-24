package ongkir

type Repository interface {
	Cost(Dataongkir *Ongkir) (*CostResponse, error)
	City(Dataongkir *Ongkir) (*Kota, error)
}

type Service interface {
	GetCity(Dataongkir *Ongkir) (*Kota, error)
	GetCost(Dataongkir *Ongkir) (*CostResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetCity(Dataongkir *Ongkir) (*Kota, error) {
	return s.repository.City(Dataongkir)
}

func (s *service) GetCost(Dataongkir *Ongkir) (*CostResponse, error) {
	return s.repository.Cost(Dataongkir)
}
