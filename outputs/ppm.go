package main

import (
	"fmt"
	"math"
	t "rtc/types"
	"strconv"
	"strings"
)

var MAX_VALUE = 255
var MAX_LINE_LENGTH = 70

func getHeaderString(c t.Canvas) string {
	return fmt.Sprintf("P3\n%d %d\n%d", c.Width, c.Height, MAX_VALUE)
}

func scaleToInt(v float64) int {
	av := v

	if av < 0 {
		av = 0
	}

	if av > 1 {
		av = 1
	}

	return int(math.RoundToEven(av * float64(MAX_VALUE)))
}

func formatValue(v float64) string {
	return strconv.Itoa(scaleToInt(v))
}

func getBodyString(c t.Canvas) string {
	lines := make([]string, 0, c.Height)

	for y := 0; y < c.Height; y++ {
		values := make([]string, c.Width*3)
		for x := 0; x < c.Width*3; x += 3 {
			pixel := c.PixelAt(x/3, y)
			values[x] = formatValue(pixel.Red())
			values[x+1] = formatValue(pixel.Green())
			values[x+2] = formatValue(pixel.Blue())
		}

		newLines := appendToLines(values)
		lines = append(lines, newLines...)
	}

	return strings.Join(lines, "\n")
}

func appendToLines(values []string) []string {
	currentLength := -1 // initial value won't have space before it
	buffer := make([]string, 0, MAX_LINE_LENGTH/3)
	lines := make([]string, 0, len(values)*3/MAX_LINE_LENGTH)

	SPACE_LENGTH := 1

	for _, v := range values {

		expectedLength := currentLength + len(v) + SPACE_LENGTH

		if expectedLength > MAX_LINE_LENGTH {
			lines = append(lines, strings.Join(buffer, " "))
			buffer = make([]string, 0, MAX_LINE_LENGTH/3)
			currentLength = 0
		}

		buffer = append(buffer, v)
		currentLength += len(v) + 1
	}

	if currentLength > 0 {
		lines = append(lines, strings.Join(buffer, " "))
	}

	return lines
}

func ToString(c t.Canvas) string {
	return fmt.Sprintf("%s\n%s\n", getHeaderString(c), getBodyString(c))
}
