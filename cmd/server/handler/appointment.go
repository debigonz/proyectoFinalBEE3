package handler

import (
	"errors"
	"examenFinal/internal/appointment"
	"examenFinal/internal/domain"
	"examenFinal/pkg/web"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	s appointment.Service
}

// NewPatientHandler crea un nuevo controller de turnos
func NewAppointmentHandler(s appointment.Service) *appointmentHandler {
	return &appointmentHandler{
		s: s,
	}
}

// Get obtiene un turno por id
//@Summary Get appointment by id
//@Tags domain.Appointment
//@Produce json
//@Success 200 {object} web.response
//@Router /appointments/id [get]
func (h appointmentHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		appoinment, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		web.Success(c, 200, appoinment)
	}
}

// Get obtiene un turno por dni paciente
//@Summary Get appointment by patient identity document
//@Tags domain.Appointment
//@Produce json
//@Success 200 {object} web.response
//@Router /appointments/dni [get]
func (h appointmentHandler) GetByDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := strings.Split(c.Request.RequestURI, "/")
		fmt.Println(ctx)
		dnic := ctx[3]
		fmt.Println(dnic)
		var err error
		if err != nil {
			web.Failure(c, 400, errors.New("invalid dni"))
			return
		}
		appoinment, err := h.s.GetByIdentityDocument(dnic)
		if err != nil {
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		web.Success(c, 200, appoinment)
	}
}

// GetAll obtiene todos los turnos
//@Summary Get appointments
//@Tags domain.Appointment
//@Produce json
//@Success 200 {object} web.response
//@Router /appointments [get]
func (h *appointmentHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		appointments, _ := h.s.GetAll()
		web.Success(c, 200, appointments)
	}
}

// Post crea un nuevo turno
//@Summary Create appointment
//@Tags domain.Appointment
//@Produce json
//@Success 201 {object} web.response
//@Router /appointments [post]
func (h *appointmentHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appoinment domain.Appointment

		err := c.ShouldBindJSON(&appoinment)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		a, err := h.s.Create(appoinment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, a)
	}
}

// Post crea un nuevo turno por dni paciente y matrícula de odontólogo
//@Summary Create appointment by patient identity document and dentist license
//@Tags domain.Appointment
//@Produce json
//@Success 201 {object} web.response
//@Router /appointments/dni/lic [post]
func (h *appointmentHandler) PostDP() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appoinment domain.Appointment
		ctx := strings.Split(c.Request.RequestURI, "/")
		dni := ctx[2]
		lic := ctx[3]

		err := c.ShouldBindJSON(&appoinment)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		a, err := h.s.CreateByDniAndLicense(appoinment, dni, lic)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, a)
	}
}

// Delete elimina un turno
//@Summary Delete a appointment by id
//@Tags domain.Appointment
//@Produce json
//@Success 200 {object} web.response
//@Router /appointments/id [delete]
func (h *appointmentHandler) Delete() gin.HandlerFunc {
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

// validateEmptyFills valida que los campos no esten vacios
func validateEmptyFills(apo *domain.Appointment) (bool, error) {
	switch {
	case apo.Date == "" || apo.Time == "":
		return false, errors.New("fields can't be empty")
	case apo.Dentist == 0 || apo.Patient == 0:
		if apo.Dentist == 0 {
			return false, errors.New("dentist doesn't exist")
		}
		if apo.Patient == 0 {
			return false, errors.New("patient doesn't exist")
		}
	}
	return true, nil
}

// validateate valida que la fecha de expiracion sea valida
func validateDate(date string) (bool, error) {
	dates := strings.Split(date, "-")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid date, must be in format: yyyy-dd-mm")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 9999) && (list[1] < 1 || list[1] > 31) && (list[2] < 1 || list[2] > 12)
	if condition {
		return false, errors.New("invalid date, date must be between 1 and 9999-31-12")
	}
	return true, nil
}

// Put actualiza un turno
//@Summary Update a appointment by id
//@Tags domain.Appointment
//@Produce json
//@Success 200 {object} web.response
//@Router /appointments/id [put]
func (h *appointmentHandler) Put() gin.HandlerFunc {
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
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var appoinment domain.Appointment
		err = c.ShouldBindJSON(&appoinment)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptyFills(&appoinment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateDate(appoinment.Date)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		a, err := h.s.Update(id, appoinment)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, a)
	}
}

// Patch actualiza un turno o alguno de sus campos
//@Summary Update some fils of appointment by id
//@Tags domain.Appointment
//@Produce json
//@Success 200 {object} web.response
//@Router /appointments/id [patch]
func (h *appointmentHandler) Patch() gin.HandlerFunc {
	//var patient Patient := domain.Patient.
	type Request struct {
		Id          int    `json:"id,omitempty"`
		Dentist     int    `json:"dentist_id,omitempty"`
		Patient     int    `json:"patient_id,omitempty"`
		Date        string `json:"date,omitempty"`
		Time        string `json:"time,omitempty"`
		Description string `json:"description,omitempty"`
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
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Appointment{
			Dentist:     r.Dentist,
			Patient:     r.Patient,
			Date:        r.Date,
			Time:        r.Time,
			Description: r.Description,
		}
		if update.Date != "" {
			valid, err := validateDate(update.Date)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		a, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, a)
	}
}
