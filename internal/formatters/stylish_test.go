package formatters

import (
	"code/internal/diff"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatStylish(t *testing.T) {
	tree := []diff.DiffNode{
		{
			Key:   "follow",
			Kind:  diff.KindAdded,
			Right: false,
		},
		{
			Key:  "proxy",
			Kind: diff.KindRemoved,
			Left: "123.234.53.22",
		},
		{
			Key:   "timeout",
			Kind:  diff.KindChanged,
			Left:  50,
			Right: 20,
		},
		{
			Key:  "group",
			Kind: diff.KindNested,
			Children: []diff.DiffNode{
				{
					Key:   "deep",
					Kind:  diff.KindAdded,
					Right: "value",
				},
			},
		},
		{
			Key:  "same",
			Kind: diff.KindUnchanged,
			Left: true,
		},
	}

	got, err := formatStylish(tree, 1)
	require.NoError(t, err)

	want := `{
  + follow: false
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
    group: {
      + deep: value
    }
    same: true
}`

	require.Equal(t, want, got)
}

func TestFormatValueNil(t *testing.T) {
	got := formatValue(nil, 1)
	require.Equal(t, "null", got)
}

func TestFormatValueMap(t *testing.T) {
	value := map[string]any{
		"b": "2",
		"a": map[string]any{
			"x": 1,
		},
	}

	got := formatValue(value, 1)

	want := `{
    a: {
        x: 1
    }
    b: 2
}`

	require.Equal(t, want, got)
}

func TestSortedMapKeys(t *testing.T) {
	data := map[string]any{
		"z": 1,
		"a": 2,
		"m": 3,
	}

	got := sortedMapKeys(data)

	require.Equal(t, []string{"a", "m", "z"}, got)
}

func TestFormatStylishUnknownKind(t *testing.T) {
	tree := []diff.DiffNode{
		{
			Key:  "broken",
			Kind: diff.NodeKind("mystery"),
		},
	}

	_, err := formatStylish(tree, 1)

	require.Error(t, err)
	require.ErrorContains(t, err, "(о_0) unknown diff node kind")
}
