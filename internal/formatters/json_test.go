package formatters

import (
	"code/internal/diff"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatJSON(t *testing.T) {
	tree := []diff.DiffNode{
		{
			Key:   "follow",
			Kind:  diff.KindAdded,
			Right: false,
		},
		{
			Key:  "group",
			Kind: diff.KindNested,
			Children: []diff.DiffNode{
				{
					Key:   "deep",
					Kind:  diff.KindChanged,
					Left:  1,
					Right: 2,
				},
			},
		},
	}

	got, err := formatJSON(tree)
	require.NoError(t, err)

	var decoded map[string]any
	err = json.Unmarshal([]byte(got), &decoded)
	require.NoError(t, err)

	require.Len(t, decoded, 2)

	follow := decoded["follow"].(map[string]any)
	require.Equal(t, "added", follow["type"])
	require.Equal(t, false, follow["new_value"])

	group := decoded["group"].(map[string]any)
	require.Equal(t, "nested", group["type"])

	children := group["children"].(map[string]any)
	deep := children["deep"].(map[string]any)
	require.Equal(t, "changed", deep["type"])
	require.Equal(t, float64(1), deep["old_value"])
	require.Equal(t, float64(2), deep["new_value"])
}
