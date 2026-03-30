package code

import (
	"path/filepath"
	"testing"
)

func TestGenDiff(t *testing.T) {
	file1 := filepath.Join("testdata", "file1.json")
	file2 := filepath.Join("testdata", "file2.json")

	got, err := GenDiff(file1, file2, "stylish")
	if err != nil {
		t.Fatalf("GenDiff returned error: %v", err)
	}

	want := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

	if got != want {
		t.Errorf("unexpected diff result\nwant:\n%s\ngot:\n%s", want, got)
	}
}
