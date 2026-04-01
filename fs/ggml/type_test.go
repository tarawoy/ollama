package ggml

import "testing"

func TestParseFileTypeTurboQuantAliases(t *testing.T) {
	tests := []struct {
		in   string
		want FileType
	}{
		{in: "tq1_0", want: fileTypeTQ1_0},
		{in: "TQ2_0", want: fileTypeTQ2_0},
		{in: "turboquant google", want: fileTypeTQ2_0},
		{in: "google-turboquant", want: fileTypeTQ2_0},
	}

	for _, tt := range tests {
		got, err := ParseFileType(tt.in)
		if err != nil {
			t.Fatalf("ParseFileType(%q) unexpected error: %v", tt.in, err)
		}
		if got != tt.want {
			t.Fatalf("ParseFileType(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestParseTensorTypeTurboQuantAliases(t *testing.T) {
	tests := []struct {
		in   string
		want TensorType
	}{
		{in: "tq1", want: tensorTypeTQ1_0},
		{in: "TQ1_0", want: tensorTypeTQ1_0},
		{in: "tq2", want: tensorTypeTQ2_0},
		{in: "turboquant google", want: tensorTypeTQ2_0},
	}

	for _, tt := range tests {
		got, err := ParseTensorType(tt.in)
		if err != nil {
			t.Fatalf("ParseTensorType(%q) unexpected error: %v", tt.in, err)
		}
		if got != tt.want {
			t.Fatalf("ParseTensorType(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
