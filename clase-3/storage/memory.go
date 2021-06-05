package storage

import "https://github.com/whiteagleinc-meli/curso_api_rest_go/clase-3/model"

// Memory .
type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

// NewMemory .
func NewMemory() Memory {
	persons := make(map[int]model.Person)

	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

// Create .
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeBil
	}

	m.currentID++
	m.Persons[m.currentID] = *person

	return nil
}