package main

import (
	"rtc/types"
	"strings"
	"testing"
)

func TestToStringHeader(t *testing.T) {
	c := types.CreateCanvas(5, 3)
	s := ToString(c)

	lines := strings.Split(s, "\n")

	if len(lines) < 3 {
		t.Fatalf("Invalid line count %d", len(lines))
	}

	if lines[0] != "P3" {
		t.Fatalf("Invalid first line %s", lines[0])
	}

	if lines[1] != "5 3" {
		t.Fatalf("Invalid second line %s", lines[1])
	}

	if lines[2] != "255" {
		t.Fatalf("Invalid third line %s", lines[2])
	}
}

func TestToStringBody(t *testing.T) {
	c := types.CreateCanvas(5, 3)
	c.WritePixel(0, 0, types.CreateColor(1.5, 0, 0))
	c.WritePixel(2, 1, types.CreateColor(0, 0.5, 0))
	c.WritePixel(4, 2, types.CreateColor(-0.5, 0, 1))

	s := ToString(c)
	lines := strings.Split(s, "\n")

	if len(lines) != 6 {
		t.Fatalf("Invalid line count %d", len(lines))
	}

	if lines[3] != "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0" {
		t.Fatalf("Invalid first body line %s", lines[3])
	}

	if lines[4] != "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0" {
		t.Fatalf("Invalid second body line %s", lines[4])
	}

	if lines[5] != "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255" {
		t.Fatalf("Invalid third body line %s", lines[5])
	}
}
