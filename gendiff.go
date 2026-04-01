package code

import (
	"code/internal/parser"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func GenDiff(leftPath, rightPath, format string) (string, error) {
	_ = format // NOTE: пока не используем

	leftData, rightData, err := parseFiles(leftPath, rightPath)
	if err != nil {
		return "", err
	}

	keysSorted := sortedKeys(leftData, rightData)
	diff := buildDiff(leftData, rightData, keysSorted)

	// склеить diff в итоговую строку
	r := fmt.Sprintf("{\n%s\n}", strings.Join(diff, "\n"))

	return r, nil
}

func parseFiles(path1, path2 string) (map[string]any, map[string]any, error) {
	data1, err := parser.Parse(path1)
	if err != nil {
		return nil, nil, err
	}

	data2, err := parser.Parse(path2)
	if err != nil {
		return nil, nil, err
	}

	return data1, data2, nil
}

func sortedKeys(data1, data2 map[string]any) []string {
	set := make(map[string]struct{})
	for k := range data1 {
		set[k] = struct{}{}
	}
	for k := range data2 {
		set[k] = struct{}{}
	}

	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

func buildDiff(leftData, rightData map[string]any, keys []string) []string {
	diff := make([]string, 0, len(keys))

	for _, key := range keys {
		v1, ok1 := leftData[key]
		v2, ok2 := rightData[key]

		switch {
		case !ok1:
			d := fmt.Sprintf("  + %s: %v", key, v2)
			diff = append(diff, d)
		case !ok2:
			d := fmt.Sprintf("  - %s: %v", key, v1)
			diff = append(diff, d)
		case !reflect.DeepEqual(v1, v2):
			d := fmt.Sprintf("  - %s: %v", key, v1)
			diff = append(diff, d)

			d = fmt.Sprintf("  + %s: %v", key, v2)
			diff = append(diff, d)
		default:
			d := fmt.Sprintf("    %s: %v", key, v1)
			diff = append(diff, d)
		}
	}

	return diff
}
