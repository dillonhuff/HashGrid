package hashgrid

import "container/list"
import "testing"

func TestEmptyGridAlpha(t *testing.T) {
	m := NewHashGrid(1.0)
	if !(m.Alpha() == 1.0) {
		t.Errorf("Expected alpha = ", 1.0, "but got", m.Alpha())
	}
}

func TestNoElemRetrieve(t *testing.T) {
	m := NewHashGrid(1.2)
	v := &Vertex{1.2, 3.4}
	l := m.GetSquare(v)
	if l != nil {
		t.Errorf("Expected ", nil, "but got non-nil")
	}
}

func TestSameElemRetrieve(t *testing.T) {
	m := NewHashGrid(5.0)
	m.Add(&Vertex{10, -12.3})
	l := list.New()
	l.PushFront(&Vertex{10, -12.3})
	if !SameElems(l, m.GetSquare(&Vertex{10, -12.3})) {
		t.Errorf("Did not retrieve correct list")
	}
}

func TestMultipleElemRetrieve(t *testing.T) {
	m := NewHashGrid(29.3)
	m.Add(&Vertex{82.1, 98.1})
	m.Add(&Vertex{-12.1, -34.4})
	m.Add(&Vertex{82.0, 99.2})
	l := list.New()
	l.PushFront(&Vertex{82.1, 98.1})
	l.PushFront(&Vertex{82.0, 99.2})
	if !SameElems(l, m.GetSquare(&Vertex{86.4, 89.987})) {
		t.Errorf("Multiple elements in same square failed")
	}
}

func TestUpperLeftCluster(t *testing.T) {
	m := NewHashGrid(10.0)
	m.Add(&Vertex{2.3, 7.9})
	m.Add(&Vertex{-11.2, 11.1})
	l := list.New()
	l.PushFront(&Vertex{2.3, 7.9})
	if !SameElems(l, m.GetCluster(&Vertex{14.5, 19.9})) {
		t.Errorf("Wrong vertices returned for cluster")
	}
}

func SameElems(l1, l2 *list.List) bool {
	if (l1.Len() != l2.Len()) {
		return false
	}
	e2 := l2.Front()
	for e1 := l1.Front(); e1 != nil; e1 = e1.Next() {
		var v1, v2 *Vertex
		v1, v2 = e1.Value.(*Vertex), e2.Value.(*Vertex)
		if (!v1.Eq(v2)) {
			return false
		}
		e2 = e2.Next()
	}
	return true
}