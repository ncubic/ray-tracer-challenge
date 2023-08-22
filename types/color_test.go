package types

import "testing"

func TestColorValues(t *testing.T) {
	color := CreateColor(-0.5, 0.4, 1.7)

	if !SafeEquals(color.Red(), -0.5) {
		t.Fatalf("Invalid value of red")
	}

	if !SafeEquals(color.Green(), 0.4) {
		t.Fatalf("Invalid value of blue")
	}

	if !SafeEquals(color.Blue(), 1.7) {
		t.Fatalf("Invalid value of green")
	}
}

func TestColorAdd(t *testing.T) {
	c1 := CreateColor(0.9, 0.6, 0.75)
	c2 := CreateColor(0.7, 0.1, 0.25)

	expected := CreateColor(1.6, 0.7, 1.0)
	actual := c1.Add(c2)

	if !expected.Equals(actual) {
		t.Fatalf("%+v and %+v should add to %+v instead of %+v", c1, c2, expected, actual)
	}
}

func TestColorSubtract(t *testing.T) {
	c1 := CreateColor(0.9, 0.6, 0.75)
	c2 := CreateColor(0.7, 0.1, 0.25)

	expected := CreateColor(0.2, 0.5, 0.5)
	actual := c1.Subtract(c2)

	if !expected.Equals(actual) {
		t.Fatalf("%+v and %+v should subtract to %+v instead of %+v", c1, c2, expected, actual)
	}
}

func TestColorMultiplyByScalar(t *testing.T) {
	c1 := CreateColor(0.2, 0.3, 0.4)
	scalar := 2.0

	expected := CreateColor(0.4, 0.6, 0.8)
	actual := c1.MultiplyByScalar(scalar)

	if !expected.Equals(actual) {
		t.Fatalf("%+v multiplied by %f should be %+v instead of %+v", c1, scalar, expected, actual)
	}
}

func TestMultiply(t *testing.T) {
	c1 := CreateColor(1, 0.2, 0.4)
	c2 := CreateColor(0.9, 1, 0.1)

	expected := CreateColor(0.9, 0.2, 0.04)
	actual := c1.Multiply(c2)

	if !expected.Equals(actual) {
		t.Fatalf("%+v and %+v should multiply to %+v instead of %+v", c1, c2, expected, actual)
	}
}
