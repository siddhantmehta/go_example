package sqrt

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestMany(t *testing.T) {
	file, err := os.Open("sqrt_cases.csv")
	if err != nil {
		t.Fatalf("Error: can't open the file - %s\n", err)
	}
	defer file.Close()
	lines := csv.NewReader(file)
	for {
		line, err := lines.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("error reading cases file - %s", err)
		}
		val, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			t.Fatalf("bad value - %s", err)
		}
		expected, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			t.Fatalf("bad value - %s", err)
		}
		t.Run(fmt.Sprintf("%f", val), func(t *testing.T) {
			out, err := Sqrt(val)
			if err != nil {
				t.Fatal(err)
			}
			if !almostEqual(out, expected) {
				t.Fatalf("%f != %f", out, expected)
			}
		})
	}
}
