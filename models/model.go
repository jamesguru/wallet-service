package models

type Wallet struct {
	Name   string `json:"name" xml:"name" validate:"required"`
	Author string `json:"author" validate:"required"`
}
