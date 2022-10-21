package domain

type Patient struct {
	Id               int    `json:"id"`
	Lastname         string `json:"lastname"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	IdentityDocument string `json:"identity_document"`
	EntryDate        string `json:"entry_date"`
}
