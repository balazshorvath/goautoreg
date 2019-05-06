package goautoreg

import (
	"testing"

	"github.com/gonum/floats"
)

func TestYuleWalkerAR_Fit_OrderTooBig(t *testing.T) {
	yw := NewYuleWalkerARFit()
	order := 3

	_, err := yw(
		[]float64{
			1,
			2,
		},
		order,
	)
	if err == nil {
		t.Fatal("Error expected")
	}
}

func TestYuleWalkerAR_Fit(t *testing.T) {
	yw := NewYuleWalkerARFit()
	order := 3
	// Checked with a python implementation
	expected := []float64{0.7466340269277845, -0.06744186046511638, -0.09547123623011006}

	coefs, err := yw(
		[]float64{
			1,
			2,
			3,
			4,
		},
		order,
	)
	if err != nil {
		t.Fatalf("Fit failed: %v", err)
	}
	if len(coefs) != order {
		t.Fatalf("Invlaid coefficient array length")
	}
	if !floats.Equal(expected, coefs) {
		t.Errorf("Expected : %v, But got : %v", expected, coefs)
	}
}
