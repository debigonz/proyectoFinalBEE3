package dentist

import (
	"examenFinal/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Create(p domain.Dentist) (domain.Dentist, error)
	Update(id int, p domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetByID busca un odontólogo por su id
func (s *service) GetByID(id int) (domain.Dentist, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}

// GetAll busca todos los odontólogos
func (s *service) GetAll() ([]domain.Dentist, error) {
	dentists, err := s.r.GetAll()
	if err != nil {
		return []domain.Dentist{}, err
	}
	return dentists, nil
}

// Create agrega un nuevo odontólogo
func (s *service) Create(d domain.Dentist) (domain.Dentist, error) {
	d, err := s.r.Create(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}

// Delete elimina un odontólogo
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un odontólogo
func (s *service) Update(id int, u domain.Dentist) (domain.Dentist, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	if u.Name != "" {
		d.Name = u.Name
	}
	if u.Lastname != "" {
		d.Lastname = u.Lastname
	}
	if u.License != "" {
		d.License = u.License
	}
	d, err = s.r.Update(id, d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}
