package code

import (
	"code/internal/parser"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// GenDiff вычисляет различия между двумя файлами и возвращает их
// в виде строки в указанном формате.
func GenDiff(leftPath, rightPath, format string) (string, error) {
	_ = format // NOTE: пока не используем ...

	leftData, rightData, err := parseFiles(leftPath, rightPath)
	if err != nil {
		return "", err
	}

	keys := sortedKeys(leftData, rightData)
	diff := buildDiff(leftData, rightData, keys)

	// ... пока просто "склеить diff в итоговую строку"
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
		leftValue, leftOK := leftData[key]
		rightValue, rightOK := rightData[key]

		switch {
		case !leftOK:
			diff = append(diff, formatLine("  +", key, rightValue))
		case !rightOK:
			diff = append(diff, formatLine("  -", key, leftValue))
		case !reflect.DeepEqual(leftValue, rightValue):
			diff = append(diff, formatLine("  -", key, leftValue))
			diff = append(diff, formatLine("  +", key, rightValue))
		default:
			diff = append(diff, formatLine("   ", key, leftValue))
		}
	}

	return diff
}

func formatLine(prefix, key string, value any) string {
	return fmt.Sprintf("%s %s: %v", prefix, key, value)
}
