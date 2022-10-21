package appoinmentS

import (
	"encoding/json"
	"errors"
	"os"

	"examenFinal/internal/domain"
)

type StoreA interface {
	GetAllA() ([]domain.Appointment, error)
	GetOneA(id int) (domain.Appointment, error)
	GetOneAA(iden string) (domain.Appointment, error)
	AddOneA(appointment domain.Appointment) error
	UpdateOneA(appointment domain.Appointment) error
	DeleteOneA(id int) error
	saveAppointments(appointment []domain.Appointment) error
	loadAppointment() ([]domain.Appointment, error)
}

type jsonStoreA struct {
	pathToFile string
}

// loadProducts carga los productos desde un archivo json
func (s *jsonStoreA) loadAppointment() ([]domain.Appointment, error) {
	var appointments []domain.Appointment
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &appointments)
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

// saveProducts guarda los productos en un archivo json
func (s *jsonStoreA) saveAppointments(appointments []domain.Appointment) error {
	bytes, err := json.Marshal(appointments)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

/*
// NewJsonStore crea un nuevo store de products
func NewStoreA(path string) StoreA {
	return &jsonStoreA{
		pathToFile: path,
	}
}
*/
// GetAll devuelve todos los productos
func (s *jsonStoreA) GetAllA() ([]domain.Appointment, error) {
	appointments, err := s.loadAppointment()
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

// GetOne devuelve un producto por su id
func (s *jsonStoreA) GetOneA(id int) (domain.Appointment, error) {
	appointments, err := s.loadAppointment()
	if err != nil {
		return domain.Appointment{}, err
	}
	for _, appointment := range appointments {
		if appointment.Id == id {
			return appointment, nil
		}
	}
	return domain.Appointment{}, errors.New("appointment not found")
}

// GetOne devuelve un producto por su id
/*
func (s *jsonStoreA) GetOneAA(iden string) (domain.Appointment, error) {
	appointments, err := s.loadAppointment()
	if err != nil {
		return domain.Appointment{}, err
	}
	for _, appointment := range appointments {
		//if appointment.Patient.IdentityDocument == iden {
			return appointment, nil
		}
	}
	return domain.Appointment{}, errors.New("appointment not found")
}
*/
// AddOne agrega un nuevo producto
func (s *jsonStoreA) AddOneA(appointment domain.Appointment) error {
	appointments, err := s.loadAppointment()
	if err != nil {
		return err
	}
	appointment.Id = len(appointments) + 1
	appointments = append(appointments, appointment)
	return s.saveAppointments(appointments)
}

// UpdateOne actualiza un producto
func (s *jsonStoreA) UpdateOneA(appointment domain.Appointment) error {
	appointments, err := s.loadAppointment()
	if err != nil {
		return err
	}
	for i, p := range appointments {
		if p.Id == appointment.Id {
			appointments[i] = appointment
			return s.saveAppointments(appointments)
		}
	}
	return errors.New("appointment not found")
}

// DeleteOne elimina un producto
func (s *jsonStoreA) DeleteOneA(id int) error {
	appointments, err := s.loadAppointment()
	if err != nil {
		return err
	}
	for i, p := range appointments {
		if p.Id == id {
			appointments = append(appointments[:i], appointments[i+1:]...)
			return s.saveAppointments(appointments)
		}
	}
	return errors.New("appointment not found")
}
