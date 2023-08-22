package types

type Color struct {
	tupple Tupple
}

func (c Color) Red() float64 {
	return c.tupple.X
}

func (c Color) Green() float64 {
	return c.tupple.Y
}

func (c Color) Blue() float64 {
	return c.tupple.Z
}

func CreateColor(red, green, blue float64) Color {
	tupple := Tupple{red, green, blue, 0.0}
	return Color{tupple}
}

func (c Color) Add(o Color) Color {
	return Color{c.tupple.Add(o.tupple)}
}

func (c Color) Subtract(o Color) Color {
	return Color{c.tupple.Subtract(o.tupple)}
}

func (c Color) MultiplyByScalar(scalar float64) Color {
	return Color{c.tupple.Multiply(scalar)}
}

func (c Color) Multiply(o Color) Color {
	return CreateColor(
		c.Red()*o.Red(),
		c.Green()*o.Green(),
		c.Blue()*o.Blue(),
	)
}

func (c Color) Equals(o Color) bool {
	return c.tupple.Equals(o.tupple)
}
