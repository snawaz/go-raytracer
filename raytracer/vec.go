
package raytracer

import (
    "math"
)

type Vec struct {
    X float64
    Y float64
    Z float64
}

type Point = Vec

func NewVec(x, y, z float64) *Vec {
    return &Vec{x, y, z}
}

func VecFrom(n float64) *Vec {
    return &Vec {n, n, n}
}

func Zeroes() *Vec {
    return VecFrom(0.0)
}

func Ones() *Vec {
    return VecFrom(1.0)
}

func (v * Vec) Negate() *Vec {
    return &Vec {-v.X, -v.Y, -v.Z}
}

func (v1 * Vec) Add(v2 *Vec) *Vec {
    return &Vec {v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 * Vec) Substract(v2 *Vec) *Vec {
    return &Vec {v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v1 * Vec) Multiply(v2 *Vec) *Vec {
    return &Vec {v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}

func (v * Vec) Scale(s float64) *Vec {
    return &Vec { v.X * s, v.Y * s, v.Z * s}
}

func (v * Vec) Shift(s float64) *Vec {
    return &Vec { v.X + s, v.Y + s, v.Z + s}
}

func (v1 * Vec) Dot(v2 *Vec) float64 {
    return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}

func (v1 * Vec) Cross(v2 *Vec) *Vec {
    return &Vec {v1.Y * v2.Z - v1.Z * v2.Y, v1.Z * v2.X - v1.X * v2.Z, v1.X * v2.Y - v1.Y * v2.X}
}

func (v * Vec) LengthSquared() float64 {
    return v.Dot(v)
}

func (v * Vec) Length() float64 {
    return math.Sqrt(v.LengthSquared())
}

func (v * Vec) Unit() *Vec {
    length := v.Length()
    return &Vec {v.X / length, v.Y / length, v.Z / length};
}
