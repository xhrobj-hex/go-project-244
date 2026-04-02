package formatters

import (
	"code/internal/diff"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatPlainValue(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  string
	}{
		{
			name:  "nil",
			value: nil,
			want:  "null",
		},
		{
			name:  "string",
			value: "hexlet",
			want:  "'hexlet'",
		},
		{
			name:  "bool",
			value: true,
			want:  "true",
		},
		{
			name:  "int",
			value: 10,
			want:  "10",
		},
		{
			name:  "complex value",
			value: map[string]any{"key": "value"},
			want:  "[complex value]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatPlainValue(tt.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFormatPlain(t *testing.T) {
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
					Key:  "abc",
					Kind: diff.KindRemoved,
					Left: 123,
				},
				{
					Key:   "deep",
					Kind:  diff.KindAdded,
					Right: map[string]any{"k": "v"},
				},
			},
		},
		{
			Key:  "same",
			Kind: diff.KindUnchanged,
			Left: "value",
		},
	}

	got, err := formatPlain(tree)
	require.NoError(t, err)

	want := `Property 'follow' was added with value: false
Property 'proxy' was removed
Property 'timeout' was updated. From 50 to 20
Property 'group.abc' was removed
Property 'group.deep' was added with value: [complex value]`

	require.Equal(t, want, got)
}
