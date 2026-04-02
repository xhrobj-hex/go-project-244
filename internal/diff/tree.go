package diff

import (
	"reflect"
	"sort"
)

// NodeKind описывает тип узла в дереве различий.
type NodeKind string

const (
	// KindAdded обозначает ключ, который есть только в правом наборе данных.
	KindAdded NodeKind = "added"

	// KindRemoved обозначает ключ, который есть только в левом наборе данных.
	KindRemoved NodeKind = "removed"

	// KindChanged обозначает ключ, значение которого изменилось.
	KindChanged NodeKind = "changed"

	// KindUnchanged обозначает ключ, значение которого не изменилось.
	KindUnchanged NodeKind = "unchanged"

	// KindNested обозначает ключ, значения которого с обеих сторон являются
	// вложенными объектами и требуют рекурсивного сравнения.
	KindNested NodeKind = "nested"
)

// DiffNode описывает узел дерева различий.
// Хранит информацию о ключе, типе изменения и дочерних узлах
// для вложенных объектов.
type DiffNode struct {
	Key      string
	Kind     NodeKind
	Left     any
	Right    any
	Children []DiffNode
}

// BuildTree строит дерево различий для двух объектов,
// представленных в виде map[string]any.
func BuildTree(leftData, rightData map[string]any) []DiffNode {
	keys := sortedUnionKeys(leftData, rightData)

	tree := make([]DiffNode, 0, len(keys))

	for _, key := range keys {
		leftValue, leftOK := leftData[key]
		rightValue, rightOK := rightData[key]

		leftObj, leftIsObj := asMap(leftValue)
		rightObj, rightIsObj := asMap(rightValue)

		switch {
		case !leftOK:
			tree = append(tree, DiffNode{
				Key:   key,
				Kind:  KindAdded,
				Right: rightValue,
			})

		case !rightOK:
			tree = append(tree, DiffNode{
				Key:  key,
				Kind: KindRemoved,
				Left: leftValue,
			})

		case leftIsObj && rightIsObj:
			tree = append(tree, DiffNode{
				Key:      key,
				Kind:     KindNested,
				Children: BuildTree(leftObj, rightObj),
			})

		case !reflect.DeepEqual(leftValue, rightValue):
			tree = append(tree, DiffNode{
				Key:   key,
				Kind:  KindChanged,
				Left:  leftValue,
				Right: rightValue,
			})

		default: // NOTE: "unchanged" (leftValue = rightValue)
			tree = append(tree, DiffNode{
				Key:   key,
				Kind:  KindUnchanged,
				Left:  leftValue,
				Right: rightValue,
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
