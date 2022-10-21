package appoinmentS

import (
	"examenFinal/internal/domain"
)

type StoreInterface interface {
	GetOne(id int) (domain.Appointment, error)
	GetAll() ([]domain.Appointment, error)
	GetDniP(dni string) ([]domain.Appointment, error)
	GetPatientByDni(dni string) (id int)
	GetDentistByLicense(lic string) (id int)
	Create(appoinment domain.Appointment) error
	UpdateOne(appoinment domain.Appointment) error
	DeleteOne(id int) error
	ValidateTime(time string) bool
}
