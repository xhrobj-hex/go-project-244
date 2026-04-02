package code

import (
	"path/filepath"
	"testing"
)

func TestGenDiffFlatJSON(t *testing.T) {
	fileLeft := filepath.Join("testdata", "fixture", "file1.json")
	fileRight := filepath.Join("testdata", "fixture", "file2.json")

	got, err := GenDiff(fileLeft, fileRight, "stylish")
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

func TestGenDiffFlatYAML(t *testing.T) {
	fileLeft := filepath.Join("testdata", "fixture", "file1.yml")
	fileRight := filepath.Join("testdata", "fixture", "file2.yml")

	got, err := GenDiff(fileLeft, fileRight, "stylish")
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
	fileLeft := filepath.Join("testdata", "fixture", "file3.json")
	fileRight := filepath.Join("testdata", "fixture", "file4.json")

	got, err := GenDiff(fileLeft, fileRight, "stylish")
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

func TestGenDiffFlatYAMLSameDataDifferentKeyOrder(t *testing.T) {
	fileLeft := filepath.Join("testdata", "fixture", "file3.yml")
	fileRight := filepath.Join("testdata", "fixture", "file4.yml")

	got, err := GenDiff(fileLeft, fileRight, "stylish")
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

func TestGenDiffMultilevelJSON(t *testing.T) {
	fileLeft := filepath.Join("testdata", "fixture", "file5.json")
	fileRight := filepath.Join("testdata", "fixture", "file6.json")

	got, err := GenDiff(fileLeft, fileRight, "stylish")
	if err != nil {
		t.Fatalf("GenDiff returned error: %v", err)
	}

	want := `{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: null
      + setting4: blah blah
      + setting5: {
            key5: value5
        }
        setting6: {
            doge: {
              - wow: 
              + wow: so much
            }
            key: value
          + ops: vops
        }
    }
    group1: {
      - baz: bas
      + baz: bars
        foo: bar
      - nest: {
            key: value
        }
      + nest: str
    }
  - group2: {
        abc: 12345
        deep: {
            id: 45
        }
    }
  + group3: {
        deep: {
            id: {
                number: 45
            }
        }
        fee: 100500
    }
}`

	if got != want {
		t.Errorf("unexpected diff result\nwant:\n%s\ngot:\n%s", want, got)
	}
}
