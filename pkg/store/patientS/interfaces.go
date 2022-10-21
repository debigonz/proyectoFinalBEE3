package patientS

import (
	"examenFinal/internal/domain"
)

type StoreInterface interface {
	GetOne(id int) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Create(patient domain.Patient) error
	UpdateOne(patient domain.Patient) error
	DeleteOne(id int) error
	ValidateIdentityDocument(identity string) bool
}
