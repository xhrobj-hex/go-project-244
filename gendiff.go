package code

import (
	"code/internal/parser"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type diff struct {
	Key   string
	Kind  string
	Left  any
	Right any
}

// GenDiff вычисляет различия между двумя файлами и возвращает их
// в виде строки в указанном формате.
func GenDiff(leftPath, rightPath, format string) (string, error) {
	leftData, rightData, err := parseFiles(leftPath, rightPath)
	if err != nil {
		return "", err
	}

	diffs := buildDiff(leftData, rightData)

	return formatDiffs(diffs, format), nil
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

func buildDiff(leftData, rightData map[string]any) []diff {
	keys := sortedKeys(leftData, rightData)

	diffs := make([]diff, 0, len(keys))

	for _, key := range keys {
		leftValue, leftOK := leftData[key]
		rightValue, rightOK := rightData[key]

		switch {
		case !leftOK:
			diffs = append(diffs, diff{
				Key:   key,
				Kind:  "right",
				Left:  nil,
				Right: rightValue,
			})
		case !rightOK:
			diffs = append(diffs, diff{
				Key:   key,
				Kind:  "left",
				Left:  leftValue,
				Right: nil,
			})
		case !reflect.DeepEqual(leftValue, rightValue):
			diffs = append(diffs, diff{
				Key:   key,
				Kind:  "both",
				Left:  leftValue,
				Right: rightValue,
			})

		default:
			diffs = append(diffs, diff{
				Key:   key,
				Kind:  "tie",
				Left:  leftValue,
				Right: nil, // NOTE: для "tie" храним только Left
			})
		}
	}

	return diffs
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

func formatDiffs(diffs []diff, format string) string {
	_ = format // NOTE: пока не используем ...

	r := make([]string, 0, len(diffs))

	for _, diff := range diffs {
		switch diff.Kind {
		case "right":
			r = append(r, formatLine("  +", diff.Key, diff.Right))
		case "left":
			r = append(r, formatLine("  -", diff.Key, diff.Left))
		case "both":
			r = append(r, formatLine("  -", diff.Key, diff.Left))
			r = append(r, formatLine("  +", diff.Key, diff.Right))
		case "tie":
			r = append(r, formatLine("   ", diff.Key, diff.Left))
		default:
			r = append(r, "  ! error <--")
		}
	}

	return fmt.Sprintf("{\n%s\n}", strings.Join(r, "\n"))
}

func formatLine(prefix, key string, value any) string {
	return fmt.Sprintf("%s %s: %v", prefix, key, value)
}
