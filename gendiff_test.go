package code

import (
	"path/filepath"
	"testing"
)

func TestGenDiffFlatJSON(t *testing.T) {
	file1 := filepath.Join("testdata", "fixture", "file1.json")
	file2 := filepath.Join("testdata", "fixture", "file2.json")

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

func TestGenDiffFlatJSONSameDataDifferentKeyOrder(t *testing.T) {
	file1 := filepath.Join("testdata", "fixture", "file3.json")
	file2 := filepath.Join("testdata", "fixture", "file4.json")

	got, err := GenDiff(file1, file2, "stylish")
	if err != nil {
		t.Fatalf("GenDiff returned error: %v", err)
	}

	want := `{
    follow: false
    host: hexlet.io
    timeout: 50
}`

	if got != want {
		t.Errorf("unexpected diff result\nwant:\n%s\ngot:\n%s", want, got)
	}
}
