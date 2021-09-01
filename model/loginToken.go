package model

type LoginToken struct {
	PasswordList map[string]string `json:"p"`
	Date         string            `json:"d"`
	IP           string            `json:"i"`
	Uuid         string            `json:"u"`
}
