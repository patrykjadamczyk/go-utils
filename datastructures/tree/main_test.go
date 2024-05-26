package tree

import (
	"testing"

	"github.com/patrykjadamczyk/go-utils/math"
)

func TestTree(t *testing.T) {
	mtn := func(
		id math.Vector2[int],
		data int,
		children ...TreeNode[math.Vector2[int], int],
	) TreeNode[math.Vector2[int], int] {
		return MakeTreeNode(id, data, children...)
	}
	mtnwoc := func(
		id math.Vector2[int],
		data int,
	) TreeNode[math.Vector2[int], int] {
		return MakeTreeNodeWithoutChildren(id, data)
	}
	zv2 := math.ZeroVector2[int]()
	c1 := mtnwoc(
		math.MakeVector2(0, 1),
		1,
	)
	tr := mtn(
		zv2,
		0,
		c1,
		mtnwoc(
			math.MakeVector2(1, 0),
			2,
		),
	)
	if !tr.ContainsID(math.MakeVector2(0, 1)) {
		t.Error("TreeNode should contain ID that was added to the tree")
	}
	if !tr.ContainsID(math.MakeVector2(1, 0)) {
		t.Error("TreeNode should contain ID that was added to the tree")
	}
	if tr.ContainsID(math.MakeVector2(1, 1)) {
		t.Error("TreeNode should not contain ID that was not on the tree")
	}
	if !tr.Contains(c1) {
		t.Error("TreeNode should contain child that was added to the tree")
	}
}
