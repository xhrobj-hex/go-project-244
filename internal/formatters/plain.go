package formatters

import (
	"code/internal/diff"
	"fmt"
	"strings"
)

func formatPlain(tree []diff.DiffNode) (string, error) {
	lines, err := walkPlain(tree, "")
	if err != nil {
		return "", err
	}

	return strings.Join(lines, "\n"), nil
}

func walkPlain(tree []diff.DiffNode, parentPath string) ([]string, error) {
	lines := make([]string, 0)

	for _, node := range tree {
		path := node.Key
		if parentPath != "" {
			path = parentPath + "." + node.Key
		}

		switch node.Kind {
		case diff.KindAdded:
			lines = append(lines,
				fmt.Sprintf(
					"Property '%s' was added with value: %s",
					path,
					formatPlainValue(node.Right),
				),
			)

		case diff.KindRemoved:
			lines = append(lines,
				fmt.Sprintf(
					"Property '%s' was removed",
					path,
				),
			)

		case diff.KindChanged:
			lines = append(lines,
				fmt.Sprintf(
					"Property '%s' was updated. From %s to %s",
					path,
					formatPlainValue(node.Left),
					formatPlainValue(node.Right),
				),
			)

		case diff.KindUnchanged:
			continue

		case diff.KindNested:
			childLines, err := walkPlain(node.Children, path)
			if err != nil {
				return nil, err
			}
			lines = append(lines, childLines...)

		default:
			return nil, fmt.Errorf("(о_0) unknown diff node kind: %q", node.Kind)
		}
	}

	return lines, nil
}

func formatPlainValue(value any) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case string:
		return fmt.Sprintf("'%s'", v)
	case map[string]any:
		return "[complex value]"
	default:
		return fmt.Sprintf("%v", v)
	}
}
