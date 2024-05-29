package tree

import (
	. "github.com/patrykjadamczyk/go-utils/base"
)

// TreeNode is a Tree structure
// TreeNode contains ID, Data of the node and children
type TreeNode[IDT any, DT any] struct {
	ID       IDT
	NodeData DT
	Children []TreeNode[IDT, DT]
}

// AddChild is a function that adds child to the node
func (t *TreeNode[IDT, DT]) AddChild(child TreeNode[IDT, DT]) {
	t.Children = append(t.Children, child)
}

// Contains checks if the child is in the tree of current node
func (t TreeNode[IDT, DT]) Contains(child TreeNode[IDT, DT]) bool {
	if t.Children == nil || len(t.Children) == 0 {
		return false
	}

	for _, c := range t.Children {
		if UniversalEquals(c.ID, child.ID).GetValue() {
			return true
		}
		if c.Contains(child) {
			return true
		}
	}
	return false
}

// ContainsID checks if the childID is in the tree of current node
func (t TreeNode[IDT, DT]) ContainsID(childID IDT) bool {
	if t.Children == nil || len(t.Children) == 0 {
		return false
	}

	for _, c := range t.Children {
		if UniversalEquals(c.ID, childID).GetValue() {
			return true
		}
		if c.ContainsID(childID) {
			return true
		}
	}
	return false
}

// GetChild is a function that gets child by ID from the current tree node
func (t TreeNode[IDT, DT]) GetChild(childID IDT) Result[TreeNode[IDT, DT]] {
	if t.Children == nil || len(t.Children) == 0 {
		return Err[TreeNode[IDT, DT]](NewError("tree does not contain child"))
	}

	for _, c := range t.Children {
		if UniversalEquals(c.ID, childID).GetValue() {
			return Ok(c)
		}
		cr := c.GetChild(childID)
		if !cr.IsError() {
			return Ok(cr.Unwrap())
		}
	}

	return Err[TreeNode[IDT, DT]](NewError("tree does not contain child"))
}

// MakeTreeNode is a function that makes TreeNode with specified ID, Node data and childen of the node
func MakeTreeNode[IDT any, DT any](id IDT, data DT, children ...TreeNode[IDT, DT]) TreeNode[IDT, DT] {
	return TreeNode[IDT, DT]{ID: id, NodeData: data, Children: children}
}

// MakeTreeNodeWithoutChildren is a function that makes TreeNode with specified ID, Node data
func MakeTreeNodeWithoutChildren[IDT any, DT any](id IDT, data DT) TreeNode[IDT, DT] {
	return TreeNode[IDT, DT]{ID: id, NodeData: data}
}
