package formatters

import (
	"code/internal/diff"
	"encoding/json"
	"fmt"
)

func formatJSON(tree []diff.DiffNode) (string, error) {
	root, err := buildJSONTree(tree)
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(root)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func buildJSONTree(tree []diff.DiffNode) (map[string]any, error) {
	result := make(map[string]any, len(tree))

	for _, node := range tree {
		entry, err := buildJSONNode(node)
		if err != nil {
			return nil, err
		}

		result[node.Key] = entry
	}

	return result, nil
}

func buildJSONNode(node diff.DiffNode) (map[string]any, error) {
	entry := map[string]any{
		"type": node.Kind,
	}

	switch node.Kind {
	case diff.KindAdded:
		entry["new_value"] = node.Right
	case diff.KindRemoved:
		entry["old_value"] = node.Left
	case diff.KindChanged, diff.KindUnchanged:
		entry["old_value"] = node.Left
		entry["new_value"] = node.Right
	case diff.KindNested:
		children, err := buildJSONTree(node.Children)
		if err != nil {
			return nil, err
		}
		entry["children"] = children
	default:
		return nil, fmt.Errorf("(о_0) unknown diff node kind: %q", node.Kind)
	}

	return entry, nil
}
