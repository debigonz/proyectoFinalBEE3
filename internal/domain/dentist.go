package domain

type Dentist struct {
	Id       int    `json:"id"`
	Lastname string `json:"lastname"`
	Name     string `json:"name"`
	License  string `json:"license"`
}
