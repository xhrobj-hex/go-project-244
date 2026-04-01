package parser

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseJSON(t *testing.T) {
	data := `{
  		"host": "hexlet.io",
  		"timeout": 50,
  		"proxy": "123.234.53.22",
  		"follow": false
	}`

	var want map[string]any
	err := json.Unmarshal([]byte(data), &want)
	if err != nil {
		t.Fatalf("prepare test data: %v", err)
	}

	dir := t.TempDir()
	path := filepath.Join(dir, "test.json")

	filedata := []byte(data)

	err = os.WriteFile(path, filedata, 0644)
	if err != nil {
		t.Fatalf("write file: %v", err)
	}

	got, err := Parse(path)
	if err != nil {
		t.Fatalf("parse data: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want: %#v, got: %#v", want, got)
	}
}

func TestParseYAML(t *testing.T) {
	data := `
host: hexlet.io
timeout: 50
proxy: 123.234.53.22
follow: false
`

	want := map[string]any{
		"host":    "hexlet.io",
		"timeout": 50,
		"proxy":   "123.234.53.22",
		"follow":  false,
	}

	dir := t.TempDir()
	path := filepath.Join(dir, "test.yaml")

	err := os.WriteFile(path, []byte(data), 0644)
	if err != nil {
		t.Fatalf("write file: %v", err)
	}

	got, err := Parse(path)
	if err != nil {
		t.Fatalf("parse yaml: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want: %#v, got: %#v", want, got)
	}
}

func TestParseUnsupportedExtension(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")

	filedata := []byte("")
	err := os.WriteFile(path, filedata, 0644)
	if err != nil {
		t.Fatalf("write file: %v", err)
	}

	_, err = Parse(path)
	if err == nil {
		t.Fatalf("want err, got nil")
	}
}

func TestParse_DirectoryInsteadOfFile(t *testing.T) {
	parentDir := t.TempDir()
	dirPath := filepath.Join(parentDir, "data.json")

	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		t.Fatalf("create directory: %v", err)
	}

	_, err = Parse(dirPath)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
