// pkg/service/person_service.go
package service

import (
	"context"
	"errors"

	"github.com/amartinezh/ms-crosslisting/pkg/model"
	"github.com/amartinezh/sdk-db/sdk_postgres"
)

type PersonService struct {
	db *sdk_postgres.PostgresSDK
}

func NewPersonService(db *sdk_postgres.PostgresSDK) *PersonService {
	return &PersonService{db: db}
}

// CreatePerson inserta una nueva persona en la base de datos
func (ps *PersonService) CreatePerson(person model.Person) error {
	query := "INSERT INTO dt.person (name) VALUES ($1) RETURNING id"
	row := ps.db.QueryRow(context.Background(), query, person.Name)
	err := row.Scan(&person.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetAllPersons obtiene todas las personas de la base de datos
func (ps *PersonService) GetAllPersons() ([]model.Person, error) {
	query := "SELECT id, name FROM dt.person"
	rows, err := ps.db.ExecuteQuery(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []model.Person
	for rows.Next() {
		var person model.Person
		if err := rows.Scan(&person.ID, &person.Name); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}

// UpdatePerson actualiza los datos de una persona en la base de datos
func (ps *PersonService) UpdatePerson(person model.Person) error {
	query := "UPDATE dt.person SET name = $1 WHERE id = $2"
	result, err := ps.db.Exec(context.Background(), query, person.Name, person.ID)
	if err != nil {
		return err
	}
	rowsAffected := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

// DeletePerson elimina una persona de la base de datos por ID
func (ps *PersonService) DeletePerson(id int) error {
	query := "DELETE FROM dt.person WHERE id = $1"
	result, err := ps.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	rowsAffected := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted")
	}
	return nil
}
