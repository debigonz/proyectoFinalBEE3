package dentistS

import (
	"database/sql"
	"examenFinal/internal/domain"
	"fmt"
	"log"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) GetOne(id int) (domain.Dentist, error) {
	var dentist domain.Dentist
	query := "SELECT * FROM dentists WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&dentist.Id, &dentist.Lastname, &dentist.Name, &dentist.License)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *sqlStore) GetAll() ([]domain.Dentist, error) {
	dentists := []domain.Dentist{}
	query := "SELECT id, lastname, name, license FROM dentists;"
	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err.Error())
	}
	var dentist domain.Dentist
	for rows.Next() {
		err := rows.Scan(&dentist.Id, &dentist.Lastname, &dentist.Name, &dentist.License)
		if err != nil {
			return dentists, err
		}
		dentists = append(dentists, dentist)
	}
	return dentists, nil
}

func (s *sqlStore) Create(dentist domain.Dentist) error {
	query := "INSERT INTO dentists (lastname, name, license) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	fmt.Println(dentist)
	res, err := stmt.Exec(dentist.Lastname, dentist.Name, dentist.License)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) UpdateOne(dentist domain.Dentist) error {
	query := "UPDATE dentists SET lastname = ?, name = ?, license = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentist.Lastname, dentist.Name, dentist.License, dentist.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DeleteOne(id int) error {
	query := "DELETE FROM dentists WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) ValidateLicense(license string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM dentists WHERE license = ?;"
	row := s.db.QueryRow(query, license)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
