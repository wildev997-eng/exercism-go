package pov

type Tree struct {
	value    string
	children []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	return &Tree{
		value:    value,
		children: children,
	}
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	return tr.reroot(from, nil)
}

func (tr *Tree) reroot(target string, up *Tree) *Tree {
	if tr == nil {
		return nil
	}

	if tr.value == target {
		kids := append([]*Tree{}, tr.children...)
		if up != nil {
			kids = append(kids, up)
		}
		return &Tree{value: tr.value, children: kids}
	}

	for i, child := range tr.children {
		others := append([]*Tree{}, tr.children[:i]...)
		others = append(others, tr.children[i+1:]...)
		if up != nil {
			others = append(others, up)
		}

		parent := &Tree{value: tr.value, children: others}
		if res := child.reroot(target, parent); res != nil {
			return res
		}
	}
	return nil
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	x := tr.FromPov(from)
	return pathDown(x, to)
}

func pathDown(tr *Tree, target string) []string {
	if tr == nil {
		return nil
	}

	if tr.value == target {
		return []string{tr.value}
	}

	for _, child := range tr.children {
		result := pathDown(child, target)

		if result != nil {
			return append([]string{tr.value}, result...)
		}
	}

	return nil
}
