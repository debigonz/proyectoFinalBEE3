package handler

import (
	"errors"
	"examenFinal/internal/dentist"
	"examenFinal/internal/domain"
	"examenFinal/pkg/web"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.Service
}

// NewDentistHandler crea un nuevo controller de odontólogos
func NewDentistHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

// Get obtiene un odontólogo por id
//@Summary Get dentist by id
//@Tags domain.Dentist
//@Produce json
//@Success 200 {object} web.response
//@Router /dentists/id [get]
func (h *dentistHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentist, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		web.Success(c, 200, dentist)
	}
}

// GetAll obtiene todos los odontólogos
//@Summary Get all dentists
//@Tags domain.Dentist
//@Produce json
//@Success 200 {object} web.response
//@Router /dentists [get]
func (h *dentistHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		dentists, _ := h.s.GetAll()
		web.Success(c, 200, dentists)
	}
}

// Post crea un nuevo odontólogo
//@Summary Create a dentist
//@Tags domain.Dentist
//@Produce json
//@Success 201 {object} web.response
//@Router /dentists [post]
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist

		err := c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		d, err := h.s.Create(dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

// Delete elimina un odontólogo
//@Summary Delete a dentist by id
//@Tags domain.Dentist
//@Produce json
//@Success 200 {object} web.response
//@Router /dentists/id [delete]
func (h *dentistHandler) Delete() gin.HandlerFunc {
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

// validateEmptys valida que los campos no esten vacios
func validateEmptys(dentist *domain.Dentist) (bool, error) {
	if dentist.Lastname == "" || dentist.Name == "" || dentist.License == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Put actualiza un odontólogo
//@Summary Update a dentist by id
//@Tags domain.Dentist
//@Produce json
//@Success 200 {object} web.response
//@Router /dentists/id [put]
func (h *dentistHandler) Put() gin.HandlerFunc {
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
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var dentist domain.Dentist
		err = c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

// Patch actualiza un odontólogo o alguno de sus campos
//@Summary Update some fills of dentist by id
//@Tags domain.Dentist
//@Produce json
//@Success 200 {object} web.response
//@Router /dentists/id [patch]
func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Lastname string `json:"lastname,omitempty"`
		Name     string `json:"name,omitempty"`
		License  string `json:"license,omitempty"`
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
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentist{
			Lastname: r.Lastname,
			Name:     r.Name,
			License:  r.License,
		}
		d, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}
