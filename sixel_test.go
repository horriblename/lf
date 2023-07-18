package main

import (
	"strings"
	"testing"
)

func TestSixelSize(t *testing.T) {
	tests := []struct {
		si     string
		hi     int
		succ   bool
		wo, ho int
	}{
		{
			"\x1bP;;;q\"1;1;1;12#0;2;97;97;97#1;2;75;75;75#0B$#1{\x1b\\", 12,
			true, 1, 6,
		},
		{
			"\x1bP;;;q\"1;1;1;12#0;2;97;97;97#1;2;75;75;75#0B$#1{-#0o$#1N\x1b\\", 30,
			true, 1, 12,
		},
		{
			"\x1bP;;;q\"1;1;1;12#0;2;97;97;97#1;2;75;75;75#0B$#1{-#0o$#1N\x1b\\", 12,
			true, 1, 12,
		},
		{
			"\x1bP;;;q\"1;1;1;12#0;2;97;97;97#1;2;75;75;75#0B$#1{-#0o$#1N\x1b\\", 11,
			false, 0, 0,
		},
		{
			"\x1bP;;;q\"1;1;1;12#0;2;97;97;97#1;2;75;75;75#0B$#1{-#0o$#1N-\x1b\\", 30,
			true, 1, 12,
		},
		{
			"\x1bP;;;q\"1;1;1;12#0;2;97;97;97#1;2;75;75;75#0B$#1{-#0o$#1N-\x1b\\", 12,
			true, 1, 12,
		},
		{
			"\x1bP;;;q\"1;1;1;12#0;2;97;97;97#1;2;75;75;75#0B$#1{-#0o$#1N-\x1b\\", 11,
			false, 0, 0,
		},
	}

	for i, test := range tests {
		reader := strings.NewReader(test.si)
		w, h, err := sixelSize(reader, test.hi)

		if !test.succ {
			if err == nil {
				t.Errorf("test #%d expected to fail", i)
			}
			continue
		} else if err != nil {
			t.Errorf("test #%d failed with error %s", i, err)
			continue
		}

		if w != test.wo || h != test.ho {
			t.Errorf("test #%d expected (%d, %d), got (%d, %d)", i, test.wo, test.ho, w, h)
		}
	}

}
