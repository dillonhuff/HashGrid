package hashgrid

import "math"
import "math/rand"
import "strconv"

type Vertex struct {
	x, y float64
}

func (v1 *Vertex) Dist(v2 *Vertex) float64 {
	dx := v1.x - v2.x;
	dy := v1.y - v2.y;
	return math.Sqrt(dx*dx + dy*dy)
}

func (v1 *Vertex) Eq(v2 *Vertex) bool {
	return v1.x == v2.x && v1.y == v2.y
}

func (v Vertex) Show() string {
	return "( " + strconv.FormatFloat(v.x, 'f', 1, 64) + ", " + strconv.FormatFloat(v.y, 'f', 1, 64) + " )"
}

func randVertex(scale float64) Vertex {
	return Vertex{scale*rand.Float64(), scale*rand.Float64()}
}