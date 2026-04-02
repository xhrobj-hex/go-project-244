package formatter

import (
	"code/internal/diff"
	"fmt"
	"sort"
	"strings"
)

const (
	indentChar  = " "
	indentWidth = 4
)

func formatTreeStylish(tree []diff.DiffNode, depth int) (string, error) {
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
			childTree, err := formatTreeStylish(node.Children, depth+1)
			if err != nil {
				return "", err
			}

			lines = append(lines, formatLine(depth, " ", node.Key, childTree))

		default:
			return "", fmt.Errorf("(о_0) unknown diff node kind: %q", node.Kind)
		}
	}

	lines = append(lines, fmt.Sprintf("%s}", closingIndent(depth)))

	return strings.Join(lines, "\n"), nil
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
