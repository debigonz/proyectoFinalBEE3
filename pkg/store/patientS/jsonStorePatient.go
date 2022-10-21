package patientS

import (
	"encoding/json"
	"errors"
	"os"

	"examenFinal/internal/domain"
)

type StoreP interface {
	GetAllP() ([]domain.Patient, error)
	GetOneP(id int) (domain.Patient, error)
	AddOneP(dentist domain.Patient) error
	UpdateOneP(dentist domain.Patient) error
	DeleteOneP(id int) error
	savePatients(dentists []domain.Patient) error
	loadPatient() ([]domain.Patient, error)
}

type jsonStoreP struct {
	pathToFile string
}

// loadProducts carga los productos desde un archivo json
func (s *jsonStoreP) loadPatient() ([]domain.Patient, error) {
	var patients []domain.Patient
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &patients)
	if err != nil {
		return nil, err
	}
	return patients, nil
}

// saveProducts guarda los productos en un archivo json
func (s *jsonStoreP) savePatients(patients []domain.Patient) error {
	bytes, err := json.Marshal(patients)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

/*
// NewJsonStore crea un nuevo store de products
func NewStoreP(path string) StoreP {
	return &jsonStoreP{
		pathToFile: path,
	}
}
*/
// GetAll devuelve todos los productos
func (s *jsonStoreP) GetAllP() ([]domain.Patient, error) {
	patients, err := s.loadPatient()
	if err != nil {
		return nil, err
	}
	return patients, nil
}

// GetOne devuelve un producto por su id
func (s *jsonStoreP) GetOneP(id int) (domain.Patient, error) {
	patients, err := s.loadPatient()
	if err != nil {
		return domain.Patient{}, err
	}
	for _, patient := range patients {
		if patient.Id == id {
			return patient, nil
		}
	}
	return domain.Patient{}, errors.New("patient not found")
}

// AddOne agrega un nuevo producto
func (s *jsonStoreP) AddOneP(patient domain.Patient) error {
	patients, err := s.loadPatient()
	if err != nil {
		return err
	}
	patient.Id = len(patients) + 1
	patients = append(patients, patient)
	return s.savePatients(patients)
}

// UpdateOne actualiza un producto
func (s *jsonStoreP) UpdateOneP(patient domain.Patient) error {
	patients, err := s.loadPatient()
	if err != nil {
		return err
	}
	for i, p := range patients {
		if p.Id == patient.Id {
			patients[i] = patient
			return s.savePatients(patients)
		}
	}
	return errors.New("patient not found")
}

// DeleteOne elimina un producto
func (s *jsonStoreP) DeleteOne(id int) error {
	patients, err := s.loadPatient()
	if err != nil {
		return err
	}
	for i, p := range patients {
		if p.Id == id {
			patients = append(patients[:i], patients[i+1:]...)
			return s.savePatients(patients)
		}
	}
	return errors.New("patient not found")
}
