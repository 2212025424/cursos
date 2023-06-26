package introduccion

import "testing"

func TestAdd(t *testing.T) {
	want := 8
	got := Add(2, 3)

	if got != want {
		t.Errorf("ERRROR: Se esperaba %d, se obtuvo: %d", want, got)
	}

	t.Log("Terminó la prueba")
}

func TestAddAcum(t *testing.T) {
	want := 8
	got := AddAcum(1, 2, 3, 4)

	if got != want {
		t.Errorf("ERRROR: Se esperaba %d, se obtuvo: %d", want, got)
	}
}
