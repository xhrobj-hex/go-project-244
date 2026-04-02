package code

import (
	"code/internal/parser"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

const (
	indentChar  = " "
	indentWidth = 4
)

type nodeKind string

const (
	kindAdded     nodeKind = "added"
	kindRemoved   nodeKind = "removed"
	kindChanged   nodeKind = "changed"
	kindUnchanged nodeKind = "unchanged"
	kindNested    nodeKind = "nested"
)

type diffNode struct {
	Key      string
	Kind     nodeKind
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
	rootDepth := 1

	return formatTree(tree, rootDepth, format), nil
}

func parseFiles(leftPath, rightPath string) (map[string]any, map[string]any, error) {
	letfData, err := parser.Parse(leftPath)
	if err != nil {
		return nil, nil, fmt.Errorf("parse left file: %w", err)
	}

	rightData, err := parser.Parse(rightPath)
	if err != nil {
		return nil, nil, fmt.Errorf("parse right file: %w", err)
	}

	return letfData, rightData, nil
}

func buildDiff(leftData, rightData map[string]any) []diffNode {
	keys := sortedUnionKeys(leftData, rightData)

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
				Kind:  kindAdded,
				Right: rightValue,
			})

		case !rightOK:
			tree = append(tree, diffNode{
				Key:  key,
				Kind: kindRemoved,
				Left: leftValue,
			})

		case leftIsObj && rightIsObj:
			tree = append(tree, diffNode{
				Key:      key,
				Kind:     kindNested,
				Children: buildDiff(leftObj, rightObj),
			})

		case !reflect.DeepEqual(leftValue, rightValue):
			tree = append(tree, diffNode{
				Key:   key,
				Kind:  kindChanged,
				Left:  leftValue,
				Right: rightValue,
			})

		default:
			tree = append(tree, diffNode{
				Key:  key,
				Kind: kindUnchanged,
				Left: leftValue, // NOTE: для "unchanged" храним только Left
			})
		}
	}

	return tree
}

func sortedUnionKeys(data1, data2 map[string]any) []string {
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

func formatTree(tree []diffNode, depth int, format string) string {
	_ = format // NOTE: пока не используем ...

	lines := []string{"{"}

	for _, node := range tree {
		switch node.Kind {
		case kindAdded:
			lines = append(lines, formatLine(depth, "+", node.Key, formatValue(node.Right, depth+1)))

		case kindRemoved:
			lines = append(lines, formatLine(depth, "-", node.Key, formatValue(node.Left, depth+1)))

		case kindChanged:
			lines = append(lines,
				formatLine(depth, "-", node.Key, formatValue(node.Left, depth+1)),
				formatLine(depth, "+", node.Key, formatValue(node.Right, depth+1)),
			)

		case kindUnchanged:
			lines = append(lines, formatLine(depth, " ", node.Key, formatValue(node.Left, depth+1)))

		case kindNested:
			lines = append(lines, formatLine(depth, " ", node.Key, formatTree(node.Children, depth+1, format)))

		default:
			panic(fmt.Sprintf("(о_0) unknown diff node kind: %q", node.Kind))
		}
	}

	lines = append(lines, fmt.Sprintf("%s}", closingIndent(depth)))

	return strings.Join(lines, "\n")
}

func formatLine(depth int, sign, key, value string) string {
	return fmt.Sprintf("%s%s %s: %s", nodeIndent(depth), sign, key, value)
}

func formatValue(value any, depth int) string {
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
					valueIndent(depth),
					key,
					formatValue(v[key], depth+1),
				),
			)
		}
		lines = append(lines, fmt.Sprintf("%s}", closingIndent(depth)))

		return strings.Join(lines, "\n")

	default:
		return fmt.Sprintf("%v", v)
	}
}

func sortedMapKeys(data map[string]any) []string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

func nodeIndent(depth int) string {
	return strings.Repeat(indentChar, depth*indentWidth-2)
}

func valueIndent(depth int) string {
	return strings.Repeat(indentChar, depth*indentWidth)
}

func closingIndent(depth int) string {
	return strings.Repeat(indentChar, (depth-1)*indentWidth)
}
