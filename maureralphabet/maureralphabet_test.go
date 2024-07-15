package maureralphabet

import (
	"testing"
)

func TestConvertToBuchstabe(t *testing.T) {
	example := "ABCDEFZ"
	test := []Buchstabe{
		Buchstabe{Str: "A", Num: 1},
		Buchstabe{Str: "B", Num: 2},
		Buchstabe{Str: "C", Num: 3},
		Buchstabe{Str: "D", Num: 4},
		Buchstabe{Str: "E", Num: 5},
		Buchstabe{Str: "F", Num: 6},
		Buchstabe{Str: "Z", Num: 26},
	}

	out, err := ConvertToBuchstaben(example)

	if err != nil {
		t.Fatalf("An Error occured: %v", err)
	}

	for index, b := range out {
		if b.Str != test[index].Str || b.Num != test[index].Num {
			t.Fatalf("Buchstabe ungültig! %v != %v", b, test[index])
		}
	}
}
