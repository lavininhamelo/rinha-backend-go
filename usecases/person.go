package usecases

import (
	"fmt"
	"github.com/google/uuid"
	"rinha-basic/cache"
	"rinha-basic/entities"
	"rinha-basic/repositories"
)

type PersonUsecase interface {
	CreatePerson(p *entities.Person) (string, error)
	GetPerson(id string) (*entities.Person, error)
	SearchPeople(term string) ([]entities.Person, error)
	CountPeople() (int, error)
}

type personUsecase struct {
	repo  repositories.PersonRepository
	cache cache.CacheStorage
}

func NewPersonUsecase(repo repositories.PersonRepository, cache cache.CacheStorage) PersonUsecase {
	return &personUsecase{repo: repo, cache: cache}
}

func (u *personUsecase) CreatePerson(p *entities.Person) (string, error) {
	exists, err := u.cache.VerifyUsername(p.Nickname)
	if err != nil {
		return "", err
	}
	if exists {
		return "", fmt.Errorf("username %s already exists", p.Nickname)
	}

	id := uuid.New().String()
	p.ID = id

	_, err = u.repo.CreatePerson(p)

	if err != nil {
		return "", err
	}

	_, err = u.cache.SaveUsername(p.Nickname)
	_, err = u.cache.SavePerson(p.ID, p)
	return id, nil
}

func (u *personUsecase) GetPerson(id string) (*entities.Person, error) {
	person, _ := u.cache.RecoveryPerson(id)

	if person != nil {
		return person, nil
	}

	p, err := u.repo.GetPerson(id)
	if err != nil {
		return nil, err
	}

	u.cache.SavePerson(p.ID, p)

	return p, nil
}

func (u *personUsecase) SearchPeople(term string) ([]entities.Person, error) {
	people, err := u.repo.SearchPeople(term)
	if err != nil {
		return nil, err
	}
	if people == nil {
		return []entities.Person{}, nil
	}
	return people, nil
}

func (u *personUsecase) CountPeople() (int, error) {
	return u.repo.CountPeople()
}
