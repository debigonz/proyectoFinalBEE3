package appointment

import (
	"errors"
	"examenFinal/internal/domain"
	appoinment "examenFinal/pkg/store/appointmentS"
	"log"
)

type Repository interface {
	Create(p domain.Appointment) (domain.Appointment, error)
	SearchPatientByDNI(dni string) (id int)
	SearchDentistByLicense(lic string) (id int)
	GetByID(id int) (domain.Appointment, error)
	GetAll() ([]domain.Appointment, error)
	GetByIdentityDocument(dni string) ([]domain.Appointment, error)
	Update(id int, p domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type repository struct {
	storage appoinment.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage appoinment.StoreInterface) Repository {
	return &repository{storage}
}

// GetByID busca un turno por su id
func (r *repository) GetByID(id int) (domain.Appointment, error) {
	appointment, err := r.storage.GetOne(id)
	if err != nil {
		return domain.Appointment{}, errors.New("appointment not found")
	}
	return appointment, nil
}

// GetAll busca todos turno
func (r *repository) GetAll() ([]domain.Appointment, error) {
	appointments, err := r.storage.GetAll()
	if err != nil {
		return []domain.Appointment{}, errors.New("there are not appointment")
	}
	return appointments, nil
}

// GetByID busca todos turnos por su dni del paciente
func (r *repository) GetByIdentityDocument(ident string) ([]domain.Appointment, error) {
	appointment, err := r.storage.GetDniP(ident)
	if err != nil {
		return []domain.Appointment{}, errors.New("appointment not found")
	}
	return appointment, nil
}

//  SearchPatientByDNI busca un paciente por su dni
func (r *repository) SearchPatientByDNI(dni string) (id int) {
	p := r.storage.GetPatientByDni(dni)
	if p == 0 {
		log.Fatal("Id did not exist")
	}
	return p
}

//  SearchDentistByLicense busca un odontólogo por su matrícula
func (r *repository) SearchDentistByLicense(lic string) (id int) {
	p := r.storage.GetDentistByLicense(lic)
	if p == 0 {
		log.Fatal("Id did not exist")
	}
	return p
}

// Create agrega un nuevo turno
func (r *repository) Create(a domain.Appointment) (domain.Appointment, error) {
	err := r.storage.Create(a)
	if err != nil {
		return domain.Appointment{}, errors.New("error adding appointment")
	}
	return a, nil
}

// Delete elimina un turno
func (r *repository) Delete(id int) error {
	err := r.storage.DeleteOne(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un turno
func (r *repository) Update(id int, a domain.Appointment) (domain.Appointment, error) {
	err := r.storage.UpdateOne(a)
	if err != nil {
		return domain.Appointment{}, errors.New("error updating turno")
	}
	return a, nil
}
