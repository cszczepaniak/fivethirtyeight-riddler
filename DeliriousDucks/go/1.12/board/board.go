package board

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

type Board []Point

func NewBoard(w, h int) Board {
	res := make([]Point, 0, w*h)
	for i := 0; i < h; i++ {
		r := make([]Point, w)
		for j := 0; j < w; j++ {
			r[j] = NewPoint(j, i)
		}
		res = append(res, r...)
	}
	return res
}
