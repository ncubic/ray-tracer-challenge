package types

type Canvas struct {
	Width  int
	Height int
	pixels [][]Color
}

func (c Canvas) WritePixel(x, y int, color Color) {
	if x >= c.Width {
		panic("Invalid X")
	}

	if y >= c.Height {
		panic("Invalid Y")
	}

	c.pixels[x][y] = color
}

func (c Canvas) PixelAt(x, y int) Color {
	if x >= c.Width {
		panic("Invalid X")
	}

	if y >= c.Height {
		panic("Invalid Y")
	}

	return c.pixels[x][y]
}

func CreateCanvas(width, height int) Canvas {
	pixels := make([][]Color, width)

	for column := range pixels {
		pixels[column] = make([]Color, height)

		for row := range pixels[column] {
			pixels[column][row] = CreateColor(0.0, 0.0, 0.0)
		}
	}

	return Canvas{
		width,
		height,
		pixels,
	}
}
