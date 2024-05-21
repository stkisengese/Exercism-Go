package sorting

import (
	"testing"
)

func TestDescribeNumber(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  string
	}{
		{
			name: "Describe 4.1", input: 4.1, want: "This is the number 4.1",
		},
		{
			name: "Describe -3.2", input: -3.2, want: "This is the number -3.2",
		},
		{
			name: "Pads to single decimal place", input: 4.0, want: "This is the number 4.0",
		},
		{
			name: "Truncates to single decimal place", input: 7.11, want: "This is the number 7.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DescribeNumber(tt.input); got != tt.want {
				t.Errorf("DescribeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testNumberBox struct {
	n int
}

func (nb testNumberBox) Number() int {
	return nb.n
}

func TestDescribeNumberBox(t *testing.T) {
	tests := []struct {
		name  string
		input NumberBox
		want  string
	}{
		{"Describe NumberBox with 4", testNumberBox{4}, "This is a box containing the number 4.0"},
		{"Describe NumberBox with -3", testNumberBox{-3}, "This is a box containing the number -3.0"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := DescribeNumberBox(test.input); got != test.want {
				t.Errorf("DescribeNumberBox(%v) = %v want %v", test.input, got, test.want)
			}
		})
	}
}

type differentFancyNumber struct {
	num string
}

func (i differentFancyNumber) Value() string {
	return i.num
}

func TestExtractFancyNumber(t *testing.T) {
	tests := []struct {
		name  string
		input FancyNumberBox
		want  int
	}{
		{
			name: "Extract fancy number 11", input: FancyNumber{"11"}, want: 11,
		},
		{
			name: "Extract fancy number 0", input: FancyNumber{"0"}, want: 0,
		},
		{
			name: "Extract a different fancy number returns 0", input: differentFancyNumber{"two"}, want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractFancyNumber(tt.input); got != tt.want {
				t.Errorf("ExtractFancyNumber(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestDescribeFancyNumberBox(t *testing.T) {
	tests := []struct {
		name  string
		input FancyNumberBox
		want  string
	}{
		{
			name: "Describe fancy number 12", input: FancyNumber{"12"}, want: "This is a fancy box containing the number 12.0",
		},
		{
			name: "Describe fancy number 0", input: FancyNumber{"0"}, want: "This is a fancy box containing the number 0.0",
		},
		{
			name: "Describe a different fancy number", input: differentFancyNumber{"three"}, want: "This is a fancy box containing the number 0.0",
		},
		{
			name: "Describe a valid different fancy number", input: differentFancyNumber{"4"}, want: "This is a fancy box containing the number 0.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DescribeFancyNumberBox(tt.input); got != tt.want {
				t.Errorf("DescribeFancyNumberBox(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestDescribeAnything(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  string
	}{
		{
			name: "Describe 7.2", input: 7.2, want: "This is the number 7.2",
		},
		{
			name: "Describe 42", input: 42, want: "This is the number 42.0",
		},
		{
			name: "Describe NumberBox with 16", input: testNumberBox{16}, want: "This is a box containing the number 16.0",
		},
		{
			name: "Describe FancyNumber with 16", input: FancyNumber{"16"}, want: "This is a fancy box containing the number 16.0",
		},
		{
			name: "Describe a different FancyNumberBox", input: differentFancyNumber{"ten"}, want: "This is a fancy box containing the number 0.0",
		},
		{
			name: "Something unknown is labelled return to sender", input: "something we did not anticipate", want: "Return to sender",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DescribeAnything(tt.input); got != tt.want {
				t.Errorf("DescribeAnything(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
