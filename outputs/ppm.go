package main

import (
	"fmt"
	"math"
	t "rtc/types"
	"strconv"
	"strings"
)

var MAX_VALUE = 255

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
	lines := make([]string, c.Height)

	for y := 0; y < c.Height; y++ {
		values := make([]string, c.Width*3)
		for x := 0; x < c.Width*3; x += 3 {
			pixel := c.PixelAt(x/3, y)
			values[x] = formatValue(pixel.Red())
			values[x+1] = formatValue(pixel.Green())
			values[x+2] = formatValue(pixel.Blue())
		}
		line := strings.Join(values, " ")
		lines[y] = line
	}

	return strings.Join(lines, "\n")
}

func ToString(c t.Canvas) string {
	return fmt.Sprintf("%s\n%s", getHeaderString(c), getBodyString(c))
}
