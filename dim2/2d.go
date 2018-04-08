package dim2

import "math"

type Point struct {
	X, Y float64
}

// ConvexHull returns the convex hull of a set of points.
func ConvexHull(points []Point) []Point {
	if len(points) < 3 {
		panic("there must be at least 3 points")
	}

	hull := []Point{}

	// Slice index of the leftmost point
	l := 0
	for i := 1; i < len(points); i++ {
		if points[i].X < points[l].X {
			l = i
		}
	}

	// Start from leftmost point, keep moving counterclockwise
	// until reach the start point again.  This loop runs O(h)
	// times where h is number of points in result or output.
	p := l
	q := 0
	for {
		hull = append(hull, points[p])

		// Search for a point 'q' such that orientation(p, x,
		// q) is counterclockwise for all points 'x'. The idea
		// is to keep track of last visited most counterclock-
		// wise point in q. If any point 'i' is more counterclock-
		// wise than q, then update q.
		q = (p + 1) % len(points)
		for i := 0; i < len(points); i++ {
			// If i is more counterclockwise than current q, then update q
			if orientation(points[p], points[i], points[q]) == 2 {
				q = i
			}
		}

		// Now q is the most counterclockwise with respect to p
		// Set p as q for next iteration, so that q is added to result 'hull'
		p = q

		// While we don't come to first point
		if p == l {
			break
		}
	}

	return hull
}

// To find orientation of ordered triplet (p, q, r).
// The function returns following values
// 0 --> p, q and r are colinear
// 1 --> Clockwise
// 2 --> Counterclockwise
func orientation(p, q, r Point) int {
	val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)
	if val == 0 {
		return 0 // Colinear
	}
	if val > 0 {
		return 1 // Clockwise
	}
	return 2 // Counter clockwise
}

// PolygonArea returns the area of the polygon described by the given slice of ordered Points.
// It uses the shoelace formula.
func PolygonArea(points []Point) float64 {
	var first, second float64
	for i := 0; i < len(points); i++ {
		first += points[i].X * points[(i+1)%len(points)].Y
		second += points[i].Y * points[(i+1)%len(points)].X
	}
	return math.Abs((first - second) / 2)
}
