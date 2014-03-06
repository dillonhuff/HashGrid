package hashgrid

import "container/list"

type HashGrid struct {
	alpha float64
	grid map[GridCoor](*list.List)
}

type GridCoor struct {
	x, y int
}

func (g HashGrid) Add(v *Vertex) {
	current := g.GetSquare(v)
	if (current == nil) {
		current = list.New()
	}
	current.PushFront(v)
	g.grid[*g.GetGridCoor(v)] = current
}

func (g HashGrid) Alpha() float64 {
	return g.alpha
}

func (g HashGrid) GetSquare(v *Vertex) *list.List {
	gCor := g.GetGridCoor(v)
	return g.grid[*gCor]
}

func (g HashGrid) GetGridCoor(v *Vertex) *GridCoor {
	xc := int(v.x / g.alpha)
	yc := int(v.y / g.alpha)
	return &GridCoor{xc, yc}
}

func NewHashGrid(alpha float64) HashGrid{
	return HashGrid{alpha, make(map[GridCoor](*list.List))}
}