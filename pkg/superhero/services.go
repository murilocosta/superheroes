package superhero

import (
	"errors"
	"strconv"
)

type SuperHeroService interface {
	AddSuper(name string) error
	ListSuper(superType SuperType) ([]*Super, error)
}

type superHeroServiceImpl struct {
	api  *SuperHeroApi
	repo SuperRepository
}

func NewSuperHeroService(api *SuperHeroApi, repo SuperRepository) SuperHeroService {
	return &superHeroServiceImpl{api, repo}
}

func (s *superHeroServiceImpl) AddSuper(name string) error {
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

func (s *superHeroServiceImpl) ListSuper(superType SuperType) ([]*Super, error) {
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
		Image:        r.Image.Url,
	}
}
