package code

import (
	"code/internal/diff"
	"code/internal/parser"
	"fmt"
	"sort"
	"strings"
)

const (
	indentChar  = " "
	indentWidth = 4
)

// GenDiff вычисляет различия между двумя файлами и возвращает их
// в виде строки в указанном формате.
func GenDiff(leftPath, rightPath, format string) (string, error) {
	leftData, rightData, err := parseFiles(leftPath, rightPath)
	if err != nil {
		return "", err
	}

	tree := diff.BuildTree(leftData, rightData)
	rootDepth := 1

	return formatTree(tree, rootDepth, format), nil
}

func parseFiles(leftPath, rightPath string) (map[string]any, map[string]any, error) {
	leftData, err := parser.Parse(leftPath)
	if err != nil {
		return nil, nil, fmt.Errorf("parse left file: %w", err)
	}

	rightData, err := parser.Parse(rightPath)
	if err != nil {
		return nil, nil, fmt.Errorf("parse right file: %w", err)
	}

	return leftData, rightData, nil
}

func formatTree(tree []diff.DiffNode, depth int, format string) string {
	_ = format // NOTE: пока не используем ...

	lines := []string{"{"}

	for _, node := range tree {
		switch node.Kind {
		case diff.KindAdded:
			lines = append(lines, formatLine(depth, "+", node.Key, formatValue(node.Right, depth+1)))

		case diff.KindRemoved:
			lines = append(lines, formatLine(depth, "-", node.Key, formatValue(node.Left, depth+1)))

		case diff.KindChanged:
			lines = append(lines,
				formatLine(depth, "-", node.Key, formatValue(node.Left, depth+1)),
				formatLine(depth, "+", node.Key, formatValue(node.Right, depth+1)),
			)

		case diff.KindUnchanged:
			lines = append(lines, formatLine(depth, " ", node.Key, formatValue(node.Left, depth+1)))

		case diff.KindNested:
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
