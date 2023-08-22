package types

import "testing"

func TestCreateCanvas(t *testing.T) {
	canvas := CreateCanvas(10, 20)

	if canvas.Width != 10 {
		t.Fatalf("Invalid width %d, expected %d", canvas.Width, 10)
	}

	if canvas.Height != 20 {
		t.Fatalf("Invalid height %d, expected %d", canvas.Height, 20)
	}

	expected := CreateColor(0, 0, 0)
	for x := range canvas.pixels {
		for y := range canvas.pixels[x] {
			color := canvas.pixels[x][y]

			if !color.Equals(expected) {
				t.Fatalf("Invalid color value at %d and %d", x, y)
			}
		}
	}
}

func TestWritePixel(t *testing.T) {
	canvas := CreateCanvas(10, 20)
	red := CreateColor(1, 0, 0)

	canvas.WritePixel(2, 3, red)
	actual := canvas.PixelAt(2, 3)

	if !actual.Equals(red) {
		t.Fatalf("Invalid color %+v", actual)
	}
}
