package models

// Admin struct describing admin details.
type Admin struct {
	ID       int    `json:"adminID"`
	Name     string `json:"adminName"`
	Password string `json:"password"`
}
