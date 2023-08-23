package cache

import "rinha-basic/entities"

type PersonCache interface {
	SaveUsername(username string) (string, error)
	VerifyUsername(username string) (bool, error)
	SavePerson(id string, person *entities.Person) (bool, error)
	RecoveryPerson(id string) (*entities.Person, error)
}
