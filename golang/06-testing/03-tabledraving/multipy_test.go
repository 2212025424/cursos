package tabledraving

import "testing"

func TestMultiply(t *testing.T) {

	table := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"2x1", 2, 1, 2},
		{"2x2", 2, 2, 4},
		{"2x3", 2, 3, 6},
		{"2x4", 2, 4, 8},
		{"2x5", 2, 5, 10},
		{"2x6", 2, 6, 12},
		{"2x7", 2, 7, 14},
		{"2x8", 2, 8, 16},
		{"2x9", 2, 9, 18},
		{"2x10", 2, 10, 20},
	}

	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			want := v.want
			got := multiply(v.a, v.b)

			if want != got {
				t.Errorf("Se obtuvo %d, se esperaba %d", got, want)
			}
		})
	}

}
