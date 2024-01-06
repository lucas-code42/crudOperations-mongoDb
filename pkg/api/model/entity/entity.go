package entity

import "fmt"

type Book struct {
	Id     string
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (b *Book) CheckFields() error {
	if b.Title == "" {
		return fmt.Errorf("campo title inválido")
	}

	if b.Author == "" {
		return fmt.Errorf("campo author inválido")
	}

	return nil
}
