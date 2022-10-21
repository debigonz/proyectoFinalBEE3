package handler

import (
	"errors"
	"examenFinal/internal/domain"
	"examenFinal/internal/patient"
	"examenFinal/pkg/web"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s patient.Service
}

// NewPatientHandler crea un nuevo controller de pacientes
func NewPatientHandler(s patient.Service) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

// Get obtiene un paciente por id
//@Summary Get patient by id
//@Tags domain.Patient
//@Produce json
//@Success 200 {object} web.response
//@Router /patients/id [get]
func (h *patientHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		web.Success(c, 200, patient)
	}
}

// GetAll obtiene todos los odontólogos
//@Summary Get all patients
//@Tags domain.Patient
//@Produce json
//@Success 200 {object} web.response
//@Router /patients [get]
func (h *patientHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		patients, _ := h.s.GetAll()
		web.Success(c, 200, patients)
	}
}

// Post crea un nuevo paciente
//@Summary Create a patient
//@Tags domain.Patient
//@Produce json
//@Success 201 {object} web.response
//@Router /patients [post]
func (h *patientHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient domain.Patient

		err := c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		p, err := h.s.Create(patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete elimina un paciente
//@Summary Delete a patient by id
//@Tags domain.Patient
//@Produce json
//@Success 200 {object} web.response
//@Router /patients/id [delete]
func (h *patientHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, nil)
	}
}

// validateEmpty valida que los campos no esten vacios
func validateEmpty(patient *domain.Patient) (bool, error) {
	if patient.Lastname == "" || patient.Name == "" || patient.Address == "" || patient.IdentityDocument == "" || patient.EntryDate == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// validateEntryDate valida que la fecha de expiracion sea valida
func validateEntryDate(entry string) (bool, error) {
	dates := strings.Split(entry, "-")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid entry date, must be in format: yyyy-dd-mm")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid entry date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 9999) && (list[1] < 1 || list[1] > 31) && (list[2] < 1 || list[2] > 12)
	if condition {
		return false, errors.New("invalid entry date, date must be between 1 and 9999-31-12")
	}
	return true, nil
}

// Put actualiza un paciente
//@Summary Update a patient by id
//@Tags domain.Patient
//@Produce json
//@Success 200 {object} web.response
//@Router /patients/id [put]
func (h *patientHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var patient domain.Patient
		err = c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmpty(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateEntryDate(patient.EntryDate)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, patient)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch actualiza un patient o alguno de sus campos
//@Summary Update some fills of patient by id
//@Tags domain.Patient
//@Produce json
//@Success 200 {object} web.response
//@Router /patients/id [patch]
func (h *patientHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Lastname         string `json:"lastname,omitempty"`
		Name             string `json:"name,omitempty"`
		Address          string `json:"address,omitempty"`
		IdentityDocument string `json:"identity_document,omitempty"`
		EntryDate        string `json:"entry_date,omitempty"`
	}
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Patient{
			Lastname:         r.Lastname,
			Name:             r.Name,
			Address:          r.Address,
			IdentityDocument: r.IdentityDocument,
			EntryDate:        r.EntryDate,
		}
		if update.EntryDate != "" {
			valid, err := validateEntryDate(update.EntryDate)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
