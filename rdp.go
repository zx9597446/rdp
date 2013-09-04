package rdp

import (
	"image"

	"math"
)

func findPerpendicularDistance(p, p1, p2 image.Point) (result float64) {
	if p1.X == p2.X {
		result = math.Abs(float64(p.X - p1.X))
	} else {
		slope := float64(p2.Y-p1.Y) / float64(p2.X-p1.X)
		intercept := float64(p1.Y) - (slope * float64(p1.X))
		result = math.Abs(slope*float64(p.X)-float64(p.Y)+intercept) / math.Sqrt(math.Pow(slope, 2)+1)
	}
	return
}

func Process(points []image.Point, epsilon float64) []image.Point {
	firstPoint := points[0]
	lastPoint := points[len(points)-1]
	if len(points) < 3 {
		return points
	}
	index := -1
	dist := float64(0)
	for i := 1; i < len(points)-1; i++ {
		cDist := findPerpendicularDistance(points[i], firstPoint, lastPoint)
		if cDist > dist {
			dist = cDist
			index = i
		}
	}
	if dist > epsilon {
		l1 := points[0 : index+1]
		l2 := points[index:]
		r1 := Process(l1, epsilon)
		r2 := Process(l2, epsilon)
		rs := append(r1[0:len(r1)-1], r2...)
		return rs
	} else {
		ret := make([]image.Point, 0)
		ret = append(ret, firstPoint, lastPoint)
		return ret
	}
	return make([]image.Point, 0)
}
