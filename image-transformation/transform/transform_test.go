package transform

import (
	"io"
	"os"
	"testing"
)

func TestTransform(t *testing.T) {
	var file io.Reader
	file, err := os.Open("input.png")
	file, err = Transform(file, "png", "1", "100")
	if err != nil {
		t.Errorf("error")
	}
}

func TestPrimitive(t *testing.T) {
	_, err := primitive("-i in.png -o out.png -n 10 -m 0", "out.png")
	if err == nil {
		t.Error("error in primitive")
	}
	_, err = primitive("-i input.png -o output.png -n 10 -m 0", "out.png")
	if err == nil {
		t.Error("error in primitive")
	}
}

func TestTempfile(t *testing.T) {
	_, err := tempfile("", "")
	if err != nil {
		t.Error("error in tempfile")
	}
}
