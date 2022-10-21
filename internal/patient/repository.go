package patient

import (
	"errors"
	"examenFinal/internal/domain"
	patient "examenFinal/pkg/store/patientS"
)

type Repository interface {
	Create(p domain.Patient) (domain.Patient, error)
	GetByID(id int) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Update(id int, p domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type repository struct {
	storage patient.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage patient.StoreInterface) Repository {
	return &repository{storage}
}

// GetByID busca un paciente por su id
func (r *repository) GetByID(id int) (domain.Patient, error) {
	patient, err := r.storage.GetOne(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil
}

// GetAll busca a todos los pacientes
func (r *repository) GetAll() ([]domain.Patient, error) {
	patients, err := r.storage.GetAll()
	if err != nil {
		return []domain.Patient{}, errors.New("there are not patients")
	}
	return patients, nil
}

// Create agrega un nuevo paciente
func (r *repository) Create(p domain.Patient) (domain.Patient, error) {
	if r.storage.ValidateIdentityDocument(p.IdentityDocument) {
		return domain.Patient{}, errors.New("identity document already exists")
	}
	err := r.storage.Create(p)
	if err != nil {
		return domain.Patient{}, errors.New("error adding patient")
	}
	return p, nil
}

// Delete elimina un paciente
func (r *repository) Delete(id int) error {
	err := r.storage.DeleteOne(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un paciente
func (r *repository) Update(id int, p domain.Patient) (domain.Patient, error) {
	if !r.storage.ValidateIdentityDocument(p.IdentityDocument) {
		return domain.Patient{}, errors.New("identity document already exists")
	}
	err := r.storage.UpdateOne(p)
	if err != nil {
		return domain.Patient{}, errors.New("error updating patient")
	}
	return p, nil
}
