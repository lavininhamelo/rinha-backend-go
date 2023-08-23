package repositories

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"rinha-basic/entities"
)

type PersonRepositoryImpl struct {
	db *sqlx.DB
}

func NewPersonRepository(db *sqlx.DB) PersonRepository {
	return &PersonRepositoryImpl{db: db}
}

func (r *PersonRepositoryImpl) CreatePerson(p *entities.Person) (string, error) {
	query := `INSERT INTO persons(nickname, name, birthday, stack, id) VALUES (?, ?, ?, ?, ?)`

	stackSerialized, err := json.Marshal(p.Stack)
	if err != nil {
		return "", err
	}

	_, err = r.db.Exec(query, p.Nickname, p.Name, p.Birthday, stackSerialized, p.ID)
	if err != nil {
		return "", err
	}

	return p.ID, nil
}

func (r *PersonRepositoryImpl) SearchPeople(term string) ([]entities.Person, error) {
	var people []entities.Person
	query := `
        SELECT DISTINCT p.id, p.nickname, p.name, p.birthday, p.stack
        FROM persons p
        WHERE p.name = ? OR p.nickname = ? OR p.stack LIKE ? 
        LIMIT 50
    `
	rows, err := r.db.Query(query, term, term, "%"+term+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var person entities.Person
		var stackSerialized string
		err = rows.Scan(&person.ID, &person.Nickname, &person.Name, &person.Birthday, &stackSerialized)
		if err != nil {
			return nil, err
		}

		var stack []string
		err = json.Unmarshal([]byte(stackSerialized), &stack)
		if err != nil {
			return nil, err
		}
		person.Stack = stack

		people = append(people, person)
	}

	return people, nil
}

func (r *PersonRepositoryImpl) GetPerson(id string) (*entities.Person, error) {
	p := &entities.Person{}
	var stackSerialized string

	query := "SELECT id, nickname, name, birthday, stack FROM persons WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Nickname, &p.Name, &p.Birthday, &stackSerialized)
	if err != nil {
		return nil, err
	}

	var stack []string
	err = json.Unmarshal([]byte(stackSerialized), &stack)
	if err != nil {
		return nil, err
	}
	p.Stack = stack

	return p, nil
}

func (r *PersonRepositoryImpl) CountPeople() (int, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM persons")
	return count, err
}
