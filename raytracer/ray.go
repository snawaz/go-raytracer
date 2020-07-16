
package raytracer

type Ray struct {
    Origin Point
    Direction Vec
}

func NewRay(origin *Point, direction *Vec) Ray{
    return Ray {*origin, *direction}
}

func (ray * Ray) PointAt(t float64) *Point {
    return ray.Origin.Add(ray.Direction.Scale(t))
}
