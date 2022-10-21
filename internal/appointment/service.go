package appointment

import (
	"examenFinal/internal/domain"
	"fmt"
)

type Service interface {
	GetByID(id int) (domain.Appointment, error)
	GetAll() ([]domain.Appointment, error)
	GetByIdentityDocument(dni string) ([]domain.Appointment, error)
	Create(p domain.Appointment) (domain.Appointment, error)
	CreateByDniAndLicense(a domain.Appointment, dni string, lic string) (domain.Appointment, error)
	Update(id int, p domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetByID busca un turno por su id
func (s *service) GetByID(id int) (domain.Appointment, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return d, nil
}

// GetAll busca todos los turnos
func (s *service) GetAll() ([]domain.Appointment, error) {
	appointments, err := s.r.GetAll()
	if err != nil {
		return []domain.Appointment{}, err
	}
	return appointments, nil
}

// GetByID busca todos los turnos por su dni de paciente
func (s *service) GetByIdentityDocument(ident string) ([]domain.Appointment, error) {
	d, err := s.r.GetByIdentityDocument(ident)
	if err != nil {
		return []domain.Appointment{}, err
	}
	return d, nil
}

// Create agrega un nuevo turno
func (s *service) Create(p domain.Appointment) (domain.Appointment, error) {
	d, err := s.r.Create(p)
	if err != nil {
		return domain.Appointment{}, err
	}
	return d, nil
}

// Create agrega un nuevo turno por dni de paciente y matrícula de odontólogo
func (s *service) CreateByDniAndLicense(a domain.Appointment, dni string, lic string) (domain.Appointment, error) {
	patient := s.r.SearchPatientByDNI(dni)
	fmt.Println(patient)
	dentist := s.r.SearchDentistByLicense(lic)
	fmt.Println(dentist)
	a.Patient = patient
	a.Dentist = dentist
	d, err := s.r.Create(a)
	if err != nil {
		return domain.Appointment{}, err
	}

	return d, nil
}

// Delete elimina un turno
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un turno
func (s *service) Update(id int, u domain.Appointment) (domain.Appointment, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	if u.Dentist != 0 {
		d.Dentist = u.Dentist
	}
	if u.Patient != 0 {
		d.Patient = u.Patient
	}
	if u.Date != "" {
		d.Date = u.Date
	}
	if u.Time != "" {
		d.Time = u.Time
	}
	if u.Description != "" {
		d.Description = u.Description
	}
	d, err = s.r.Update(id, d)
	if err != nil {
		return domain.Appointment{}, err
	}
	return d, nil
}
