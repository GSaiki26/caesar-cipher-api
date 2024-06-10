package utils

import (
	"testing"
)

// Functions
func TestConvertToCesar(t *testing.T) {
	// Create a sub test to check if ABC is converted to BCD
	t.Run("'ABC' to 'BCD'", func(t *testing.T) {
		got := ConvertToCesar("ABC", 1)
		want := "BCD"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("'ABC' to 'XYZ'", func(t *testing.T) {
		got := ConvertToCesar("ABC", 23)
		want := "XYZ"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("'Hello World!' to 'Ifmmp Xpsme!'", func(t *testing.T) {
		got := ConvertToCesar("Hello World!", 1)
		want := "Ifmmp Xpsme!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
