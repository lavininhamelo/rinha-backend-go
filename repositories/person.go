package repositories

import (
	"rinha-basic/entities"
)

type PersonRepository interface {
	CreatePerson(p *entities.Person) (string, error)
	GetPerson(id string) (*entities.Person, error)
	SearchPeople(term string) ([]entities.Person, error)
	CountPeople() (int, error)
}
