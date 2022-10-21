package appoinmentS

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

func (s *sqlStore) GetOne(id int) (domain.Appointment, error) {
	var appoinment domain.Appointment
	query := "SELECT * FROM appointments WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&appoinment.Id, &appoinment.Dentist, &appoinment.Patient, &appoinment.Date, &appoinment.Time, &appoinment.Description)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appoinment, nil
}

func (s *sqlStore) GetAll() ([]domain.Appointment, error) {
	appointments := []domain.Appointment{}
	query := "SELECT id, dentist_id, patient_id, date, time, description FROM appointments;"
	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err.Error())
	}
	var appointment domain.Appointment
	for rows.Next() {
		err := rows.Scan(&appointment.Id, &appointment.Dentist, &appointment.Patient, &appointment.Date, &appointment.Time, &appointment.Description)
		if err != nil {
			return appointments, err
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}

func (s *sqlStore) GetDniP(dni string) ([]domain.Appointment, error) {
	appointments := []domain.Appointment{}
	query := "SELECT appointments.id, dentist_id, patient_id, date, time, description FROM appointments INNER JOIN patients ON appointments.patient_id = patients.id WHERE patients.identity_document = ?;"
	rows, err := s.db.Query(query, dni)
	if err != nil {
		log.Println(err.Error())
	}
	var appointment domain.Appointment
	for rows.Next() {
		err := rows.Scan(&appointment.Id, &appointment.Dentist, &appointment.Patient, &appointment.Date, &appointment.Time, &appointment.Description)
		if err != nil {
			return appointments, err
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}

func (s *sqlStore) Create(appointment domain.Appointment) error {
	query := "INSERT INTO appointments (dentist_id, patient_id, date, time, description) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	fmt.Println(appointment)
	res, err := stmt.Exec(appointment.Dentist, appointment.Patient, appointment.Date, appointment.Time, appointment.Description)
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

func (s *sqlStore) GetPatientByDni(identity_document string) int {
	var id int
	//var patient domain.Patient
	query := "SELECT id FROM patients WHERE identity_document = ?;"
	row := s.db.QueryRow(query, identity_document)
	err := row.Scan(&id)
	if err != nil {
		return id
	}
	fmt.Println("paciente id: ", id)
	return id
}

func (s *sqlStore) GetDentistByLicense(lic string) int {
	var id int
	//var dentist domain.Dentist
	query := "SELECT id FROM dentists WHERE license = ?;"
	row := s.db.QueryRow(query, lic)
	err := row.Scan(&id)
	if err != nil {
		return id
	}
	fmt.Println("dentista id: ", id)
	return id
}

/*
func (s *sqlStore) CreateByDniAndLicense(appointment domain.Appointment, lic string, dni string) error {
	if s.GetDentistByLicense(lic) > 0 && s.GetPatientByDni(dni) > 0 {
		if s.GetDentistByLicense(lic) == appointment.Dentist && s.GetPatientByDni(dni) == appointment.Dentist {
			query := "INSERT INTO appointments (dentist_id, patient_id, date, time) VALUES (?, ?, ?, ?);"
			stmt, err := s.db.Prepare(query)
			if err != nil {
				return err
			}
			fmt.Println(appointment)
			res, err := stmt.Exec(appointment.Dentist, appointment.Patient, appointment.Date, appointment.Time)
			if err != nil {
				fmt.Println(err)
				return err
			}
			_, err = res.RowsAffected()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
*/
func (s *sqlStore) UpdateOne(appointment domain.Appointment) error {
	query := "UPDATE appointments SET dentist_id = ?, patient_id = ?, date = ?, time = ?, description = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(appointment.Dentist, appointment.Patient, appointment.Date, appointment.Time, appointment.Description, appointment.Id)
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
	query := "DELETE FROM appointments WHERE id = ?;"
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

func (s *sqlStore) ValidateTime(time string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM appointments WHERE time = ?;"
	row := s.db.QueryRow(query, time)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
