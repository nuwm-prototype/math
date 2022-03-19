package math

import "testing"

func TestNew(t *testing.T) {
	galoisImpl4 := New(4)
	if galoisImpl4 == nil {
		t.Error("Galois implementation for 4 bit words is null")
	}
	galoisImpl8 := New(8)
	if galoisImpl8 == nil {
		t.Error("Galois implementation for 8 bit words is null")
	}
	galoisImpl16 := New(16)
	if galoisImpl16 != nil {
		t.Error("Galois implementation for 16 bit words must not present")
	}
}

func TestAdd(t *testing.T) {
	galoisImpl4 := New(4)
	addResult := galoisImpl4.Add(11, 7)
	if addResult != 12 {
		t.Errorf("11 + 7 must be equal to 12 in galois field 2^4, but equals %d", addResult)
	}
}

func TestSub(t *testing.T) {
	galoisImpl4 := New(4)
	subResult := galoisImpl4.Sub(11, 7)
	if subResult != 12 {
		t.Errorf("11 - 7 must be equal to 12 in galois field 2^4, but equals %d", subResult)
	}
}

func TestMul(t *testing.T) {
	galoisImpl4 := New(4)
	mulResult := galoisImpl4.Mul(3, 7)
	if mulResult != 9 {
		t.Errorf("3 * 7 must be equal to 9 in galois field 2^4, but equals %d", mulResult)
	}
	mulResult = galoisImpl4.Mul(13, 10)
	if mulResult != 11 {
		t.Errorf("13 * 10 must be equal to 11 in galois field 2^4, but equals %d", mulResult)
	}
}

func TestDiv(t *testing.T) {
	galoisImpl4 := New(4)
	divResult := galoisImpl4.Div(3, 7)
	if divResult != 10 {
		t.Errorf("3 / 7 must be equal to 10 in galois field 2^4, but equals %d", divResult)
	}
	divResult = galoisImpl4.Div(13, 10)
	if divResult != 3 {
		t.Errorf("13 / 10 must be equal to 3 in galois field 2^4, but equals %d", divResult)
	}
}
