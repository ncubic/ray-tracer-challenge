package types

import "math"

const EPSILON float64 = 0.00001

type Tupple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (t Tupple) IsVector() bool {
	if SafeEquals(t.W, 0.0) {
		return true
	} else {
		return false
	}
}

func (t Tupple) IsPoint() bool {
	return SafeEquals(t.W, 1.0)
}

func (t Tupple) Equals(o Tupple) bool {
	ok := SafeEquals(t.X, o.X) && SafeEquals(t.Y, o.Y) && SafeEquals(t.Z, o.Z) && SafeEquals(t.W, o.W)
	return ok
}

func (t Tupple) Add(o Tupple) Tupple {
	return Tupple{
		t.X + o.X,
		t.Y + o.Y,
		t.Z + o.Z,
		t.W + o.W,
	}
}

func (t Tupple) Subtract(o Tupple) Tupple {
	return Tupple{
		t.X - o.X,
		t.Y - o.Y,
		t.Z - o.Z,
		t.W - o.W,
	}
}

func (t Tupple) Negate() Tupple {
	return Tupple{
		t.X * -1,
		t.Y * -1,
		t.Z * -1,
		t.W * -1,
	}
}

func (t Tupple) Multiply(scalar float64) Tupple {
	return Tupple{
		t.X * scalar,
		t.Y * scalar,
		t.Z * scalar,
		t.W * scalar,
	}
}

func (t Tupple) Divide(divisor float64) Tupple {
	return Tupple{
		t.X / divisor,
		t.Y / divisor,
		t.Z / divisor,
		t.W / divisor,
	}
}

func (t Tupple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

func (t Tupple) Normalize() Tupple {
	magnitude := t.Magnitude()
	return Tupple{
		t.X / magnitude,
		t.Y / magnitude,
		t.Z / magnitude,
		t.W / magnitude,
	}
}

func (t Tupple) Dot(o Tupple) float64 {
	return t.X*o.X +
		t.Y*o.Y +
		t.Z*o.Z +
		t.W*o.W
}

func (t Tupple) Cross(o Tupple) Tupple {
	return CreateVector(
		t.Y*o.Z-t.Z*o.Y,
		t.Z*o.X-t.X*o.Z,
		t.X*o.Y-t.Y*o.X,
	)
}

func CreatePoint(x, y, z float64) Tupple {
	return Tupple{x, y, z, 1.0}
}

func CreateVector(x, y, z float64) Tupple {
	return Tupple{x, y, z, 0.0}
}

func SafeEquals(n1, n2 float64) bool {
	if math.Abs(n1-n2) < EPSILON {
		return true
	} else {
		return false
	}
}
