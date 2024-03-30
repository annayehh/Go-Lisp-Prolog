package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

// TODO: implement a type that satisfies the DisjointSet interface.
type node struct {
	parent *node
	value  int
	rank   int
}

type disjointSet struct {
	nodes map[int]*node
}

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() *disjointSet {
	return &disjointSet{nodes: make(map[int]*node)}
}

func (ds *disjointSet) MakeSet(x int) {
	ds.nodes[x] = &node{
		parent: nil,
		value:  x,
		rank:   0,
	}
}

func (ds *disjointSet) UnionSet(s, t int) int {
	sNode := ds.findOrCreate(s)
	tNode := ds.findOrCreate(t)
	return ds.union(sNode, tNode).value
}

func (ds *disjointSet) FindSet(s int) int {
	return ds.findOrCreate(s).value
}

func (ds *disjointSet) findOrCreate(s int) *node {
	if _, ok := ds.nodes[s]; !ok {
		// Create a new node since it doesn't exist
		ds.MakeSet(s)
	}
	return ds.find(ds.nodes[s])
}

func (ds *disjointSet) union(x, y *node) *node {
	xRoot := ds.find(x)
	yRoot := ds.find(y)

	if xRoot == yRoot {
		return xRoot
	}

	if xRoot.rank < yRoot.rank {
		xRoot.parent = yRoot
		return yRoot
	} else if xRoot.rank > yRoot.rank {
		yRoot.parent = xRoot
		return xRoot
	} else {
		yRoot.parent = xRoot
		xRoot.rank++
		return xRoot
	}
}

func (ds *disjointSet) find(x *node) *node {
	if x.parent == nil {
		return x
	}
	x.parent = ds.find(x.parent)
	return x.parent
}
