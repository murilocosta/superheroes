package superhero

import (
	"errors"
	"strconv"
)

type Service interface {
	Create(name string) error
	Delete(id int64) error
	ListByType(superType SuperType) ([]*Super, error)
}

type serviceImpl struct {
	api  *SuperApi
	repo SuperRepository
}

func NewService(api *SuperApi, repo SuperRepository) Service {
	return &serviceImpl{api, repo}
}

func (s *serviceImpl) Create(name string) error {
	if name == "" {
		return errors.New("Super name is required")
	}

	resp, err := s.api.FindByName(name)
	if err != nil {
		return err
	}

	var superArray []*Super
	for _, r := range resp {
		hero := convertToSuper(r)
		superArray = append(superArray, hero)
	}
	s.repo.SaveMany(superArray)

	return nil
}

func (s *serviceImpl) Delete(id int64) error {
	if id <= 0 {
		return errors.New("Must suply a valid ID")
	}

	res, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if res == nil {
		return errors.New("Could not find the ID")
	}

	err = s.repo.Delete(res.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceImpl) ListByType(superType SuperType) ([]*Super, error) {
	return s.repo.List(superType)
}

func convertToSuper(r *RemoteSuper) *Super {
	var algn SuperType
	if r.Biography.Alignment == "good" {
		algn = HeroType
	} else {
		algn = VillainType
	}

	uuid, _ := strconv.ParseInt(r.ID, 10, 64)
	intelligence, _ := strconv.ParseInt(r.Powerstats.Intelligence, 10, 64)
	power, _ := strconv.ParseInt(r.Powerstats.Power, 10, 64)

	return &Super{
		Type:         algn,
		UUID:         uuid,
		Name:         r.Name,
		FullName:     r.Biography.FullName,
		Intelligence: intelligence,
		Power:        power,
		Occupation:   r.Work.Occupation,
		Image:        r.Image.URL,
	}
}
