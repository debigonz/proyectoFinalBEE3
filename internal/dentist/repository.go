package dentist

import (
	"errors"
	"examenFinal/internal/domain"
	dentist "examenFinal/pkg/store/dentistS"
)

type Repository interface {
	Create(p domain.Dentist) (domain.Dentist, error)
	GetByID(id int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Update(id int, p domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type repository struct {
	storage dentist.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage dentist.StoreInterface) Repository {
	return &repository{storage}
}

// GetByID busca un odontólogo por su id
func (r *repository) GetByID(id int) (domain.Dentist, error) {
	dentist, err := r.storage.GetOne(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")
	}
	return dentist, nil
}

// GetAll todos los odontólogo
func (r *repository) GetAll() ([]domain.Dentist, error) {
	dentists, err := r.storage.GetAll()
	if err != nil {
		return []domain.Dentist{}, errors.New("there are not dentist")
	}
	return dentists, nil
}

// Create agrega un nuevo odontólogo
func (r *repository) Create(d domain.Dentist) (domain.Dentist, error) {
	if r.storage.ValidateLicense(d.License) {
		return domain.Dentist{}, errors.New("license already exists")
	}
	err := r.storage.Create(d)
	if err != nil {
		return domain.Dentist{}, errors.New("error adding dentist")
	}
	return d, nil
}

// Delete elimina un dentista
func (r *repository) Delete(id int) error {
	err := r.storage.DeleteOne(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un dentista
func (r *repository) Update(id int, d domain.Dentist) (domain.Dentist, error) {
	if !r.storage.ValidateLicense(d.License) {
		return domain.Dentist{}, errors.New("license already exists")
	}
	err := r.storage.UpdateOne(d)
	if err != nil {
		return domain.Dentist{}, errors.New("error updating dentist")
	}
	return d, nil
}
