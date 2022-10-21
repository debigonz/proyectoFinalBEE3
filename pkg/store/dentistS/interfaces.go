package dentistS

import (
	"examenFinal/internal/domain"
)

type StoreInterface interface {
	GetOne(id int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Create(dentist domain.Dentist) error
	UpdateOne(dentist domain.Dentist) error
	DeleteOne(id int) error
	ValidateLicense(license string) bool
}
