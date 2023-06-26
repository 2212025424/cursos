package storage

import (
	"testing"

	"github.com/2212025424/api/model"
)

func TestCreate_empty_person(t *testing.T) {
	m := NewMemory()
	err := m.Create(nil)

	if err == nil {
		t.Error("Se esperaba un error y se obtuvo nil")
	}

	if err != model.ErrPersonCanNotBeNil {
		t.Errorf("Se esperaba el error: %v - Y se obtuvo: %v", model.ErrPersonCanNotBeNil, err)
	}
}

func TestCreate_count_elements(t *testing.T) {
	m := NewMemory()
	p := model.Person{Name: "kike"}

	err := m.Create(&p)

	if err != nil {
		t.Errorf("No se esperaba un error, se obtuvo %v", err)
	}

	if m.currentID != 1 {
		t.Errorf("Se esperaba un 1, se obtuvo %d", m.currentID)
	}
}
