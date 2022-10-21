package dentistS

import (
	"encoding/json"
	"errors"
	"os"

	"examenFinal/internal/domain"
)

type Store interface {
	GetAll() ([]domain.Dentist, error)
	GetOne(id int) (domain.Dentist, error)
	AddOne(dentist domain.Dentist) error
	UpdateOne(dentist domain.Dentist) error
	DeleteOne(id int) error
	saveDentists(dentists []domain.Dentist) error
	loadDentist() ([]domain.Dentist, error)
}

type jsonStore struct {
	pathToFile string
}

// loadProducts carga los productos desde un archivo json
func (s *jsonStore) loadDentist() ([]domain.Dentist, error) {
	var dentists []domain.Dentist
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &dentists)
	if err != nil {
		return nil, err
	}
	return dentists, nil
}

// saveProducts guarda los productos en un archivo json
func (s *jsonStore) saveDentists(dentists []domain.Dentist) error {
	bytes, err := json.Marshal(dentists)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

// NewJsonStore crea un nuevo store de products
func NewStore(path string) Store {
	return &jsonStore{
		pathToFile: path,
	}
}

// GetAll devuelve todos los productos
func (s *jsonStore) GetAll() ([]domain.Dentist, error) {
	dentists, err := s.loadDentist()
	if err != nil {
		return nil, err
	}
	return dentists, nil
}

// GetOne devuelve un producto por su id
func (s *jsonStore) GetOne(id int) (domain.Dentist, error) {
	dentists, err := s.loadDentist()
	if err != nil {
		return domain.Dentist{}, err
	}
	for _, dentist := range dentists {
		if dentist.Id == id {
			return dentist, nil
		}
	}
	return domain.Dentist{}, errors.New("dentist not found")
}

// AddOne agrega un nuevo producto
func (s *jsonStore) AddOne(dentist domain.Dentist) error {
	dentists, err := s.loadDentist()
	if err != nil {
		return err
	}
	dentist.Id = len(dentists) + 1
	dentists = append(dentists, dentist)
	return s.saveDentists(dentists)
}

// UpdateOne actualiza un producto
func (s *jsonStore) UpdateOne(product domain.Dentist) error {
	dentists, err := s.loadDentist()
	if err != nil {
		return err
	}
	for i, p := range dentists {
		if p.Id == product.Id {
			dentists[i] = product
			return s.saveDentists(dentists)
		}
	}
	return errors.New("dentist not found")
}

// DeleteOne elimina un producto
func (s *jsonStore) DeleteOne(id int) error {
	dentists, err := s.loadDentist()
	if err != nil {
		return err
	}
	for i, p := range dentists {
		if p.Id == id {
			dentists = append(dentists[:i], dentists[i+1:]...)
			return s.saveDentists(dentists)
		}
	}
	return errors.New("dentist not found")
}
