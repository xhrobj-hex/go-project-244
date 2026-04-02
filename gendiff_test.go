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

func TestGenDiffJSONNested(t *testing.T) {
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

func TestGenDiffJSONPlain(t *testing.T) {
	fileLeft := filepath.Join("testdata", "fixture", "file5.json")
	fileRight := filepath.Join("testdata", "fixture", "file6.json")

	got, err := GenDiff(fileLeft, fileRight, "plain")
	if err != nil {
		t.Fatalf("GenDiff returned error: %v", err)
	}

	want := `Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]`

	if got != want {
		t.Errorf("unexpected plain diff result\nwant:\n%s\ngot:\n%s", want, got)
	}
}

func TestGenDiffYAMLPlain(t *testing.T) {
	fileLeft := filepath.Join("testdata", "fixture", "file5.yml")
	fileRight := filepath.Join("testdata", "fixture", "file6.yml")

	got, err := GenDiff(fileLeft, fileRight, "plain")
	if err != nil {
		t.Fatalf("GenDiff returned error: %v", err)
	}

	want := `Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]`

	if got != want {
		t.Errorf("unexpected plain diff result\nwant:\n%s\ngot:\n%s", want, got)
	}
}
