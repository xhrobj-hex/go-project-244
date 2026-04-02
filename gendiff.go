package code

import (
	"code/internal/parser"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type diffNode struct {
	Key      string
	Kind     string // NOTE: right/left/both/tie/deep
	Left     any
	Right    any
	Children []diffNode
}

// GenDiff вычисляет различия между двумя файлами и возвращает их
// в виде строки в указанном формате.
func GenDiff(leftPath, rightPath, format string) (string, error) {
	leftData, rightData, err := parseFiles(leftPath, rightPath)
	if err != nil {
		return "", err
	}

	tree := buildDiff(leftData, rightData)

	return formatTree(tree, 1, format), nil
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

func buildDiff(leftData, rightData map[string]any) []diffNode {
	keys := sortedKeys(leftData, rightData)

	tree := make([]diffNode, 0, len(keys))

	for _, key := range keys {
		leftValue, leftOK := leftData[key]
		rightValue, rightOK := rightData[key]

		leftObj, leftIsObj := asMap(leftValue)
		rightObj, rightIsObj := asMap(rightValue)

		switch {
		case !leftOK:
			tree = append(tree, diffNode{
				Key:   key,
				Kind:  "right",
				Right: rightValue,
			})

		case !rightOK:
			tree = append(tree, diffNode{
				Key:  key,
				Kind: "left",
				Left: leftValue,
			})

		case leftIsObj && rightIsObj:
			tree = append(tree, diffNode{
				Key:      key,
				Kind:     "deep",
				Children: buildDiff(leftObj, rightObj),
			})

		case !reflect.DeepEqual(leftValue, rightValue):
			tree = append(tree, diffNode{
				Key:   key,
				Kind:  "both",
				Left:  leftValue,
				Right: rightValue,
			})

		default:
			tree = append(tree, diffNode{
				Key:  key,
				Kind: "tie",
				Left: leftValue, // NOTE: для "tie" храним только Left
			})
		}
	}

	return tree
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

func asMap(value any) (map[string]any, bool) {
	obj, ok := value.(map[string]any)
	return obj, ok
}

func formatTree(tree []diffNode, level int, format string) string {
	_ = format // NOTE: пока не используем ...

	lines := []string{"{"}

	for _, node := range tree {
		switch node.Kind {
		case "right":
			lines = append(lines,
				fmt.Sprintf(
					"%s+ %s: %s",
					nodePadding(level),
					node.Key,
					formatValue(node.Right, level+1),
				),
			)

		case "left":
			lines = append(lines,
				fmt.Sprintf(
					"%s- %s: %s",
					nodePadding(level),
					node.Key,
					formatValue(node.Left, level+1),
				),
			)

		case "both":
			lines = append(lines,
				fmt.Sprintf(
					"%s- %s: %s",
					nodePadding(level),
					node.Key,
					formatValue(node.Left, level+1),
				),
			)
			lines = append(lines,
				fmt.Sprintf(
					"%s+ %s: %s",
					nodePadding(level),
					node.Key,
					formatValue(node.Right, level+1),
				),
			)

		case "tie":
			lines = append(lines,
				fmt.Sprintf(
					"%s%s%s%s: %s",
					nodePadding(level),
					paddingChar(),
					paddingChar(),
					node.Key,
					formatValue(node.Left, level+1),
				),
			)

		case "deep":
			lines = append(lines,
				fmt.Sprintf(
					"%s%s%s%s: %s",
					nodePadding(level),
					paddingChar(),
					paddingChar(),
					node.Key,
					formatTree(node.Children, level+1, format),
				),
			)

		default:
			lines = append(lines, "**! error <--")
		}
	}

	lines = append(lines, fmt.Sprintf("%s}", closePadding(level)))

	return strings.Join(lines, "\n")
}

func sortedMapKeys(data map[string]any) []string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

func formatValue(value any, level int) string {
	switch v := value.(type) {
	case nil:
		return "null"

	case map[string]any:
		keys := sortedMapKeys(v)

		lines := []string{"{"}
		for _, key := range keys {
			lines = append(lines,
				fmt.Sprintf(
					"%s%s: %s",
					valuePadding(level),
					key,
					formatValue(v[key], level+1),
				),
			)
		}
		lines = append(lines, fmt.Sprintf("%s}", closePadding(level)))

		return strings.Join(lines, "\n")

	default:
		return fmt.Sprintf("%v", v)
	}
}

func nodePadding(level int) string {
	return strings.Repeat(paddingChar(), level*4-2)
}

func valuePadding(level int) string {
	return strings.Repeat(paddingChar(), level*4)
}

func closePadding(level int) string {
	return strings.Repeat(paddingChar(), (level-1)*4)
}

func paddingChar() string {
	return string('.')
}
