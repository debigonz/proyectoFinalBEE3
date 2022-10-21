package domain

type Appointment struct {
	Id          int    `json:"id"`
	Dentist     int    `json:"dentist_id"`
	Patient     int    `json:"patient_id"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Description string `json:"description"`
}
