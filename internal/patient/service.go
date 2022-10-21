package patient

import (
	"examenFinal/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Create(p domain.Patient) (domain.Patient, error)
	Update(id int, p domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetByID busca un paciente por su id
func (s *service) GetByID(id int) (domain.Patient, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return d, nil
}

// GetAll busca todos los pacientes
func (s *service) GetAll() ([]domain.Patient, error) {
	patients, err := s.r.GetAll()
	if err != nil {
		return []domain.Patient{}, err
	}
	return patients, nil
}

// Create agrega un nuevo paciente
func (s *service) Create(p domain.Patient) (domain.Patient, error) {
	d, err := s.r.Create(p)
	if err != nil {
		return domain.Patient{}, err
	}
	return d, nil
}

// Delete elimina un paciente
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un paciente
func (s *service) Update(id int, u domain.Patient) (domain.Patient, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	if u.Name != "" {
		p.Name = u.Name
	}
	if u.Lastname != "" {
		p.Lastname = u.Lastname
	}
	if u.Address != "" {
		p.Address = u.Address
	}
	if u.IdentityDocument != "" {
		p.IdentityDocument = u.IdentityDocument
	}
	if u.EntryDate != "" {
		p.EntryDate = u.EntryDate
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}
