
package raytracer

import (
    "math"
)

type Camera struct{
    Origin          *Point
    LowerLeftCorner *Point
    Horizontal      *Vec
    Vertical        *Vec
    u, v, w         *Vec
    LensRadius       float64
}

func NewCamera(lookFrom, lookAt *Point, viewUp *Vec, verticalFov, aspectRatio, aperture, focusDistance float64) Camera {
    theta := verticalFov * math.Pi / 180.0
    h := math.Tan (theta / 2.0)
    viewportHeight := 2.0 * h
    viewportWidth := aspectRatio * viewportHeight

    w := lookFrom.Substract(lookAt).Unit()
    u := viewUp.Cross(w).Unit()
    v := w.Cross(u)

    origin := lookFrom
    horizontal := u.Scale(viewportWidth * focusDistance)
    vertical := v.Scale(viewportHeight * focusDistance)
    lowerLeftCorner := origin.Substract(horizontal.Scale(1.0/2.0)).Substract(vertical.Scale(1.0/2.0)).Substract(w.Scale(focusDistance))
    lensRadius := aperture / 2.0
    return Camera {origin, lowerLeftCorner, horizontal, vertical, u, v, w, lensRadius}
}

func (c * Camera) RayAt(s, t float64) Ray {
    rd := SampleVecInUnitDisk().Scale(c.LensRadius)
    offset := c.u.Scale(rd.X).Add(c.v.Scale(rd.Y))
    origin := c.Origin.Add(offset)
    direction := c.LowerLeftCorner.Add(c.Horizontal.Scale(s)).Add(c.Vertical.Scale(t)).Substract(origin)
    return NewRay(origin, direction) 
}
