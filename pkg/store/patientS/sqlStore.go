package patientS

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

func (s *sqlStore) GetOne(id int) (domain.Patient, error) {
	var patient domain.Patient
	query := "SELECT * FROM patients WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&patient.Id, &patient.Lastname, &patient.Name, &patient.Address, &patient.IdentityDocument, &patient.EntryDate)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (s *sqlStore) GetAll() ([]domain.Patient, error) {
	patients := []domain.Patient{}
	query := "SELECT id, lastname, name, address, identity_document, entry_date FROM patients;"
	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err.Error())
	}
	var patient domain.Patient
	for rows.Next() {
		err := rows.Scan(&patient.Id, &patient.Lastname, &patient.Name, &patient.Address, &patient.IdentityDocument, &patient.EntryDate)
		if err != nil {
			return patients, err
		}
		patients = append(patients, patient)
	}
	return patients, nil
}

func (s *sqlStore) Create(patient domain.Patient) error {
	query := "INSERT INTO patients (lastname, name, address, identity_document, entry_date) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	fmt.Println(patient)
	res, err := stmt.Exec(patient.Lastname, patient.Name, patient.Address, patient.IdentityDocument, patient.EntryDate)
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

func (s *sqlStore) UpdateOne(patient domain.Patient) error {
	query := "UPDATE patients SET lastname = ?, name = ?, address = ?, identity_document = ?, entry_date = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(patient.Lastname, patient.Name, patient.Address, patient.IdentityDocument, patient.EntryDate, patient.Id)
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
	query := "DELETE FROM patients WHERE id = ?;"
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

func (s *sqlStore) ValidateIdentityDocument(identity string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM patients WHERE identity_document = ?;"
	row := s.db.QueryRow(query, identity)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
