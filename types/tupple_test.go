package types

import "testing"

func TestIsPoint(t *testing.T) {
	tupple := Tupple{4.3, -4.2, 3.1, 1.0}

	if !tupple.IsPoint() {
		t.Fatalf("Tupple %+v should be a point", tupple)
	}

	if tupple.IsVector() {
		t.Fatalf("Tupple %+v should not be a vector", tupple)
	}
}

func TestIsVector(t *testing.T) {
	tupple := Tupple{4.3, -4.2, 3.1, 0.0}

	if !tupple.IsVector() {
		t.Fatalf("Tupple %+v should be a vector", tupple)
	}

	if tupple.IsPoint() {
		t.Fatalf("Tupple %+v should not be a point", tupple)
	}
}

func TestCreatePoint(t *testing.T) {
	p := CreatePoint(4.0, -4.0, 3.0)

	if !SafeEquals(p.W, 1.0) {
		t.Fatalf("Point needs to have w set to 0.0")
	}
}

func TestCreateVector(t *testing.T) {
	v := CreateVector(4, -4, 3)

	if !SafeEquals(v.W, 0.0) {
		t.Fatalf("Point needs to have w set to 1.0")
	}
}

func TestEquals(t *testing.T) {
	p1 := Tupple{1.0, 2.3, 4.1, 1.0}
	p2 := Tupple{1.0, 2.3, 4.1, 1.0}

	if !p1.Equals(p2) {
		t.Fatalf("%+v and %+v should be equal", p1, p2)
	}

	p1 = Tupple{1.0, 2.3, 4.1, 0.0}
	p2 = Tupple{1.0, 2.3, 4.1, 1.0}

	if p1.Equals(p2) {
		t.Fatalf("%+v and %+v should not be equal", p1, p2)
	}

	p1 = Tupple{1.0, 2.4, 4.1, 1.0}
	p2 = Tupple{1.0, 2.3, 4.1, 1.0}

	if p1.Equals(p2) {
		t.Fatalf("%+v and %+v should not be equal", p1, p2)
	}

	p1 = Tupple{1.1, 2.3, 4.1, 1.0}
	p2 = Tupple{1.0, 2.3, 4.1, 1.0}

	if p1.Equals(p2) {
		t.Fatalf("%+v and %+v should not be equal", p1, p2)
	}

	p1 = Tupple{1.0, 2.3, 4.1, 1.0}
	p2 = Tupple{1.0, 2.3, 4.2, 1.0}

	if p1.Equals(p2) {
		t.Fatalf("%+v and %+v should not be equal", p1, p2)
	}
}

func TestAdd(t *testing.T) {
	a1 := Tupple{3, -2, 5, 1}
	a2 := Tupple{-2, 3, 1, 0}
	e := Tupple{1, 1, 6, 1}

	r := a1.Add(a2)
	if !e.Equals(r) {
		t.Fatalf("%+v and %+v should add to %+v insted of %+v", a1, a2, e, r)
	}
}

func TestSubtractTwoPoints(t *testing.T) {
	a1 := CreatePoint(3, 2, 1)
	a2 := CreatePoint(5, 6, 7)
	e := CreateVector(-2, -4, -6)

	r := a1.Subtract(a2)
	if !e.Equals(r) {
		t.Fatalf("%+v and %+v should subtract to %+v insted of %+v", a1, a2, e, r)
	}
}

func TestSubtractVectorFromAPoint(t *testing.T) {
	a1 := CreatePoint(3, 2, 1)
	a2 := CreateVector(5, 6, 7)
	e := CreatePoint(-2, -4, -6)

	r := a1.Subtract(a2)
	if !e.Equals(r) {
		t.Fatalf("%+v and %+v should subtract to %+v insted of %+v", a1, a2, e, r)
	}
}

func TestSubtractTwoVectors(t *testing.T) {
	a1 := CreateVector(3, 2, 1)
	a2 := CreateVector(5, 6, 7)
	e := CreateVector(-2, -4, -6)

	r := a1.Subtract(a2)
	if !e.Equals(r) {
		t.Fatalf("%+v and %+v should subtract to %+v insted of %+v", a1, a2, e, r)
	}
}

func TestSubtractVectorFromZero(t *testing.T) {
	a1 := CreateVector(0, 0, 0)
	a2 := CreateVector(1, -2, 3)
	e := CreateVector(-1, 2, -3)

	r := a1.Subtract(a2)
	if !e.Equals(r) {
		t.Fatalf("%+v and %+v should subtract to %+v insted of %+v", a1, a2, e, r)
	}
}

func TestNegate(t *testing.T) {
	tup := Tupple{1, -2, 3, -4}
	e := Tupple{-1, 2, -3, 4}

	r := tup.Negate()
	if !e.Equals(r) {
		t.Fatalf("%+v should negate to %+v insted of %+v", tup, e, r)
	}
}

func TestMultiplyByScalar(t *testing.T) {
	tup := Tupple{1, -2, 3, -4}
	scalar := 3.5
	expected := Tupple{3.5, -7, 10.5, -14}

	r := tup.Multiply(scalar)
	if !expected.Equals(r) {
		t.Fatalf("%+v should multiply to %+v insted of %+v", tup, expected, r)
	}
}

func TestMultiplyByFraction(t *testing.T) {
	tup := Tupple{1, -2, 3, -4}
	fraction := 0.5
	expected := Tupple{0.5, -1, 1.5, -2}

	r := tup.Multiply(fraction)
	if !expected.Equals(r) {
		t.Fatalf("%+v should multiply to %+v insted of %+v", tup, expected, r)
	}
}

func TestDivide(t *testing.T) {
	tup := Tupple{1, -2, 3, -4}
	divisor := 2.0
	expected := Tupple{0.5, -1, 1.5, -2}

	r := tup.Divide(divisor)
	if !expected.Equals(r) {
		t.Fatalf("%+v should divide to %+v insted of %+v", tup, expected, r)
	}
}

func TestMagnitude(t *testing.T) {
	v := CreateVector(1, 0, 0)
	em := 1.0

	r := v.Magnitude()
	if !SafeEquals(em, r) {
		t.Fatalf("Magnitude of %+v should be %f insted of %f", v, em, r)
	}

	v = CreateVector(0, 1, 0)
	em = 1.0

	r = v.Magnitude()
	if !SafeEquals(em, r) {
		t.Fatalf("Magnitude of %+v should be %f insted of %f", v, em, r)
	}

	v = CreateVector(0, 0, 1)
	em = 1.0

	r = v.Magnitude()
	if !SafeEquals(em, r) {
		t.Fatalf("Magnitude of %+v should be %f insted of %f", v, em, r)
	}

	v = CreateVector(1, 2, 3)
	em = 3.741657

	r = v.Magnitude()
	if !SafeEquals(em, r) {
		t.Fatalf("Magnitude of %+v should be %f insted of %f", v, em, r)
	}

	v = CreateVector(-1, -2, -3)
	em = 3.741657

	r = v.Magnitude()
	if !SafeEquals(em, r) {
		t.Fatalf("Magnitude of %+v should be %f insted of %f", v, em, r)
	}
}

func TestNormalize(t *testing.T) {
	v := CreateVector(4, 0, 0)
	e := CreateVector(1, 0, 0)

	r := v.Normalize()
	if !e.Equals(r) {
		t.Fatalf("Normalized vector of %+v should be %+v insted of %+v", v, e, r)
	}

	v = CreateVector(1, 2, 3)
	e = CreateVector(0.26726, 0.53452, 0.80178)

	r = v.Normalize()
	if !e.Equals(r) {
		t.Fatalf("Normalized vector of %+v should be %+v insted of %+v", v, e, r)
	}
}

func TestNormalizeMagnitude(t *testing.T) {
	v := CreateVector(1, 2, 3)
	n := v.Normalize()
	em := 1.0

	r := n.Magnitude()
	if !SafeEquals(r, em) {
		t.Fatalf("Magnitude of normalized vector %+v should be %f insted of %f", v, em, r)
	}
}

func TestDotProduct(t *testing.T) {
	a := CreateVector(1, 2, 3)
	b := CreateVector(2, 3, 4)

	expected := 20.0
	actual := a.Dot(b)

	if !SafeEquals(expected, actual) {
		t.Fatalf("Dot product of vectors %+v and %+v should be %f insted of %f", a, b, expected, actual)
	}
}

func TestCrossProduct(t *testing.T) {
	a := CreateVector(1, 2, 3)
	b := CreateVector(2, 3, 4)

	expected := CreateVector(-1, 2, -1)
	actual := a.Cross(b)
	if !expected.Equals(actual) {
		t.Fatalf("Cross product of vectors %+v and %+v should be %f insted of %f", a, b, expected, actual)
	}

	expected = CreateVector(1, -2, 1)
	actual = b.Cross(a)
	if !expected.Equals(actual) {
		t.Fatalf("Cross product of vectors %+v and %+v should be %f insted of %f", b, a, expected, actual)
	}
}
