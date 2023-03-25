package html

import (
	"os"
	"testing"
)

func TestCleanHTML(t *testing.T) {
	content, _ := os.ReadFile("../testdata/dirty-page.html")

	want, _ := os.ReadFile("../testdata/clean-page.html")
	got := CleanHTML(string(content))

	if string(want) != got {
		t.Errorf("got %q want %q", got, want)
	}
}
