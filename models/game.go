package models

type Move struct {
	From  string `json:"from"`
	To    string `json:"to"`
	White bool   `json:"white"`
}
