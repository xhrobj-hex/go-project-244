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
